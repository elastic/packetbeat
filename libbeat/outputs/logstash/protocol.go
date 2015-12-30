package logstash

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
	"time"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

type protocol struct {
	conn          net.Conn
	timeout       time.Duration
	compressLevel int

	eventsBuffer *bytes.Buffer
}

var (
	// ErrProtocolError is returned if an protocol error was detected in the
	// conversation with lumberjack server.
	ErrProtocolError = errors.New("lumberjack protocol error")

	errAllEventsEncoding = errors.New("failed to encode all events")
)

var (
	codeVersion byte = '2'

	codeWindowSize    = []byte{codeVersion, 'W'}
	codeJSONDataFrame = []byte{codeVersion, 'J'}
	codeCompressed    = []byte{codeVersion, 'C'}
)

func newClientProcol(
	conn net.Conn,
	timeout time.Duration,
	compressLevel int,
) (*protocol, error) {

	// validate by creating and discarding zlib writer with configured level
	if compressLevel > 0 {
		tmp := bytes.NewBuffer(nil)
		w, err := zlib.NewWriterLevel(tmp, compressLevel)
		if err != nil {
			return nil, err
		}
		w.Close()
	}

	return &protocol{
		conn:          conn,
		timeout:       timeout,
		compressLevel: compressLevel,
		eventsBuffer:  bytes.NewBuffer(nil),
	}, nil
}

func (p *protocol) sendEvents(events []common.MapStr) (uint32, error) {
	conn := p.conn

	if len(events) == 0 {
		return 0, nil
	}

	// serialize all raw events into output buffer, removing all events encoding failed for
	count, err := p.serializeEvents(events)
	if count == 0 {
		// encoding of all events failed. Let's stop here and report all events
		// as exported so no one tries to send/encode the same events once again
		// The compress/encode function already prints critical per failed encoding
		// failure.
		return 0, errAllEventsEncoding
	}

	// send window size:
	if err := p.sendWindowSize(count); err != nil {
		return 0, err
	}

	if p.compressLevel > 0 {
		err = p.sendCompressed(p.eventsBuffer.Bytes())
	} else {
		_, err = conn.Write(p.eventsBuffer.Bytes())
	}
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *protocol) recvACK() (uint32, error) {
	conn := p.conn

	if err := conn.SetReadDeadline(time.Now().Add(p.timeout)); err != nil {
		return 0, err
	}

	response := make([]byte, 6)
	ackbytes := 0
	for ackbytes < 6 {
		n, err := conn.Read(response[ackbytes:])
		if err != nil {
			return 0, err
		}
		ackbytes += n
	}

	isACK := response[0] == codeVersion && response[1] == 'A'
	if !isACK {
		return 0, ErrProtocolError
	}
	seq := binary.BigEndian.Uint32(response[2:])
	return seq, nil
}

// wait for ACK (accept partial ACK to reset timeout)
// reset timeout timer for every ACK received.
func (p *protocol) awaitACK(count uint32) (uint32, error) {
	var ackSeq uint32
	var err error

	// read until all acks
	for ackSeq < count {
		ackSeq, err = p.recvACK()
		if err != nil {
			return ackSeq, err
		}
	}
	return ackSeq, nil
}

func (p *protocol) sendWindowSize(window uint32) error {
	conn := p.conn

	if err := conn.SetWriteDeadline(time.Now().Add(p.timeout)); err != nil {
		return err
	}
	if _, err := conn.Write(codeWindowSize); err != nil {
		return err
	}
	return writeUint32(conn, window)
}

func (p *protocol) sendCompressed(payload []byte) error {
	conn := p.conn

	if err := conn.SetWriteDeadline(time.Now().Add(p.timeout)); err != nil {
		return err
	}

	if _, err := conn.Write(codeCompressed); err != nil {
		return err
	}

	if err := writeUint32(conn, uint32(len(payload))); err != nil {
		return err
	}

	for len(payload) > 0 {
		n, err := conn.Write(payload)
		payload = payload[n:]
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *protocol) serializeEvents(events []common.MapStr) (uint32, error) {
	p.eventsBuffer.Reset()

	if p.compressLevel > 0 {
		w, _ := zlib.NewWriterLevel(p.eventsBuffer, p.compressLevel)
		count, err := p.doSerializeEvents(w, events)
		if err != nil {
			return 0, err
		}
		if err := w.Close(); err != nil {
			debug("Finalizing zlib compression failed with: %s", err)
			return 0, err
		}
		return count, nil
	}

	return p.doSerializeEvents(p.eventsBuffer, events)
}

func (p *protocol) doSerializeEvents(out io.Writer, events []common.MapStr) (uint32, error) {
	var sequence uint32
	for _, event := range events {
		sequence++
		err := p.serializeDataFrame(out, event, sequence)
		if err != nil {
			logp.Critical("failed to encode event: %v", err)
			sequence-- //forget this last broken event and continue
		}
	}
	return sequence, nil
}

func (p *protocol) serializeDataFrame(
	out io.Writer,
	event common.MapStr,
	seq uint32,
) error {
	// Write JSON Data Frame:
	// version: uint8 = '2'
	// code: uint8 = 'J'
	// seq: uint32
	// payloadLen (bytes): uint32
	// payload: JSON document

	jsonEvent, err := json.Marshal(event)
	if err != nil {
		debug("Fail to convert the event to JSON: %s", err)
		return err
	}

	if _, err := out.Write(codeJSONDataFrame); err != nil { // version + code
		return err
	}
	if err := writeUint32(out, seq); err != nil {
		return err
	}
	if err := writeUint32(out, uint32(len(jsonEvent))); err != nil {
		return err
	}
	if _, err := out.Write(jsonEvent); err != nil {
		return err
	}

	return nil
}

func writeUint32(out io.Writer, v uint32) error {
	return binary.Write(out, binary.BigEndian, v)
}

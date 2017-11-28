package smtp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/mail"
	"time"

	"github.com/satori/go.uuid"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"

	"github.com/elastic/beats/packetbeat/protos"
	"github.com/elastic/beats/packetbeat/protos/tcp"
)

// Transaction Publisher.
type transPub struct {
	sendRequest     bool
	sendResponse    bool
	sendDataHeaders bool
	sendDataBody    bool

	results protos.Reporter
}

func (pub *transPub) onTransaction(
	trans transaction,
	sessionID uuid.UUID,
) error {
	if pub.results == nil {
		return nil
	}

	pub.results(pub.createEvent(trans, sessionID))

	return nil
}

func (pub *transPub) createEvent(
	trans transaction,
	sessionID uuid.UUID,
) beat.Event {
	fields := common.MapStr{
		"type":   "smtp",
		"status": common.OK_STATUS,
	}
	d := common.MapStr{"session_id": sessionID}
	drequ := common.MapStr{}
	dresp := common.MapStr{}

	var ts time.Time

	switch t := trans.(type) {

	case *transPrompt:
		d["type"] = "PROMPT"
		if pub.sendResponse {
			fields["response"] =
				common.NetString(t.resp.raw.BufferedBytes())
		}
		if t.resp.statusCode >= 400 {
			fields["status"] = common.SERVER_ERROR_STATUS
		}
		fields["bytes_out"] = t.resp.Size
		dresp["code"] = t.resp.statusCode
		if len(t.resp.statusPhrases) > 0 {
			dresp["phrases"] = t.resp.statusPhrases
		}
		fields["src"], fields["dst"] = pub.getEndpoints(t.resp)
		ts = t.resp.Ts

	case *transCommand:
		d["type"] = "COMMAND"
		if pub.sendRequest {
			fields["request"] =
				common.NetString(t.requ.raw.BufferedBytes())
		}
		if pub.sendResponse {
			fields["response"] =
				common.NetString(t.resp.raw.BufferedBytes())
		}
		fields["bytes_in"] = t.requ.Size
		fields["bytes_out"] = t.resp.Size
		if t.resp.statusCode >= 400 {
			fields["status"] = common.SERVER_ERROR_STATUS
		}
		// Response time in milliseconds
		fields["responsetime"] =
			int32(t.resp.Ts.Sub(t.requ.Ts).Nanoseconds() / 1e6)
		drequ["command"] = t.requ.command
		drequ["param"] = t.requ.param
		dresp["code"] = t.resp.statusCode
		dresp["phrases"] = t.resp.statusPhrases
		fields["src"], fields["dst"] = pub.getEndpoints(t.requ)
		ts = t.resp.Ts

	case *transMail:
		d["type"] = "MAIL"
		if t.reversePath != nil {
			drequ["envelope_sender"] = t.reversePath
		}
		if t.forwardPaths != nil {
			drequ["envelope_recipients"] = t.forwardPaths
		}
		if bytes.Equal(t.requ.command, constEOD) {
			headers, body, err := pub.parsePayload(t)
			if err != nil {
				msg := fmt.Sprintf("Failed to parse data payload: %s", err)
				t.Notes = append(t.Notes, msg)
				debugf(msg)
			} else {
				if pub.sendDataHeaders {
					drequ["headers"] = headers
				}
				if pub.sendDataBody {
					drequ["body"] = body
				}
			}
		}
		fields["bytes_in"] = t.BytesIn
		fields["bytes_out"] = t.BytesOut
		fields["status"] = t.Status
		if len(t.Notes) > 0 {
			fields["notes"] = t.Notes
		}
		fields["src"], fields["dst"] = pub.getEndpoints(t.requ)
		ts = t.resp.Ts

	}

	if len(drequ) > 0 {
		d["request"] = drequ
	}
	if len(dresp) > 0 {
		d["response"] = dresp
	}

	fields["smtp"] = d

	return beat.Event{
		Timestamp: ts,
		Fields:    fields,
	}
}

func (pub *transPub) parsePayload(t *transMail) (
	map[string]common.NetString,
	common.NetString,
	error,
) {
	if !pub.sendDataHeaders && !pub.sendDataBody {
		return nil, nil, nil
	}

	var headers map[string]common.NetString
	var body []byte
	var err error

	payload, err := mail.ReadMessage(&t.requ.raw)
	if err != nil {
		return nil, nil, err
	}

	headers = make(map[string]common.NetString)

	for k := range payload.Header {
		headers[k] = common.NetString(payload.Header.Get(k))
	}

	if body, err = ioutil.ReadAll(payload.Body); err != nil {
		return nil, nil, err
	}

	return headers, common.NetString(body), nil
}

func (pub *transPub) getEndpoints(m *message) (
	*common.Endpoint,
	*common.Endpoint,
) {
	src := &common.Endpoint{
		IP:   m.Tuple.SrcIP.String(),
		Port: m.Tuple.SrcPort,
		Proc: string(m.CmdlineTuple.Src),
	}
	dst := &common.Endpoint{
		IP:   m.Tuple.DstIP.String(),
		Port: m.Tuple.DstPort,
		Proc: string(m.CmdlineTuple.Dst),
	}
	if m.Direction == tcp.TCPDirectionReverse {
		src, dst = dst, src
	}

	return src, dst
}

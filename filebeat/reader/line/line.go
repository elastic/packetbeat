package line

import (
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/streambuf"
	"github.com/elastic/beats/libbeat/logp"
)

// lineReader reads lines from underlying reader, decoding the input stream
// using the configured codec. The reader keeps track of bytes consumed
// from raw input stream for every decoded line.
type Reader struct {
	lineScanner *lineScanner
}

// New creates a new reader object
func New(input io.Reader, codec encoding.Encoding, bufferSize int) (*Reader, error) {
	encoder := codec.NewEncoder()

	// Create newline char based on encoding
	nl, _, err := transform.Bytes(encoder, []byte{'\n'})
	if err != nil {
		return nil, err
	}

	decReader := newDecoderReader(input, codec)
	lineScanner := newLineScanner(decReader, nl, bufferSize)

	return &Reader{
		lineScanner: lineScanner,
	}, nil
}

// Next reads the next line until the new line character
func (r *Reader) Next() ([]byte, int, error) {
	return r.lineScanner.scan()
}

func (r *Reader) GetState() common.MapStr {
	return common.MapStr{
		"decoder": common.MapStr{
			"decoder": r.lineScanner.in.decodedOffset,
			"encoded": r.lineScanner.in.encodedOffset,
			"file":    r.lineScanner.in.fileOffset,
		},
		"scanner": common.MapStr{
			"line": r.lineScanner.offset,
		},
	}
}

type decoderReader struct {
	in         io.Reader
	decoder    transform.Transformer
	decoderBuf *streambuf.Buffer

	fileOffset    int
	encodedOffset int
	decodedOffset int
}

func newDecoderReader(in io.Reader, codec encoding.Encoding) *decoderReader {
	return &decoderReader{
		in:            in,
		decoder:       codec.NewDecoder(),
		decoderBuf:    streambuf.New(nil),
		fileOffset:    0,
		encodedOffset: 0,
		decodedOffset: 0,
	}
}

func (r *decoderReader) read(buf []byte) (int, error) {
	var buffer []byte
	out := streambuf.New(nil)

	if r.decoderBuf.Len() == 0 {
		buffer = make([]byte, 1024)
		n, err := r.in.Read(buffer)
		r.decoderBuf.Append(buffer[:n])

		if err != nil {
			return 0, err
		}
		if n == 0 {
			return 0, streambuf.ErrNoMoreBytes
		}

		r.fileOffset += n
	}

	inBytes := r.decoderBuf.Bytes()
	outBytes := make([]byte, 2048)

	end := len(inBytes)
	start := 0
	for start < end {
		nDst, nSrc, err := r.decoder.Transform(outBytes, inBytes[start:end], false)
		if err != nil {
			if err != transform.ErrShortDst {
				out.Write(inBytes[:end])
				start = end
				break
			}
		}

		start += nSrc
		r.decoderBuf.Advance(nSrc)
		r.decoderBuf.Reset()
		out.Write(outBytes[:nDst])

		r.encodedOffset += nSrc
		r.decodedOffset += nDst
	}

	dec, err := out.Collect(out.Len())
	if err != nil {
		panic(err)
	}

	return copy(buf, dec), nil
}

type lineScanner struct {
	in         *decoderReader
	nl         []byte
	bufferSize int

	buf    *streambuf.Buffer
	offset int
}

func newLineScanner(in *decoderReader, nl []byte, bufferSize int) *lineScanner {
	return &lineScanner{
		in:         in,
		nl:         nl,
		bufferSize: bufferSize,
		buf:        streambuf.New(nil),
		offset:     0,
	}
}

// Scan reads from the underlying decoder reader and returns decoded lines.
func (s *lineScanner) scan() ([]byte, int, error) {
	idx := s.buf.Index(s.nl)
	for !newLineFound(idx) {

		b := make([]byte, s.bufferSize)
		n, err := s.in.read(b)
		if err != nil {
			return nil, 0, err
		}

		// This can happen if something goes wrong during decoding
		if n == 0 {
			logp.Err("Empty buffer returned by read")
		}

		s.buf.Append(b[:n])
		idx = s.buf.Index(s.nl)
	}

	return s.line(idx)
}

// newLineFound checks if a new line was found.
func newLineFound(i int) bool {
	return i != -1
}

// line sets the offset of the scanner and returns a line.
func (s *lineScanner) line(i int) ([]byte, int, error) {
	line, err := s.buf.CollectUntil(s.nl)
	if err != nil {
		panic(err)
	}

	s.offset = s.offset + i
	s.buf.Reset()
	return line, len(line), nil
}

package reader

import (
	"io"
	"time"

	"github.com/elastic/beats/filebeat/harvester/encoding"
)

// Encode reader produces lines by reading lines from an io.Reader
// through a decoder converting the reader it's encoding to utf-8.
type Encode struct {
	reader *Line
}

// NewEncode creates a new Encode reader from input reader by applying
// the given codec.
func NewEncode(
	in io.ReadCloser,
	codec encoding.Encoding,
	bufferSize int,
) (*Encode, error) {
	r, err := NewLine(in, codec, bufferSize)
	return &Encode{r}, err
}

// Next reads the next line from it's initial io.Reader
func (r *Encode) Next() (Message, error) {
	c, sz, err := r.reader.Next()
	return Message{
		Ts:      time.Now(),
		Content: c,
		Bytes:   sz,
	}, err
}

func (r *Encode) Close() error {
	return r.reader.Close()
}

package strip_newline

import (
	"github.com/elastic/beats/filebeat/reader"
	"github.com/elastic/beats/libbeat/common"
)

// StripNewline reader removes the last trailing newline characters from
// read lines.
type StripNewline struct {
	reader reader.Reader
}

// New creates a new line reader stripping the last tailing newline.
func New(r reader.Reader) *StripNewline {
	return &StripNewline{r}
}

// Next returns the next line.
func (p *StripNewline) Next() (reader.Message, error) {
	message, err := p.reader.Next()
	if err != nil {
		return message, err
	}

	L := message.Content
	message.Content = L[:len(L)-lineEndingChars(L)]

	return message, err
}

// isLine checks if the given byte array is a line, means has a line ending \n
func isLine(l []byte) bool {
	return l != nil && len(l) > 0 && l[len(l)-1] == '\n'
}

// lineEndingChars returns the number of line ending chars the given by array has
// In case of Unix/Linux files, it is -1, in case of Windows mostly -2
func lineEndingChars(l []byte) int {
	if !isLine(l) {
		return 0
	}

	if len(l) > 1 && l[len(l)-2] == '\r' {
		return 2
	}
	return 1
}

func (r *StripNewline) GetState() common.MapStr {
	return r.reader.GetState()
}

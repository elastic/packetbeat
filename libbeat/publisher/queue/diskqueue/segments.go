// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package diskqueue

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Every data frame read from the queue is assigned a unique sequential
// integer, which is used to keep track of which frames have been
// acknowledged.
// This id is not stable between restarts; the value 0 is always assigned
// to the oldest remaining frame on startup.
type frameID uint64

// segmentOffset is a byte index into the segment's data region.
type segmentOffset uint64

// The metadata for a single segment file.
type queueSegment struct {
	// A segment id is globally unique within its originating queue.
	id segmentID

	// The settings for the queue that created this segment. Used for locating
	// the queue file on disk and determining its checksum behavior.
	queueSettings *Settings

	// Whether the file for this segment exists on disk yet. If it does
	// not, then calling getWriter() will create it and return a writer
	// positioned at the start of the data region.
	created bool

	// The size in bytes of the segment file on disk. This is updated when
	// the segment is written to, and should always correspond to the end of
	// a complete data frame.
	size uint64

	// The header metadata for this segment file. This field is nil if the
	// segment has not yet been opened for reading. It should only be
	// accessed by the reader loop.
	header *segmentHeader

	// The number of frames read from this segment during this session. This
	// does not necessarily equal the number of frames in the segment, even
	// after reading is complete, since the segment may have been partially
	// read during a previous session.
	//
	// Used to count how many frames still need to be acknowledged by consumers.
	framesRead int64

	// If this segment is being written or read, then reader / writer
	// contain the respective file handles. To get a valid reader / writer for
	// a segment that may not yet be open, call getReader / getWriter instead.
	reader *os.File
	writer *os.File
}

type segmentHeader struct {
	version      uint32
	checksumType ChecksumType
}

// segmentReader is a wrapper around io.Reader that provides helpers and
// metadata for decoding segment files.
/*type segmentReader struct {
	// The segment this reader was generated from.
	segment *queueSegment

	// The underlying data reader.
	raw io.Reader

	// The current byte offset of the reader within the file.
	curPosition segmentOffset

	// The position at which this reader should stop reading. This is often
	// the end of the file, but it may be earlier when the queue is reading
	// and writing to the same segment.
	endPosition segmentOffset

	// The checksumType field from this segment file's header.
	checksumType ChecksumType
}*/

/*type segmentWriter struct {
	segment  *queueSegment
	file     *os.File
	position int64
}*/

// ChecksumType specifies what checksum algorithm the queue should use to
// verify its data frames.
type ChecksumType int

// ChecksumTypeNone: Don't compute or verify checksums.
// ChecksumTypeCRC32: Compute the checksum with the Go standard library's
//   "hash/crc32" package.
const (
	ChecksumTypeNone = iota

	ChecksumTypeCRC32
)

// Each data frame has a 32-bit lengths and 1 32-bit checksum
// in the header, and a duplicate 32-bit length in the footer.
const frameHeaderSize = 8
const frameFooterSize = 4
const frameMetadataSize = frameHeaderSize + frameFooterSize

// Each segment header has a 32-bit version and a 32-bit checksum type.
const segmentHeaderSize = 8

// Sort order: we store loaded segments in ascending order by their id.
type bySegmentID []*queueSegment

func (s bySegmentID) Len() int           { return len(s) }
func (s bySegmentID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bySegmentID) Less(i, j int) bool { return s[i].size < s[j].size }

func scanExistingSegments(path string) ([]*queueSegment, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read queue directory '%s': %w", path, err)
	}

	segments := []*queueSegment{}
	for _, file := range files {
		if file.Size() <= segmentHeaderSize {
			// Ignore segments that don't have at least some data beyond the
			// header (this will always be true of segments we write unless there
			// is an error).
			continue
		}
		components := strings.Split(file.Name(), ".")
		if len(components) == 2 && strings.ToLower(components[1]) == "seg" {
			// Parse the id as base-10 64-bit unsigned int. We ignore file names that
			// don't match the "[uint64].seg" pattern.
			if id, err := strconv.ParseUint(components[0], 10, 64); err == nil {
				segments = append(segments,
					&queueSegment{
						id:      segmentID(id),
						created: true,
						size:    uint64(file.Size()),
					})
			}
		}
	}
	sort.Sort(bySegmentID(segments))
	return segments, nil
}

// Should only be called from the reader loop.
func (segment *queueSegment) getReader() (io.ReadCloser, error) {
	if segment.reader != nil {
		return segment.reader, nil
	}
	path := segment.queueSettings.segmentPath(segment.id)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(
			"Couldn't open segment %d: %w", segment.id, err)
	}
	//reader := bufio.NewReader(file)
	header, err := readSegmentHeader(file)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read segment header: %w", err)
	}
	segment.header = header

	return file, nil
}

// Should only be called from the writer loop.
func (segment *queueSegment) getWriter() (io.WriteCloser, error) {
	if segment.writer != nil {
		return segment.writer, nil
	}
	path := segment.queueSettings.segmentPath(segment.id)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	segment.writer = file
	header := &segmentHeader{
		version:      0,
		checksumType: segment.queueSettings.ChecksumType,
	}
	err = writeSegmentHeader(file, header)
	// TODO: write header
	return file, nil
}

func readSegmentHeader(in io.Reader) (*segmentHeader, error) {
	header := segmentHeader{}
	if header.version != 0 {
		return nil, fmt.Errorf("Unrecognized schema version %d", header.version)
	}
	panic("TODO: not implemented")
	//return nil, nil
}

func writeSegmentHeader(out io.Writer, header *segmentHeader) error {
	panic("TODO: not implemented")
}

// The number of bytes occupied by all the queue's segment files. This
// should only be called from the core loop.
func (segments *diskQueueSegments) sizeOnDisk() uint64 {
	total := uint64(0)
	if segments.writing != nil {
		total += segments.writing.size
	}
	for _, segment := range segments.reading {
		total += segment.size
	}
	for _, segment := range segments.acking {
		total += segment.size
	}
	for _, segment := range segments.acked {
		total += segment.size
	}
	return total
}

// nextDataFrame returns the bytes of the next data frame, or an error if the
// frame couldn't be read. If an error is returned, the caller should log it
// and drop the containing segment. A nil return value with no error means
// there are no frames to read.
/*func (reader *segmentReader) nextDataFrame() ([]byte, error) {
	if reader.curPosition >= reader.endPosition {
		return nil, nil
	}
	var frameLength uint32
	err := binary.Read(reader.raw, binary.LittleEndian, &frameLength)
	if err != nil {
		return nil, fmt.Errorf(
			"Disk queue couldn't read next frame length: %w", err)
	}

	// Bounds checking to make sure we can read this frame.
	if reader.curPosition+segmentOffset(frameLength) > reader.endPosition {
		// This frame extends past the end of our data region, which
		// should never happen unless there is data corruption.
		return nil, fmt.Errorf(
			"Data frame length (%d) exceeds remaining data (%d)",
			frameLength, reader.endPosition-reader.curPosition)
	}
	if frameLength <= frameMetadataSize {
		// Actual enqueued data must have positive length
		return nil, fmt.Errorf(
			"Data frame with no data (length %d)", frameLength)
	}

	// Read the actual frame data
	dataLength := frameLength - frameMetadataSize
	data := make([]byte, dataLength)
	_, err = io.ReadFull(reader.raw, data)
	if err != nil {
		return nil, fmt.Errorf(
			"Couldn't read data frame from disk: %w", err)
	}

	// Read the footer (length + checksum)
	var duplicateLength uint32
	err = binary.Read(reader.raw, binary.LittleEndian, &duplicateLength)
	if err != nil {
		return nil, fmt.Errorf(
			"Disk queue couldn't read trailing frame length: %w", err)
	}
	if duplicateLength != frameLength {
		return nil, fmt.Errorf(
			"Disk queue: inconsistent frame length (%d vs %d)",
			frameLength, duplicateLength)
	}

	// Validate the checksum
	var checksum uint32
	err = binary.Read(reader.raw, binary.LittleEndian, &checksum)
	if err != nil {
		return nil, fmt.Errorf(
			"Disk queue couldn't read data frame's checksum: %w", err)
	}
	if computeChecksum(data, reader.checksumType) != checksum {
		return nil, fmt.Errorf("Disk queue: bad data frame checksum")
	}

	reader.curPosition += segmentOffset(frameLength)
	return data, nil
}*/

// returns the number of indices by which ackedUpTo was advanced.
func (dq *diskQueue) ack(frame frameID) int {
	dq.ackLock.Lock()
	defer dq.ackLock.Unlock()
	dq.acked[frame] = true
	ackedCount := 0
	for ; dq.acked[dq.ackedUpTo]; dq.ackedUpTo++ {
		delete(dq.acked, dq.ackedUpTo)
		ackedCount++
	}
	return ackedCount
}

func computeChecksum(data []byte, checksumType ChecksumType) uint32 {
	switch checksumType {
	case ChecksumTypeNone:
		return 0
	case ChecksumTypeCRC32:
		hash := crc32.NewIEEE()
		frameLength := uint32(len(data) + frameMetadataSize)
		binary.Write(hash, binary.LittleEndian, &frameLength)
		hash.Write(data)
		return hash.Sum32()
	default:
		panic("segmentReader: invalid checksum type")
	}
}

/*
func (dq *diskQueue) segmentReaderForPosition(
	pos bufferPosition,
) (*segmentReader, error) {
	panic("TODO: not implemented")
}
*/

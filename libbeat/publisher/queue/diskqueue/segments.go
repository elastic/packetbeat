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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/elastic/elastic-agent-libs/logp"
)

// diskQueueSegments encapsulates segment-related queue metadata.
type diskQueueSegments struct {

	// A list of the segments that have not yet been completely written, sorted
	// by increasing segment ID. When the first entry has been completely
	// written, it is removed from this list and appended to reading.
	//
	// If the reading list is empty, the queue may read from a segment that is
	// still being written, but it will always be writing[0], since later
	// entries do not yet exist on disk.
	writing []*queueSegment

	// A list of the segments that have been completely written but have
	// not yet been completely read, sorted by increasing segment ID. When the
	// first entry has been completely read, it is removed from this list and
	// appended to acking.
	reading []*queueSegment

	// A list of the segments that have been completely read but have not yet
	// been completely acknowledged, sorted by increasing segment ID. When the
	// first entry has been completely acknowledged, it is removed from this
	// list and appended to acked.
	acking []*queueSegment

	// A list of the segments that have been completely read and acknowledged
	// and are ready to be deleted. When a segment is successfully deleted, it
	// is removed from this list and discarded.
	acked []*queueSegment

	// The next sequential unused segment ID. This is what will be assigned
	// to the next queueSegment we create.
	nextID segmentID

	// writingSegmentSize tracks the expected on-disk size of the current write
	// segment after all scheduled frames have finished writing. This is used in
	// diskQueue.enqueueWriteFrame to detect when to roll over to a new segment.
	writingSegmentSize uint64

	// nextReadFrameID is the first frame ID in the current or pending
	// read request.
	nextReadFrameID frameID

	// nextReadPosition is the next absolute byte offset on disk that should be
	// read from the current read segment. The current read segment is either
	// reading[0], or writing[0] if the reading list is empty.
	// A value of 0 means the first frame in the segment (the
	// exact byte position depends on the file header, see
	// diskQueue.maybeReadPending()).
	nextReadPosition uint64
}

// segmentID is a unique persistent integer id assigned to each created
// segment in ascending order.
type segmentID uint64

// The metadata for a single segment file.
type queueSegment struct {
	// A segment id is globally unique within its originating queue.
	id segmentID

	// schemaVersion is used to determine on disk format, data serialization,
	// and encryption
	schemaVersion uint32

	// The number of bytes occupied by this segment on-disk, as of the most
	// recent completed writerLoop request.
	byteCount uint64

	// The ID of the first frame that was / will be read from this segment.
	// This field is only valid after a read request has been sent for
	// this segment. (Currently it is only used to handle consumer ACKs,
	// which can only happen after reading has begun on the segment.)
	firstFrameID frameID

	// The number of frames contained in this segment on disk, as of the
	// most recent completed writerLoop request (this does not include
	// segments which are merely scheduled to be written).
	frameCount uint32

	// The number of frames read from this segment during this session. This
	// does not necessarily equal the number of frames in the segment, even
	// after reading is complete, since the segment may have been partially
	// read during a previous session.
	//
	// Used to count how many frames still need to be acknowledged by consumers.
	framesRead uint64
}

type segmentHeader struct {
	// The schema version for this segment file
	version uint32

	// If the segment file has been completely written, this field contains
	// the number of data frames, which is used to track the number of
	// pending events left in the queue from previous sessions.
	// If the segment file has not been completely written, this field is zero.
	// Only present in schema version >= 1.
	frameCount uint32
}

type WriteCloseSyncer interface {
	io.Writer
	io.Closer
	Sync() error
}

// segmentHeaderSizeV0 schemaVersion 0 header size, uint32 for version
const segmentHeaderSizeV0 = 4

// segmentHeaderSizeV1 schemaVersion 1 header size, uint32 for version, uint32 for count
const segmentHeaderSizeV1 = 8

// segmentHeaderSizeV2 schemaVersion 2 header size uint32 for version
const segmentHeaderSizeV2 = 4

// maxSegmentVersion is the highest supported version
const maxSegmentVersion = 2

// Sort order: we store loaded segments in ascending order by their id.
type bySegmentID []*queueSegment

func (s bySegmentID) Len() int           { return len(s) }
func (s bySegmentID) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s bySegmentID) Less(i, j int) bool { return s[i].id < s[j].id }

// Scan the given path for segment files, and return them in a list
// ordered by segment id.
func scanExistingSegments(logger *logp.Logger, pathStr string) ([]*queueSegment, error) {
	files, err := ioutil.ReadDir(pathStr)
	if err != nil {
		return nil, fmt.Errorf("couldn't read queue directory '%s': %w", pathStr, err)
	}

	segments := []*queueSegment{}
	for _, file := range files {
		components := strings.Split(file.Name(), ".")
		if len(components) == 2 && strings.ToLower(components[1]) == "seg" {
			// Parse the id as base-10 64-bit unsigned int. We ignore file names that
			// don't match the "[uint64].seg" pattern.
			if id, err := strconv.ParseUint(components[0], 10, 64); err == nil {
				fullPath := path.Join(pathStr, file.Name())
				header, err := readSegmentHeaderWithFrameCount(fullPath)
				if header == nil {
					logger.Errorf("couldn't load segment file '%v': %v", fullPath, err)
					continue
				}
				// If we get an error but still got a valid header back, then we
				// were able to read at least some frames, so we keep this segment
				// but issue a warning.
				if err != nil {
					logger.Warnf(
						"error loading segment file '%v', data may be incomplete: %v",
						fullPath, err)
				}
				segments = append(segments, &queueSegment{
					id:            segmentID(id),
					schemaVersion: header.version,
					frameCount:    header.frameCount,
					byteCount:     uint64(file.Size()),
				})
			}
		}
	}
	sort.Sort(bySegmentID(segments))
	return segments, nil
}

// headerSize returns the logical size ("logical" because it may not have
// been written to disk yet) of this segment file's header region. The
// segment's first data frame begins immediately after the header.
func (segment *queueSegment) headerSize() uint64 {
	switch segment.schemaVersion {
	case 0:
		return segmentHeaderSizeV0
	case 1:
		return segmentHeaderSizeV1
	case 2:
		return segmentHeaderSizeV2
	default:
		return uint64(0)
	}
}

// The initial release of the disk queue used JSON to encode events
// on disk. Since then, we have switched to CBOR to address issues
// with encoding multi-byte characters, and for lower encoding
// overhead.
func (segment *queueSegment) shouldUseJSON() bool {
	return segment.schemaVersion == 0
}

// Should only be called from the reader loop. If successful, returns an open
// file handle positioned at the beginning of the segment's data region.
func (segment *queueSegment) getReader(queueSettings Settings) (*segmentReader, error) {
	path := queueSettings.segmentPath(segment.id)
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(
			"couldn't open segment %d: %w", segment.id, err)
	}
	err = binary.Read(file, binary.LittleEndian, &segment.schemaVersion)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("couldn't read segment version: %w", err)
	}

	if segment.schemaVersion > maxSegmentVersion {
		file.Close()
		return nil, fmt.Errorf("unknown segment version %d: %w", segment.schemaVersion, err)
	}

	if segment.schemaVersion == 1 {
		err = binary.Read(file, binary.LittleEndian, &segment.frameCount)
		if err != nil {
			file.Close()
			return nil, fmt.Errorf("couldn't read segment frame count: %w", err)
		}
	}

	sr := &segmentReader{}
	sr.src = file
	sr.version = segment.schemaVersion

	if sr.version == 0 || sr.version == 1 {
		return sr, nil
	}

	if sr.version == 2 {
		sr.er, err = NewEncryptionReader(sr.src, queueSettings.EncryptionKey)
		if err != nil {
			sr.src.Close()
			return nil, fmt.Errorf("couldn't create encryption reader: %w", err)
		}
	}

	return sr, nil

}

// Should only be called from the writer loop.
func (segment *queueSegment) getWriter(queueSettings Settings) (*segmentWriter, error) {
	path := queueSettings.segmentPath(segment.id)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}

	sw := &segmentWriter{}
	sw.dst = file

	segment.schemaVersion = queueSettings.SchemaVersion
	sw.version = queueSettings.SchemaVersion

	if err := sw.WriteHeader(); err != nil {
		return nil, err
	}

	if sw.version == 0 || sw.version == 1 {
		return sw, nil
	}

	if sw.version == 2 {
		sw.ew, err = NewEncryptionWriter(sw.dst, queueSettings.EncryptionKey)
		if err != nil {
			sw.dst.Close()
			return nil, fmt.Errorf("couldn't create encryption writer: %w", err)
		}
	}

	return sw, nil
}

// getWriterWithRetry tries to create a file handle for writing via
// queueSegment.getWriter. On error, it retries as long as the given
// retry callback returns true. This is used for timed retries when
// creating a queue segment from the writer loop.
func (segment *queueSegment) getWriterWithRetry(
	queueSettings Settings, retry func(err error, firstTime bool) bool,
) (*segmentWriter, error) {
	firstTime := true
	file, err := segment.getWriter(queueSettings)
	for err != nil && retry(err, firstTime) {
		// Set firstTime to false so the retry callback can perform backoff
		// etc if needed.
		firstTime = false

		// Try again
		file, err = segment.getWriter(queueSettings)
	}
	return file, err
}

// readSegmentHeaderWithFrameCount reads the header from the beginning
// of the file at the given path. If the header's frameCount is 0
// (whether because it is from an old version or because the segment
// file was not closed cleanly), it attempts to calculate it manually
// by scanning the file, and returns a struct with the "correct"
// frame count.
func readSegmentHeaderWithFrameCount(path string) (*segmentHeader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf(
			"couldn't open segment file '%s': %w", path, err)
	}
	defer file.Close()
	// Wrap the handle to retry non-fatal errors and always return the full
	// requested data length if possible, then read the raw header.
	reader := autoRetryReader{file}
	header, err := readSegmentHeader(reader)
	if err != nil {
		return nil, err
	}
	// If the header has a positive frame count then there is
	// no more work to do, so return immediately.
	if header.frameCount > 0 {
		return header, nil
	}
	// If we made it here, we loaded a valid header but the frame count is
	// zero, so we need to check it with a manual scan. This can
	// only happen in one of two uncommon situations:
	// - The segment was created by an old version that didn't track frame count
	// - The segment file was not closed cleanly during the previous session
	//   and still has the placeholder value of 0.
	// In either case, the right thing to do is to scan the file
	// and fill in the frame count manually.
	for {
		var frameLength uint32
		err = binary.Read(reader, binary.LittleEndian, &frameLength)
		if err != nil {
			// EOF at a frame boundary means we successfully scanned all frames.
			if errors.Is(err, io.EOF) && header.frameCount > 0 {
				return header, nil
			}
			// All other errors mean we are done scanning, exit the loop.
			break
		}
		// Length is encoded in both the first and last four bytes of a frame. To
		// detect truncated / corrupted frames, seek to the last four bytes of
		// the current frame to make sure the trailing length matches before
		// advancing to the next frame (otherwise we might accept an impossible
		// length).
		_, err = file.Seek(int64(frameLength-8), io.SeekCurrent)
		if err != nil {
			break
		}
		var duplicateLength uint32
		err = binary.Read(reader, binary.LittleEndian, &duplicateLength)
		if err != nil {
			break
		}
		if frameLength != duplicateLength {
			err = fmt.Errorf(
				"mismatched frame length: %v vs %v", frameLength, duplicateLength)
			break
		}

		header.frameCount++
	}
	// If we ended up here instead of returning directly, then
	// we encountered an error. We still return a valid header as
	// long as we successfully scanned at least one frame first.
	if header.frameCount > 0 {
		return header, err
	}
	return nil, err
}

// readSegmentHeader decodes a raw header from the given reader and
// returns it as a struct.
func readSegmentHeader(in io.Reader) (*segmentHeader, error) {
	header := &segmentHeader{}
	err := binary.Read(in, binary.LittleEndian, &header.version)
	if err != nil {
		return nil, err
	}
	switch header.version {
	case 0: // do nothing
	case 1:
		err = binary.Read(in, binary.LittleEndian, &header.frameCount)
		if err != nil {
			return nil, err
		}
	case 2: // do nothing
	default:
		return nil, fmt.Errorf("unrecognized schema version %d", header.version)
	}
	return header, nil
}

// The number of bytes occupied by all the queue's segment files. This
// should only be called from the core loop.
func (segments *diskQueueSegments) sizeOnDisk() uint64 {
	total := uint64(0)
	for _, segment := range segments.writing {
		total += segment.byteCount
	}
	for _, segment := range segments.reading {
		total += segment.byteCount
	}
	for _, segment := range segments.acking {
		total += segment.byteCount
	}
	for _, segment := range segments.acked {
		total += segment.byteCount
	}
	return total
}

type segmentReader struct {
	src     io.ReadSeekCloser
	er      *EncryptionReader
	version uint32
}

func (r *segmentReader) Read(p []byte) (int, error) {
	switch r.version {
	case 0, 1:
		return r.src.Read(p)
	case 2:
		return r.er.Read(p)
	default:
		return 0, fmt.Errorf("unsupported schema version: %d", r.version)
	}
}

func (r *segmentReader) Close() error {
	switch r.version {
	case 0, 1:
		return r.src.Close()
	case 2:
		return r.er.Close()
	default:
		return fmt.Errorf("unsupported schema version: %d", r.version)
	}
}

func (r *segmentReader) Seek(offset int64, whence int) (int64, error) {
	switch r.version {
	case 0, 1:
		return r.src.Seek(offset, whence)
	case 2:
		//can't seek before segment header
		if (offset + int64(whence)) < segmentHeaderSizeV2 {
			return 0, fmt.Errorf("illegal seek offset %d, whence %d", offset, whence)
		}
		if _, err := r.src.Seek(segmentHeaderSizeV2, io.SeekStart); err != nil {
			return 0, err
		}
		if err := r.er.Reset(); err != nil {
			return 0, err
		}
		written, err := io.CopyN(io.Discard, r.er, (offset+int64(whence))-segmentHeaderSizeV2)
		return written + segmentHeaderSizeV2, err
	default:
		return 0, fmt.Errorf("unsupported schema version: %d", r.version)
	}
}

type segmentWriter struct {
	dst     *os.File
	ew      *EncryptionWriter
	version uint32
}

func (w *segmentWriter) Write(p []byte) (int, error) {
	switch w.version {
	case 0, 1:
		return w.dst.Write(p)
	case 2:
		return w.ew.Write(p)
	default:
		return 0, fmt.Errorf("unsupported schema version: %d", w.version)
	}
}

func (w *segmentWriter) Close() error {
	switch w.version {
	case 0, 1:
		return w.dst.Close()
	case 2:
		return w.ew.Close()
	default:
		return fmt.Errorf("unsupported schema version: %d", w.version)
	}
}

func (w *segmentWriter) Seek(offset int64, whence int) (int64, error) {
	switch w.version {
	case 0, 1:
		return w.dst.Seek(offset, whence)
	default:
		return 0, fmt.Errorf("unsupported schema version: %d", w.version)
	}
}

func (w *segmentWriter) Sync() error {
	switch w.version {
	case 0, 1:
		return w.dst.Sync()
	case 2:
		return w.ew.Sync()
	default:
		return fmt.Errorf("unsupported schema version: %d", w.version)
	}
}

func (w *segmentWriter) WriteHeader() error {
	if _, err := w.dst.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := binary.Write(w.dst, binary.LittleEndian, w.version); err != nil {
		return err
	}

	// Version 1 has count
	if w.version == 1 {
		if err := binary.Write(w.dst, binary.LittleEndian, uint32(0)); err != nil {
			return err
		}
	}

	return nil
}

func (w *segmentWriter) UpdateCount(count uint32) error {
	// Only Version 1 stores count
	if w.version != 1 {
		return nil
	}

	// Seek to count on disk
	if _, err := w.dst.Seek(4, io.SeekStart); err != nil {
		return err
	}

	return binary.Write(w.dst, binary.LittleEndian, count)
}

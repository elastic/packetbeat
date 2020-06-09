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

func (dq *diskQueue) readerLoop() {
	curFrameID := frameID(0)
	for {
		dq.frameWrittenCond.Wait()
		reader, errs := dq.nextSegmentReader()
		for _, err := range errs {
			// Errors encountered while reading should be logged.
			dq.settings.Logger.Error(err)
		}
		if reader == nil {
			// We couldn't find a readable segment, wait for a new
			// data frame to be written.
			dq.frameWrittenCond.Wait()
			if dq.closed.Load() {
				// The queue has been closed, shut down.
				// TODO: cleanup (write the final read position)
				return
			}
			continue
		}
		// If we made it here, we have a nonempty reader and we want
		// to send all its frames to dq.outChan.
		framesRead := int64(0)
		for {
			bytes, err := reader.nextDataFrame()
			if err != nil {
				// An error at this stage generally means there has been
				// data corruption. For now, in this case we log the
				// error, then discard any remaining frames. When all
				// successfully read frames have been acknowledged, we
				// delete the underlying file.
				break
			}
			if bytes == nil {
				// If bytes is nil with no error, we've reached the end
				// of this segmentReader. Update the segment's frame count.
				break
			}
			framesRead++
			output := diskQueueOutput{
				data:    bytes,
				segment: reader.segment,
				frame:   curFrameID,
			}
			select {
			case dq.outChan <- output:
				curFrameID++
			case <-dq.done:
			}
		}
		reader.segment.framesRead += framesRead
	}
}

func (dq *diskQueue) writerLoop() {
	for {
		var frameBytes []byte
		select {
		case frameBytes = <-dq.inChan:

		case <-dq.done:
			break
		}
	}
}

// This is only called from the writer loop.
func (dq *diskQueue) writeFrameData(bytes []byte) error {
	frameSize := len(bytes) + frameMetadataSize

	// We can check segment size without holding a lock because the writer
	// loop (i.e. the current thread) is the only place that modifies it
	// after segmentFile creation.
	if dq.writingSegment == nil ||
		dq.writingSegment.size + frameSize > dq.settings.MaxSegmentSize {
		// There is no writing segment, or the incoming data frame doesn't
		// fit in the current one, so we need to create a new one.
	}
	
	.capacity() < frameSize {

	} != nil && dq.writingSegment.

	// while (free bytes) < frameSize {
	// block
	// }

	if dq.writingSegment == nil {
		// There is no current output segment, create a new one.

	}

	return nil
}
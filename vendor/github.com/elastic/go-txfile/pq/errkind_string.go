// Code generated by "stringer -type=ErrKind -linecomment=true"; DO NOT EDIT.

package pq

import "strconv"

const _ErrKind_name = "no errorfailed to initialize queueinvalid parameterinvalid page sizeinvalid queue configqueue is already closedreader is already closedwriter is already closedno queue rootqueue root is invalidunsupported queue versioninvalid ack on empty queuetoo many events ackedfailed to seek to next pagefailed to read page"

var _ErrKind_index = [...]uint16{0, 8, 34, 51, 68, 88, 111, 135, 159, 172, 193, 218, 244, 265, 292, 311}

func (i ErrKind) String() string {
	if i < 0 || i >= ErrKind(len(_ErrKind_index)-1) {
		return "ErrKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrKind_name[_ErrKind_index[i]:_ErrKind_index[i+1]]
}

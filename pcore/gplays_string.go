// Code generated by "stringer -type=GPLays"; DO NOT EDIT.

package pcore

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GPeOut-0]
	_ = x[GPeIn-1]
	_ = x[GPeTA-2]
	_ = x[GPi-3]
	_ = x[GPLaysN-4]
}

const _GPLays_name = "GPeOutGPeInGPeTAGPiGPLaysN"

var _GPLays_index = [...]uint8{0, 6, 11, 16, 19, 26}

func (i GPLays) String() string {
	if i < 0 || i >= GPLays(len(_GPLays_index)-1) {
		return "GPLays(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GPLays_name[_GPLays_index[i]:_GPLays_index[i+1]]
}

func (i *GPLays) FromString(s string) error {
	for j := 0; j < len(_GPLays_index)-1; j++ {
		if s == _GPLays_name[_GPLays_index[j]:_GPLays_index[j+1]] {
			*i = GPLays(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: GPLays")
}

// Code generated by "stringer -type=Valences"; DO NOT EDIT.

package pbwm

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Appetitive-0]
	_ = x[Aversive-1]
	_ = x[ValencesN-2]
}

const _Valences_name = "AppetitiveAversiveValencesN"

var _Valences_index = [...]uint8{0, 10, 18, 27}

func (i Valences) String() string {
	if i < 0 || i >= Valences(len(_Valences_index)-1) {
		return "Valences(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Valences_name[_Valences_index[i]:_Valences_index[i+1]]
}

func (i *Valences) FromString(s string) error {
	for j := 0; j < len(_Valences_index)-1; j++ {
		if s == _Valences_name[_Valences_index[j]:_Valences_index[j+1]] {
			*i = Valences(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Valences")
}

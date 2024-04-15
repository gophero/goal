package errorx

import (
	"fmt"
	"reflect"
)

var (
	ServerBusy  = NewPreferredErrf("sever is busy")
	BadReq      = NewPreferredErrf("bad request")
	NoAuth      = NewPreferredErrf("unauthorized access")
	FrequentReq = NewPreferredErrf("frequent requests")
)

// PreferredError 首选 error
type PreferredError struct {
	error
}

func (e *PreferredError) Error() string {
	return e.error.Error()
}

func NewPreferredErr(err error) error {
	return &PreferredError{error: err}
}

func NewPreferredErrf(format string, args ...any) error {
	return &PreferredError{error: fmt.Errorf(format, args...)}
}

func IsPreferred(err error) bool {
	if err == nil {
		return false
	}
	return reflect.TypeOf(err).AssignableTo(reflect.TypeOf(&PreferredError{}))
}

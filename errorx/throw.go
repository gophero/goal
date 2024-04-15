package errorx

import (
	"errors"
	"fmt"
)

// Throw method will panic the given err if it's not nil.
func Throw(err error) {
	if err != nil {
		panic(err)
	}
}

// Throwf method will panic the given err if it's not nil, and using custom text as panic message.
func Throwf(err error, format string, v ...any) {
	if err != nil {
		panic(fmt.Sprintf(format, v...))
	}
}

// New method will create an error using the given string and arguments.
func New(s string, v ...any) error {
	if v == nil {
		return errors.New(s)
	}
	return errors.New(fmt.Sprintf(s, v...))
}

// Throwv method expect there must no error, i.e. err argument is nil, if it is not, then
// panic it, otherwise return the first argument.
func Throwv[T any](v T, err error) T {
	Throw(err)
	return v
}

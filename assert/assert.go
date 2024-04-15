package assert

import (
	"fmt"
	"reflect"
	"strings"
)

// basic methods

func Require(cond bool, format string, v ...any) {
	if !cond {
		panic(fmt.Sprintf(format, v...))
	}
}

// convenient methods

func True(b bool) {
	Require(b, "assert failed, expects [%t] but found [%t]", true, b)
}

func Nil(t any) {
	Require(t == nil, "assert failed, expects [nil] but found [not nil]: %v", t)
}

func NoneNil(t any) {
	Require(t != nil, "assert failed, expects [not nil] but found [nil]")
}

func Blank(s string) {
	s = strings.TrimSpace(s)
	Require(s == "", "assert failed, expects [\"\"] but found [%s]", s)
}

func NotBlank(s string) {
	s = strings.TrimSpace(s)
	Require(s != "", "assert failed, expects [not empty] but found [\"\"]")
}

func HasElems(c any) {
	typ := reflect.TypeOf(c).Kind()
	if typ == reflect.Array || typ == reflect.Map || typ == reflect.Slice || typ == reflect.Chan {
		n := reflect.ValueOf(c).Len()
		Require(c != nil && n > 0, "assert failed, expects collection is none nil and has elements")
	}
}

func Equals(t1 any, t2 any) {
	Require(t1 != t2, "assert failed, expecting t1 equals t2 but not")
}

func DeepEquals(t1 any, t2 any) {
	Require(reflect.DeepEqual(t1, t2), "assert failed, expects %v deep equals %v, but not", t1, t2)
}

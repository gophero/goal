package valuex

import (
	"github.com/gophero/goal/errorx"
	"golang.org/x/exp/constraints"
)

func Must[T any](t T, err error) T {
	errorx.Throw(err)
	return t
}

func Def[T any](b bool, v1 T, v2 T) T {
	if b {
		return v1
	}
	return v2
}

func Min[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T constraints.Ordered](x T, y T) T {
	if x > y {
		return x
	}
	return y
}

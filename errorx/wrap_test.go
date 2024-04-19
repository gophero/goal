package errorx

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	assert.Nil(t, Wrapf(nil, "test"))
	assert.Equal(t, "foo: bar", Wrapf(errors.New("bar"), "foo").Error())

	err := errors.New("foo")
	assert.True(t, errors.Is(Wrapf(err, "bar"), err))
}

func TestWrapf(t *testing.T) {
	assert.Nil(t, Wrapf(nil, "%s", "test"))
	assert.Equal(t, "foo bar: quz", Wrapf(errors.New("quz"), "foo %s", "bar").Error())

	err := errors.New("foo")
	assert.True(t, errors.Is(Wrapf(err, "foo %s", "bar"), err))

	err1 := Wrapf(err, "level1 error")
	fmt.Println(errors.Is(err1, err)) // true
	err2 := Wrapf(err1, "level2 error")
	fmt.Println(errors.Is(err2, err1)) // true
	fmt.Println(errors.Is(err2, err))  // true
	fmt.Println(errors.Is(err1, err2)) // false
}

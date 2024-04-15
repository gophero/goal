package errorx

import (
	"errors"
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
}

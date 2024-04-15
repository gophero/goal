package errorx_test

import (
	"context"
	"github.com/gophero/goal/errorx"
	"github.com/gophero/goal/testx"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRescue(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer errorx.Recover(testx.NewLog(), func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}

func TestRescueCtx(t *testing.T) {
	var count int32
	assert.NotPanics(t, func() {
		defer errorx.RecoverCtx(context.Background(), testx.NewLog(), func() {
			atomic.AddInt32(&count, 2)
		}, func() {
			atomic.AddInt32(&count, 3)
		})

		panic("hello")
	})
	assert.Equal(t, int32(5), atomic.LoadInt32(&count))
}

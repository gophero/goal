package logx_test

import (
	"fmt"
	"github.com/gophero/goal/logx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefLogger(t *testing.T) {
	go func() {
		err := recover()
		fmt.Printf("need an error: %v", err)
		assert.True(t, err != nil)
		logx.Default.Debug("this is a debug log")
		logx.Default.Info("this is an info log")
		logx.Default.Warn("this is an warning log")
		logx.Default.Error("this is an error log")
		logx.Default.Fatal("this is a fatal log")
	}()
}

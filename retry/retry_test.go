package retry_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gophero/goal/random"
	"github.com/gophero/goal/retry"
	"github.com/stretchr/testify/assert"
)

var (
	succ = func() error { return nil }
	fail = func() error { return errors.New("make an error") }
)

func TestAlwaysSuccess(t *testing.T) {
	for i := 0; i < 100; i++ {
		if !retry.Do(succ, 10) {
			t.Errorf("test failed")
			break
		}
	}
}

func TestAlwaysFail(t *testing.T) {
	for i := 0; i < 100; i++ {
		if retry.Do(fail, 0) {
			t.Errorf("test failed")
			break
		}
	}
}

func TestDo(t *testing.T) {
	var hasErr bool
	randFunc := func() error {
		if random.Int(2) == 0 {
			fmt.Println("retry fail...")
			hasErr = true
			return errors.New("make an error")
		} else {
			fmt.Println("retry success...")
			hasErr = false
			return nil
		}
	}
	b := retry.Do(randFunc, 4)
	assert.True(t, b, !hasErr)
}

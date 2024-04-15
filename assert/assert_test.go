package assert_test

import (
	"github.com/gophero/goal/assert"
	"github.com/gophero/goal/testx"
	"testing"
)

func TestRequire(t *testing.T) {
	logger := testx.Wrap(t)
	logger.Title("test assert.Require method")

	logger.Case("give a true condition, should not panic")
	func() {
		defer func() {
			err := recover()
			if err != nil {
				logger.Fail("should not get an error: %v", err)
			} else {
				logger.Pass("should not get an error")
			}
		}()
		assert.Require(true, "not match the condition")
	}()

	logger.Case("give a false condition, should panic and get an error")
	func() {
		msg := "not match the condition"
		defer func() {
			err := recover()
			if err == nil {
				logger.Fail("should get an error with message: " + msg)
			} else {
				logger.Pass("should get an error: %v", err)
			}
		}()
		assert.Require(false, "not match the condition")
	}()
}

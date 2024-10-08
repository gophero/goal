package iox_test

import (
	"github.com/gophero/goal/iox"
	"github.com/gophero/goal/testx"
	"testing"
)

func TestExistsFile(t *testing.T) {
	lg := testx.Wrap(t)

	lg.Case("give an existing file")
	f := "./file_test.go"
	lg.Require(iox.File.Exists(f), "should exist")

	lg.Case("give an existing dir, but is not a file")
	f = "."
	lg.Require(!iox.File.Exists(f), "should not exist")
}

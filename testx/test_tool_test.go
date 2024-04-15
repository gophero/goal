package testx_test

import (
	"github.com/gophero/goal/testx"
	"testing"
)

func TestWrap(t *testing.T) {
	tr := testx.Wrap(t)
	tr.Case("wrapping testing.T")
	tr.Require(tr != nil, "wrapping should be success")
}

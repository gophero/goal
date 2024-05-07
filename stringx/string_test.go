package stringx_test

import (
	"fmt"
	"testing"

	"github.com/gophero/goal/assert"
	"github.com/gophero/goal/stringx"
	"github.com/gophero/goal/testx"
)

func TestBlurEmail(t *testing.T) {
	type Case struct {
		email  string
		expect string
	}
	cases := []Case{
		{"1313831783@qq.com", "13****3@qq.com"},
		{"belonk@126.com", "be****k@126.com"},
	}
	for _, c := range cases {
		dst := stringx.BlurEmail(c.email)
		if dst != c.expect {
			t.Errorf("test failed, expect: %v, but found: %v", c.expect, dst)
		}
	}
}

func TestEndsWith(t *testing.T) {
	assert.True(stringx.EndsWith("", ""))
	assert.True(stringx.EndsWith("a", ""))
	assert.True(!stringx.EndsWith("", "a"))

	s := "aaabb123b"
	assert.True(stringx.EndsWith(s, "b"))
	assert.True(stringx.EndsWith(s, "3b"))
	assert.True(stringx.EndsWith(s, "23b"))
	assert.True(stringx.EndsWith(s, "123b"))
	assert.True(!stringx.EndsWith(s, "a"))

	assert.True(stringx.StartsWith("", ""))
	assert.True(stringx.StartsWith("a", ""))
	assert.True(!stringx.StartsWith("", "a"))

	assert.True(stringx.StartsWith(s, "a"))
	assert.True(stringx.StartsWith(s, "aa"))
	assert.True(stringx.StartsWith(s, "aaa"))
	assert.True(stringx.StartsWith(s, "aaab"))
	assert.True(!stringx.StartsWith(s, "aaab1"))
	assert.True(!stringx.StartsWith(s, "1aaab1"))
}

func TestCamelCaseToUnderscore(t *testing.T) {
	cs := [][]string{
		{"HelloWorld", "hello_world"},
		{"helloWorld", "hello_world"},
		{"Helloworld", "helloworld"},
		{"AbcDEFGh", "abc_def_gh"},
		{"AbcDefGh", "abc_def_gh"},
		{"abcDefGh", "abc_def_gh"},
		{"abcDefGh😄", "abc_def_gh😄"},
	}

	tl := testx.Wrap(t)
	tl.Case("camelcase to underscore")

	for _, c := range cs {
		r := stringx.CamelCaseToUnderscore(c[0])
		tl.Require(r == c[1], "expect result is: %v, but is: %v", c[1], r)
	}
}

func TestUnderscoreToCamelCase(t *testing.T) {
	cs := [][]string{
		{"HelloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"Helloworld", "helloworld"},
		{"AbcDefGh", "abc_def_gh"},
		{"AbcDefGh中文", "abc_def_gh中文"},
	}

	tl := testx.Wrap(t)
	tl.Case("camelcase to underscore")

	for _, c := range cs {
		r := stringx.UnderscoreToCamelCase(c[1])
		tl.Require(r == c[0], "expect result is: %v, but is: %v", c[0], r)
	}
}

func TestFormatIntWithComma(t *testing.T) {
	d := 123456789
	s := stringx.FormatIntWithComma(int64(d))
	fmt.Println(s)
    fmt.Println(s == "123,456,789")
	assert.Equals("123,456,789", s)
}

func TestFormatFloatWithComma(t *testing.T) {
	f := 123456789.987654321
    fmt.Printf("%f\n",f)
	s := stringx.FormatFloatWithComma(f)
	fmt.Println(s)
	assert.Equals("123,456,789.987654", s)
	s = stringx.FormatFloatWithComma(f, 2)
	fmt.Println(s)
	assert.Equals("123,456,789.99", s)
}

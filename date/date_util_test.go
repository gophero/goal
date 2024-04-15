package date

import (
	"fmt"
	"testing"
)

func TestParseTime(t *testing.T) {
	r := ParseTime("10")
	fmt.Println(r)
	r = ParseTime("1:00")
	fmt.Println(r)
	r = ParseTime("01:01:00")
	fmt.Println(r)
}

package color_test

import (
	"github.com/gophero/goal/color"
	"github.com/gophero/goal/testx"
	"strings"
	"testing"
)

func TestRgb(t *testing.T) {
	cases := map[string][]int64{
		"000000": {0, 0, 0},
		"FFFFFF": {255, 255, 255},
		"FF0000": {255, 0, 0},
		"00FF00": {0, 255, 0},
		"0000FF": {0, 0, 255},
		"C8C8C8": {200, 200, 200},
	}
	logger := testx.Wrap(t)

	logger.Case("test rgb.Hex()")
	for k, v := range cases {
		rgb := color.NewRGB(v[0], v[1], v[2])
		hex := rgb.Hex()
		logger.Require(strings.ToUpper(hex) == k, "hex is correct. expects hex to be %s, found %s", hex, k)
	}

	logger.Case("test color.Hex2RGB()")
	for _, v := range cases {
		rgb := color.NewRGB(v[0], v[1], v[2])
		hex := rgb.Hex()
		rgb = color.Hex2RGB(hex)
		logger.Require(rgb.Red == v[0] && rgb.Green == v[1] && rgb.Blue == v[2], "rgb result should be correct, dest rgb is: %v, found: %v", rgb, v)
	}
}

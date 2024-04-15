package color

import "strconv"

type RGB struct {
	Red, Green, Blue int64
}

func NewRGB(r, g, b int64) RGB {
	return RGB{r, g, b}
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func (color RGB) Hex() string {
	r := t2x(color.Red)
	g := t2x(color.Green)
	b := t2x(color.Blue)
	return r + g + b
}

func Hex2RGB(hex string) RGB {
	r, _ := strconv.ParseInt(hex[:2], 16, 10)
	g, _ := strconv.ParseInt(hex[2:4], 16, 18)
	b, _ := strconv.ParseInt(hex[4:], 16, 10)
	return RGB{r, g, b}
}

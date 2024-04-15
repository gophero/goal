package printer

import (
	"fmt"
	"github.com/gophero/goal/conv"
	"strconv"
	"strings"
)

func NewLine() {
	fmt.Println()
}

func Ln() {
	fmt.Printf("\n")
}

func NewSepLine() {
	fmt.Println("========================================================================================")
}

func Printf(format string, s ...any) {
	fmt.Printf(format, s...)
}

func Println(format string, s ...any) {
	if format[len(format)-1:] != "\n" {
		format += "\n"
	}
	fmt.Printf(format, s...)
}

func Printw(width int, cols ...any) {
	printw(width, false, cols...)
}

// Printwln only support string、int、uint、float(保留两位小数)
func Printwln(width int, cols ...any) {
	printw(width, true, cols...)
}

func printw(width int, ln bool, cols ...any) {
	w := make([]any, len(cols))
	sb := strings.Builder{}
	for i, c := range cols {
		switch c.(type) {
		case string:
			sb.WriteString("%-" + strconv.Itoa(width) + "s")
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			sb.WriteString("%-" + strconv.Itoa(width) + "d")
		case float32, float64:
			// 获取小数位数
			s := strings.Split(fmt.Sprintf("%v", c), ".")
			var d = "2"
			if len(s) == 2 {
				d = conv.IntToStr(len(s[1]))
			}
			sb.WriteString("%-" + strconv.Itoa(width) + "." + d + "f")
		default:
			panic(fmt.Errorf("unsupported type: %v - %T", c, c))
		}
		w[i] = width
	}
	if ln {
		sb.WriteString("\n")
	}
	fmt.Printf(sb.String(), cols...)
}

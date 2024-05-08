package mathx

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"testing"

	"github.com/gophero/goal/assert"
)

func TestRound(t *testing.T) {
	r := math.Round(3.141592653) // round返回最接近的整数，从0开始4舍五入到整数，不论正负
	fmt.Printf("%.2f\n", r)      // 3.00
	r = math.Round(-3.141592653)
	fmt.Printf("%.2f\n", r) // -3.00
	r = math.Round(10.5)
	fmt.Printf("%.2f\n", r) // 11.00
	r = math.Round(-10.5)
	fmt.Printf("%.2f\n", r) // -11.00
	r = math.Round(-10.4)
	fmt.Printf("%.2f\n", r) // -10.00

	fmt.Printf("%.1f\n", 10.55)  // 10.6
	fmt.Printf("%.1f\n", 10.54)  // 10.5
	fmt.Printf("%.1f\n", -10.55) // -10.6
	fmt.Printf("%.1f\n", -10.54) // -10.5
}

func TestTrunc(t *testing.T) {
	// math.trunc 返回float的整数部分
	fmt.Printf("%.2f\n", math.Trunc(3.14))  // 3.00
	fmt.Printf("%.2f\n", math.Trunc(-3.14)) // -3.00
	fmt.Printf("%.2f\n", math.Trunc(3.55))  // 3.00
	fmt.Printf("%.2f\n", math.Trunc(-3.55)) // -3.00
}

func TestModf(t *testing.T) {
	i, frac := math.Modf(3.14)
	fmt.Printf("%.2f, %.2f\n", i, frac) // 3.00, 0.14
	i, frac = math.Modf(-2.71)
	fmt.Printf("%.2f, %.2f\n", i, frac) // -2.00, -0.71
}

func TestRound1(t *testing.T) {
	r := Round(3.144, 2)
	assert.True(fmt.Sprintf("%.2f", r) == "3.14")
	r = Round(3.145, 2)
	assert.True(fmt.Sprintf("%.2f", r) == "3.15")
	r = Round(3.145, -1)
	assert.True(fmt.Sprintf("%.3f", r) == "3.145")
	r = Round(3.145, 0)
	assert.True(fmt.Sprintf("%.2f", r) == "3.00")
}

func TestTrunc1(t *testing.T) {
	r := Trunc(3.145, -1)
	assert.True(fmt.Sprintf("%.3f", r) == "3.145")
	r = Trunc(3.145, 0)
	assert.True(fmt.Sprintf("%.2f", r) == "3.00")
	r = Trunc(3.145, 2)
	assert.True(fmt.Sprintf("%.2f", r) == "3.14")
	r = Trunc(3.144, 2)
	assert.True(fmt.Sprintf("%.2f", r) == "3.14")

	r = Trunc(3.144, 1)
	assert.True(fmt.Sprintf("%.2f", r) == "3.10")
	r = Trunc(3.144, 0)
	assert.True(fmt.Sprintf("%.2f", r) == "3.00")
}

func TestCeil(t *testing.T) {
	assert.True(fmt.Sprintf("%d", Ceil(3.0)) == "3")
	assert.True(fmt.Sprintf("%d", Floor(3.00)) == "3")

	assert.True(fmt.Sprintf("%d", Ceil(3.14)) == "4")
	assert.True(fmt.Sprintf("%d", Floor(3.14)) == "3")
}

func TestCeilr(t *testing.T) {
	assert.True(fmt.Sprintf("%.2f", Ceilr(3.14, 2)) == "3.14")
	assert.True(fmt.Sprintf("%.2f", Ceilr(3.14, 1)) == "3.20")
	assert.True(fmt.Sprintf("%.2f", Ceilr(3.14, 0)) == "4.00")
}

func TestCeilf(t *testing.T) {
	a := Ceilrf(3.14, 2)
	assert.True(a == "3.14")
	a = Ceilrf(3.14, 1)
	assert.True(a == "3.2")
	a = Ceilrf(3.14, 0)
	assert.True(a == "4")
}

func TestTrunc2(t *testing.T) {
	var f float64
	// f = -32.8584074
	// fmt.Println(Trunc(f, 2))
	// f = 113.03999999999999
	// fmt.Println(Trunc(f, 2))
	f = 437.09999999999997   // 精度很大，计算机无法准确表示
	fmt.Printf("%.5f\n", f)  // 437.10000
	fmt.Println(Trunc(f, 2)) // 437.1
	f = 437.09199999999997
	fmt.Println(Trunc(f, 2)) // 437.09
	f = 437.09599999999997
	fmt.Println(Trunc(f, 2)) // 437.09
}

func TestBig(t *testing.T) {
	var x, y, z big.Float
	x.SetInt64(100)                  // x is automatically set to 64bit precision
	y.SetFloat64(437.09999999999997) // y is automatically set to 53bit precision
	z.SetPrec(64)
	z.Mul(&x, &y)
	// prec：精度，acc：数否精确
	fmt.Printf("x = %.14f (%s, prec = %d, acc = %s)\n", &x, x.Text('f', 0), x.Prec(), x.Acc())
	fmt.Printf("y = %.14f (%s, prec = %d, acc = %s)\n", &y, y.Text('f', 0), y.Prec(), y.Acc())
	fmt.Printf("z = %.14f (%s, prec = %d, acc = %s)\n", &z, z.Text('f', 0), z.Prec(), z.Acc())
}

func FuzzTrunc(f *testing.F) {
	f.Add(3.1415926)
	f.Add(3.14)
	f.Add(math.Pi)
	f.Add(math.E)
	f.Fuzz(func(t *testing.T, x float64) {
		n := rand.Intn(1 << 3)
		y := Trunc(x, n)
		if x != y {
			xs := fmt.Sprintf("%v", y)
			ss := strings.Split(xs, ".")
			if len(ss) > 1 {
				assert.True(len(ss[1]) <= n) // 小数位数必然 <= 2
			}
			//if x > 0 {
			//	assert.True(x > y) // x 精度损失，但是截断后必然 > y
			//}
		}
	})
}

func TestFormatCommaFloat(t *testing.T) {
	f := 123456789.987654321
	fmt.Printf("%f\n", f)
	// default mod
	s := FmtCommaFloat(f)
	fmt.Println(s)
	assert.Equals("123,456,789.987654", s)
	// half down
	s = FmtCommaFloat(f, NewPrec(2, HalfDown))
	fmt.Println(s)
	assert.Equals("123,456,789.98", s)
	// half up
	s = FmtCommaFloat(f, NewPrec(2, HalfUp))
	fmt.Println(s)
	assert.Equals("123,456,789.99", s)
}

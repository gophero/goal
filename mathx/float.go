package mathx

import (
	"fmt"
	"github.com/gophero/goal/conv"
	"github.com/gophero/goal/valuex"
	"math"
	"strconv"
)

type Float interface {
	~float32 | ~float64
}

// Round 将 f 执行四舍五入保留 n 位小数。
func Round[T Float](f T, n int) T {
	if n < 0 {
		return f
	}
	if n == 0 {
		return T(math.Round(float64(f)))
	}
	return T(valuex.Must(strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(n)+"f", f), 64)))
}

// Trunc 将 f 保留 n 为小数，第 n 位后的小数全部舍去。
func Trunc[T Float](f T, n int) T {
	if n < 0 {
		return f
	}
	if n == 0 {
		return T(math.Trunc(float64(f)))
	}
	base := math.Pow10(n)
	//x, y := math.Modf(float64(f)) // 损失精度
	//i := int64(y * base)
	//y = float64(i) / base
	//return T(x + y)
	i := int64(Mul(float64(f), base))
	x := float64(i) / base
	return T(x)
}

// Truncf 将 f 先执行 Trunc(f, n) 再执行 Format 格式化为保留 n 为小数的字符串。
func Truncf[T Float](f T, n int) string {
	return Format(Trunc(f, n), n)
}

// Format 将 f 四舍五入保留 n 位小数并格式化为字符串。
func Format[T Float](f T, n int) string {
	return fmt.Sprintf("%."+conv.IntToStr(n)+"f", f)
}

// Ceil 将 f 向上取整
func Ceil[T Float](f T) int64 {
	x, y := math.Modf(float64(f))
	if y == 0 {
		return int64(x)
	}
	return int64(x + 1)
}

// Floor 将 f 向下取整，效果等同于 Trunc(f, 0)。
func Floor[T Float](f T) int64 {
	x, _ := math.Modf(float64(f))
	return int64(x)
}

// Ceilr 将 f 保留 n 为小数，如果 n 位小数后边还有值则第 n 位加1。
func Ceilr[T Float](f T, n int) T {
	if n < 0 {
		return f
	}
	if n == 0 {
		return T(Ceil(f))
	}
	x, y := math.Modf(float64(f))
	if y == 0 {
		return f
	}
	cardinal := math.Pow10(n)
	d := f * T(cardinal)
	a, b := math.Modf(float64(d)) // TODO 处理精度损失
	if b == 0 {
		return T(x + Trunc(y, n))
	}
	return T((a + 1) / cardinal)
}

// Ceilrf 将 f 先执行 Cerilr(f, n)，然后再执行 Format 格式化为保留 n 位小数的字符串。
func Ceilrf[T Float](f T, n int) string {
	return Format(Ceilr(f, n), n)
}

package mathx

import "math/big"

func Add(addend1 float64, addend2 float64) float64 {
	result, _ := NewFromFloat(addend1).Add(NewFromFloat(addend2)).Float64()
	return result
}

func Sub(minuend float64, subtrahend float64) float64 {
	result, _ := NewFromFloat(minuend).Sub(NewFromFloat(subtrahend)).Float64()
	return result
}

func Mul(multiplicand float64, multiplier float64) float64 {
	result, _ := NewFromFloat(multiplicand).Mul(NewFromFloat(multiplier)).Float64()
	return result
}

func Div(dividend float64, divisor float64, precision int32) float64 {
	result, _ := NewFromFloat(dividend).DivRound(NewFromFloat(divisor), precision).Float64()
	return result
}

func DivTrunc(dividend float64, divisor float64, precision int32) float64 {
	result, _ := NewFromFloat(dividend).Div(NewFromFloat(divisor)).Truncate(precision).Float64()
	return result
}

func DivInt64(dividend int64, divisor int64, precision int32) float64 {
	result, _ := NewFromInt(dividend).DivRound(NewFromInt(divisor), precision).Float64()
	return result
}

func DivTruncInt64(dividend int64, divisor int64, divisorTwo int64, precision int32) float64 {
	result, _ := NewFromInt(dividend).Div(NewFromInt(divisor)).Div(NewFromInt(divisorTwo)).Truncate(precision).Float64()
	return result
}

func MulInt64AndFloat64(multiplicand int64, multiplier float64) float64 {
	result, _ := NewFromInt(multiplicand).Mul(NewFromFloat(multiplier)).Float64()
	return result
}

func MulBigFloat(multiplicand *big.Float, multiplier *big.Float) *big.Float {
	m1, _ := multiplicand.Float64()
	m2, _ := multiplier.Float64()
	result := Mul(m1, m2)
	return big.NewFloat(result)
}

// func CompareToBigFloat(f1 *big.Float, f2 *big.Float) bool {
//	return f1.
// }

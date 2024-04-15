package conv_test

import (
	"fmt"
	"github.com/gophero/goal/conv"
	"math"
	"math/big"
	"reflect"
	"testing"
)

func ExampleInt64ToHex() {
	var x int64 = math.MinInt64
	var y int64 = math.MaxInt64
	var z int64 = math.MaxInt32

	fmt.Println(x, y, z) // -9223372036854775808 9223372036854775807 2147483647

	xs := conv.Int64ToHex(x)
	ys := conv.Int64ToHex(y)
	zs := conv.Int64ToHex(z)
	fmt.Println(xs, ys, zs) // -8000000000000000 7fffffffffffffff 7fffffff
}

func ExampleHexToInt64() {
	xs := "-8000000000000000"
	ys := "7fffffffffffffff"
	zs := "7fffffff"

	x := conv.HexToInt64(xs)
	y := conv.HexToInt64(ys)
	z := conv.HexToInt64(zs)
	fmt.Println(x, y, z) // -9223372036854775808 9223372036854775807 2147483647
}

func TestHexToInt64(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "string param: 0", args: args{src: "0"}, want: 0},
		{name: "string param: -1", args: args{src: "-1"}, want: -1},
		{name: "string param: 10", args: args{src: "A"}, want: 10},
		{name: "string param: 11", args: args{src: "B"}, want: 11},
		{name: "string param: 17", args: args{src: "11"}, want: 17},
		{name: "string param: -8000000000000000", args: args{src: "-8000000000000000"}, want: math.MinInt64},
		{name: "string param: 7fffffffffffffff", args: args{src: "7fffffffffffffff"}, want: math.MaxInt64},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.HexToInt64(tt.args.src); got != tt.want {
				t.Errorf("HexToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32ToStr(t *testing.T) {
	type args struct {
		src int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "2147483647"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-2147483648"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.Int32ToStr(tt.args.src); got != tt.want {
				t.Errorf("Int32ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -9223372036854775808 9223372036854775807
func TestInt64ToHex(t *testing.T) {
	type args struct {
		src int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "7fffffff"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-80000000"},
		{name: "param: 9223372036854775807", args: args{src: math.MaxInt64}, want: "7fffffffffffffff"},
		{name: "param: -9223372036854775808", args: args{src: math.MinInt64}, want: "-8000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.Int64ToHex(tt.args.src); got != tt.want {
				t.Errorf("Int64ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ToStr(t *testing.T) {
	type args struct {
		src int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "2147483647"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-2147483648"},
		{name: "param: 9223372036854775807", args: args{src: math.MaxInt64}, want: "9223372036854775807"},
		{name: "param: -9223372036854775808", args: args{src: math.MinInt64}, want: "-9223372036854775808"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.Int64ToStr(tt.args.src); got != tt.want {
				t.Errorf("Int64ToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToStr(t *testing.T) {
	type args struct {
		src int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "param: 0", args: args{src: 0}, want: "0"},
		{name: "param: -1", args: args{src: -1}, want: "-1"},
		{name: "param: 2147483647", args: args{src: math.MaxInt32}, want: "2147483647"},
		{name: "param: -2147483648", args: args{src: math.MinInt32}, want: "-2147483648"},
		{name: "param: 9223372036854775807", args: args{src: math.MaxInt64}, want: "9223372036854775807"},
		{name: "param: -9223372036854775808", args: args{src: math.MinInt64}, want: "-9223372036854775808"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.IntToStr(tt.args.src); got != tt.want {
				t.Errorf("IntToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntsToStr(t *testing.T) {
	type args struct {
		is []int64
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "case 1", args: args{is: []int64{-1, 0, 1}}, want: []string{"-1", "0", "1"}},
		{name: "case 2", args: args{is: []int64{math.MinInt64, 0, math.MaxInt64}}, want: []string{"-9223372036854775808", "0", "9223372036854775807"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.IntsToStr(tt.args.is); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntsToStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinBigInt(t *testing.T) {
	var bigs = []*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(math.MaxInt64)}
	type args struct {
		ints []*big.Int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case 1", args: args{ints: []*big.Int{big.NewInt(0)}}, want: "0"},
		{name: "case 2", args: args{ints: []*big.Int{big.NewInt(-1)}}, want: "-1"},
		{name: "case 3", args: args{ints: bigs}, want: "1,0,9223372036854775807"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.JoinBigInt(tt.args.ints); got != tt.want {
				t.Errorf("JoinBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitStrToInt(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "case 1", args: args{s: "-1,0,1", sep: ","}, want: []int64{-1, 0, 1}},
		{name: "case 2", args: args{s: "-1", sep: ","}, want: []int64{-1}},
		{name: "case 3", args: args{s: "0", sep: ""}, want: []int64{0}},
		{name: "case 4", args: args{s: "0#1#2#3", sep: "#"}, want: []int64{0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.SplitStrToInt(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToFloat64(t *testing.T) {
	type args struct {
		amount string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "case 1", args: args{amount: "0"}, want: 0},
		{name: "case 2", args: args{amount: "1"}, want: 1},
		{name: "case 3", args: args{amount: "-1"}, want: -1},
		{name: "case 4", args: args{amount: "3.1415"}, want: 3.1415},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := conv.StrToFloat64(tt.args.amount); got != tt.want {
				t.Errorf("StrToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

package conv

import (
	"fmt"
	"github.com/gophero/goal/errorx"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

// StrToInt covert string to int
func StrToInt(str string) int {
	n, err := strconv.Atoi(str)
	errorx.Throw(err)
	return n
}

// StrToInt32 covert string to int32
func StrToInt32(str string) int32 {
	n, err := strconv.ParseInt(str, 10, 32)
	errorx.Throw(err)
	return int32(n)
}

// StrToInt64 covert string to int64
func StrToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	errorx.Throw(err)
	return n
}

// IntToStr covert int to string
func IntToStr(src int) string {
	return strconv.Itoa(src)
}

// Int64ToStr covert int64 to string
func Int64ToStr(src int64) string {
	return strconv.FormatInt(src, 10)
}

// Int32ToStr covert int32 to string
func Int32ToStr(src int32) string {
	return strconv.FormatInt(int64(src), 10)
}

// StrToUint covert string to uint
func StrToUint(str string) uint {
	n, err := strconv.ParseUint(str, 10, 64)
	errorx.Throw(err)
	return uint(n)
}

// StrToUint64 covert string to uint64
func StrToUint64(str string) uint64 {
	n, err := strconv.ParseUint(str, 10, 64)
	errorx.Throw(err)
	return n
}

// StrToUint32 covert string to uint32
func StrToUint32(str string) uint32 {
	n, err := strconv.ParseInt(str, 10, 32)
	errorx.Throw(err)
	return uint32(n)
}

// UintToStr covert uint to string
func UintToStr(src uint) string {
	return strconv.FormatUint(uint64(src), 10)
}

// Uint64ToStr covert uint64 to string
func Uint64ToStr(src uint64) string {
	return strconv.FormatUint(src, 10)
}

// Uint32ToStr covert uint32 to string
func Uint32ToStr(src uint32) string {
	return strconv.FormatUint(uint64(src), 10)
}

// JoinBigInt join a slice of big.Int to a string with ","
func JoinBigInt(ints []*big.Int) string {
	var temp = make([]string, len(ints))
	for k, v := range ints {
		temp[k] = fmt.Sprintf("%d", v.Int64())
	}
	var result = strings.Join(temp, ",")
	return result
}

// StrToFloat64 convert string to float64
func StrToFloat64(amount string) float64 {
	float, err := strconv.ParseFloat(amount, 64)
	errorx.Throw(err)
	return float
}

// Int64ToHex convert int64 to hex string
func Int64ToHex(src int64) string {
	return strconv.FormatInt(src, 16)
}

// HexToInt64 convert hex string to int64
func HexToInt64(src string) int64 {
	id, err := strconv.ParseInt(src, 16, 64)
	errorx.Throw(err)
	return id
}

// SplitStrToInt split a string to int slice with given separator.
func SplitStrToInt(s string, sep string) []int64 {
	ss := strings.Split(s, sep)
	var is []int64
	for _, i := range ss {
		it, err := strconv.ParseInt(i, 10, 64)
		errorx.Throw(err)
		is = append(is, it)
	}
	return is
}

// IntsToStr convert a int slice to a string slice.
func IntsToStr(is []int64) []string {
	if is == nil || len(is) == 0 {
		return nil
	}
	var ss = make([]string, len(is))
	for i, v := range is {
		ss[i] = strconv.FormatInt(v, 10)
	}
	return ss
}

// StrsToInt convert a str slice to a int slice.
func StrsToInt(ss []string) []int64 {
	if ss == nil || len(ss) == 0 {
		return nil
	}
	var is = make([]int64, len(ss))
	for _, s := range ss {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil
		}
		is = append(is, i)
	}
	return is
}

func StringToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return b
}

const (
	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func IntToBase62(n int) string {
	if n == 0 {
		return string(base62[0])
	}

	var result []byte
	for n > 0 {
		result = append(result, base62[n%62])
		n /= 62
	}

	// 反转字符串
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func InterToArray(obj interface{}) []uint {
	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		s := reflect.ValueOf(obj)
		arrays := make([]uint, 0)
		for i := 0; i < s.Len(); i++ {
			ele := s.Index(i)
			e := ele.Interface().(*big.Int)
			arrays = append(arrays, uint(e.Uint64()))
		}
		return arrays
	}
	return make([]uint, 0)
}

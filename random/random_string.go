package random

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const alphanumeric = letters + numbers
const hexStr = "0123456789abcdef"

func Alphabetic(len int) string {
	return String([]byte(letters), len)
}

func Numeric(len int) string {
	return String([]byte(numbers), len)
}

func Alphanumeric(length int) string {
	return String([]byte(alphanumeric), length)
}

func Hex(length int) string {
	return String([]byte(hexStr), length)
}

func String(chars []byte, length int) string {
	selCharLen := len(chars)
	if selCharLen < 2 {
		panic("illegal argument: length of given chars need > 1")
	}
	if length <= 0 {
		panic("illegal argument: length need > 0")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	builder := strings.Builder{}

	for i := 0; i < length; i++ {
		i := r.Intn(selCharLen)
		builder.WriteByte(chars[i])
	}
	return builder.String()
}

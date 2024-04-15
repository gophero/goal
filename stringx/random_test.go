package stringx

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	////for i := 0; i < 10; i++ {
	s := Randn(16)
	fmt.Println(s)
	////}
}

func TestRandId(t *testing.T) {
	s := RandId()
	fmt.Println(s)
}

func TestCryptoRead(t *testing.T) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b) // crypto/rand 包实现了一个加密安全的随机数生成器，基于操作系统。
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// The slice should now contain random bytes instead of only zeroes.
	fmt.Println(bytes.Equal(b, make([]byte, c))) // false，已经读取到字节
	fmt.Println(b)                               // 非 0
}

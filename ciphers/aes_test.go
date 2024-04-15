package ciphers_test

import (
	"fmt"
	"github.com/gophero/goal/ciphers"
	"github.com/gophero/goal/random"
	"testing"
)

var plainText = "exampleplaintext" // 16 bytes
var key = []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}

func TestAES(t *testing.T) {
	// ECB
	s, _ := ciphers.AES.Encrypt([]byte(plainText), key, ciphers.ECB, nil)
	fmt.Printf("%x\n", s) // 固定密文：40d32c0de54cc6b82c39f22e641455d3a254be88e037ddd9d79fb6411c3f9df8
	s, _ = ciphers.AES.Decrypt(s, key, ciphers.ECB, nil)
	fmt.Println(string(s))

	// CBC
	// 随机 iv，长度为16字节，每次密文不同
	r := random.Hex(len(key))
	fmt.Println("random iv:", r)
	iv := []byte(r)
	s, _ = ciphers.AES.Encrypt([]byte(plainText), key, ciphers.CBC, iv)
	fmt.Printf("%x\n", s) // 随机密文
	s, _ = ciphers.AES.Decrypt(s, key, ciphers.CBC, iv)
	fmt.Println(string(s))

	// 固定 iv，每次加密密文相同
	iv = key
	s, _ = ciphers.AES.Encrypt([]byte(plainText), key, ciphers.CBC, iv)
	fmt.Printf("%x\n", s) // 固定密文：0735437968e811771051aa81734b1098b8353285c0c9517a752a429a3efc44fe
	s, _ = ciphers.AES.Decrypt(s, key, ciphers.CBC, iv)
	fmt.Println(string(s))
}

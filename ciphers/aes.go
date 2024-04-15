package ciphers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/gophero/goal/errorx"
)

// AES 加密算法
//
// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
var AES = &aeser{}

type aeser struct {
}

type AesMode int

const (
	// ECB 加密模式：https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#ECB
	// ECB 模式存在安全性问题，不建议使用，详见：https://crypto.stackexchange.com/questions/20941/why-shouldnt-i-use-ecb-encryption/20946#20946
	ECB AesMode = iota
	// CBC 加密模式：https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#CBC
	CBC
)

func (a *aeser) Encrypt(rawBytes []byte, key []byte, mode AesMode, iv []byte) ([]byte, error) {
	switch mode {
	case ECB:
		return aesEncryptECB(rawBytes, key)
	case CBC:
		return aesEncryptCBC(rawBytes, key, iv)
	default:
		return nil, errorx.New("unsupported encrypt mode: %v", mode)
	}
}

func (a *aeser) Decrypt(cipherBytes []byte, key []byte, mode AesMode, iv []byte) ([]byte, error) {
	switch mode {
	case ECB:
		return aesDecryptECB(cipherBytes, key)
	case CBC:
		return aesDecryptCBC(cipherBytes, key, iv)
	default:
		return nil, errorx.New("unsupported decrypt mode: %v", mode)
	}
}

// pkcs7Padding pkcs7 填充，详见：https://en.wikipedia.org/wiki/Padding_(cryptography)
func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	// 判断缺少几位长度，最少1，最多 blockSize
	padding := blockSize - len(cipherText)%blockSize
	// 将填充的数量 padding 作为每字节填充内容（反向填充时可以直接获取），并复制 padding 个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// 填充完成，拼接字节
	return append(cipherText, padText...)
}

// pkcs7UnPadding pkcs7 取消填充
func pkcs7UnPadding(cipherText []byte) []byte {
	length := len(cipherText)
	// 获取填充的个数，最后一字节的整数代表了已经填充的字节数
	unPadding := int(cipherText[length-1])
	// 去掉填充的字节
	return cipherText[:(length - unPadding)]
}

func aesEncryptECB(rawText []byte, key []byte) ([]byte, error) {
	// 创建 cipher，如果 key 长度不是16、24、32字节，则会 panic
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// AES分组后每组长度为16字节，128位，所以 blockSize = 16 字节
	bs := block.BlockSize()
	// 使用 pkcs#7 加密模式进行填充
	rawText = pkcs7Padding(rawText, bs)
	// 填充完成后，被加密字节数组的长度必须块大小的倍数，即16的倍数
	if len(rawText)%bs != 0 {
		panic("block size padding failed")
	}

	out := make([]byte, len(rawText))
	dst := out
	// 对原文依次分组加密
	for len(rawText) > 0 {
		block.Encrypt(dst, rawText[:bs])
		rawText = rawText[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func aesDecryptECB(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(cipherText)%bs != 0 {
		panic("illegal ciphertext length")
	}

	out := make([]byte, len(cipherText))
	dst := out
	// 解密
	for len(cipherText) > 0 {
		block.Decrypt(dst, cipherText[:bs])
		cipherText = cipherText[bs:]
		dst = dst[bs:]
	}
	out = pkcs7UnPadding(out)
	return out, nil
}

func aesEncryptCBC(rawBytes []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	rawBytes = pkcs7Padding(rawBytes, blockSize)
	// 创建 CBC 加密器，初始向量 iv 长度必须等于块长度 blockSize
	blockMode := cipher.NewCBCEncrypter(block, iv) // 初始向量的长度必须等于块block的长度16字节
	dst := make([]byte, len(rawBytes))
	blockMode.CryptBlocks(dst, rawBytes)
	return dst, nil
}

func aesDecryptCBC(cipherBytes []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv) // 初始向量的长度必须等于块block的长度16字节
	rawBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(rawBytes, cipherBytes)
	rawBytes = pkcs7UnPadding(rawBytes)
	return rawBytes, nil
}

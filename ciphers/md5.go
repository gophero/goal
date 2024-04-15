package ciphers

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

package base64x

import "encoding/base64"

// StdEncoding use base64.StdEncoding to encode and decode
var StdEncoding base64std

// RawStdEncoding use base64.RawStdEncoding to encode and decode
var RawStdEncoding base64raw

type base64std struct {
}

func (base64std) Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (base64std) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.StdEncoding.Strict().DecodeString(str)
	} else {
		return base64.StdEncoding.DecodeString(str)
	}
}

type base64raw struct {
}

func (base64raw) Encode(b []byte) string {
	return base64.RawStdEncoding.EncodeToString(b)
}

func (base64raw) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.RawStdEncoding.Strict().DecodeString(str)
	} else {
		return base64.RawStdEncoding.DecodeString(str)
	}
}

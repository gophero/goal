package base64x

import "encoding/base64"

// URLEncoding use base64.URLEncoding to encode and decode
var URLEncoding base64Url

// RawURLEncoding use base64.RawURLEncoding to encode and decode
var RawURLEncoding base64RawUrl

type base64Url struct {
}

func (base64Url) Encode(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

func (base64Url) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.URLEncoding.Strict().DecodeString(str)
	} else {
		return base64.URLEncoding.DecodeString(str)
	}
}

type base64RawUrl struct {
}

// Encode encodes the given bytes using strict base64url encoding.
func (base64RawUrl) Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

// Decode decodes the given string using strict base64url encoding.
func (base64RawUrl) Decode(str string, strict ...bool) ([]byte, error) {
	if len(strict) > 0 && strict[0] {
		return base64.RawURLEncoding.Strict().DecodeString(str)
	} else {
		return base64.RawURLEncoding.DecodeString(str)
	}
}

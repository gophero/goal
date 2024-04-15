package base64x_test

import (
	"fmt"
	"github.com/gophero/goal/ciphers/base64x"
	"github.com/gophero/goal/testx"
	"reflect"
	"testing"
)

const (
	plaintext = "abcdefghijjklmnopqrstuvwxyz0123456789`~-_=+[]\\{}|;':\",./<>?"
	base64enc = "YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB-LV89K1tdXHt9fDsnOiIsLi88Pj8"
)

func Test_Base64UrlEncode(t *testing.T) {
	lg := testx.Wrap(t)
	lg.Case("testing base64s.URLEncoding.Encode")
	enc := base64x.RawURLEncoding.Encode([]byte(plaintext))
	fmt.Println(enc)
	lg.Require(base64enc == enc, "result should match")
}

func Test_Base64UrlDecode(t *testing.T) {
	lg := testx.Wrap(t)
	lg.Case("testing base64s.URLEncoding.Decode")
	dec, err := base64x.RawURLEncoding.Decode(base64enc, true)
	lg.Require(err == nil, "requires no error")
	lg.Require(reflect.DeepEqual([]byte(plaintext), dec), "results should be matched")
}

package base64x_test

import (
	"github.com/gophero/goal/ciphers/base64x"
	"github.com/gophero/goal/testx"
	"math/big"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tt := []struct {
		input func() *big.Int
		want  string
	}{
		{
			input: func() *big.Int {
				bi := big.Int{}
				return bi.SetInt64(65537)
			},
			want: "AQAB",
		},
		{
			input: func() *big.Int {
				return &big.Int{}
			},
			want: "AA",
		},
	}

	lg := testx.Wrap(t)
	lg.Case("Test base64s.Base64UrlUint.Encode")
	for _, test := range tt {
		got := base64x.Base64UrlUint.Encode(test.input())

		lg.Require(test.want == got, "result shoud match")
	}
}

func Test_Base64UrlUintDecode(t *testing.T) {
	tt := []struct {
		input string
		valid bool
		want  func() *big.Int
	}{
		{
			input: "AQAB",
			valid: true,
			want: func() *big.Int {
				bi := &big.Int{}
				return bi.SetBytes([]byte{1, 0, 1})
			},
		},
		{
			input: "AA",
			valid: true,
			want: func() *big.Int {
				bi := &big.Int{}
				return bi.SetBytes([]byte{0})
			},
		},
		{
			input: "",
			valid: true,
			want: func() *big.Int {
				return nil
			},
		},
		{
			input: "Invalid Base 64 URL ===",
			valid: false,
		},
	}

	lg := testx.Wrap(t)
	lg.Case("Test base64s.Base64UrlUint.Decode")

	for _, test := range tt {
		got, err := base64x.Base64UrlUint.Decode(test.input)
		if test.valid {
			lg.Require(err == nil, "requires no error")
		} else {
			lg.Require(err != nil, "requires an error")
			continue
		}

		lg.Require(reflect.DeepEqual(got, test.want()), "Input: %s, Got: %+v, Want: %+v", test.input, got, test.want())
	}
}

var s2 = "h0HtbcA_ud27f5vc4U_9OsB2fn3Ar5QD6bpuHB1VGTXDB_zIko2ENmtHQmJAZJEEGJxA5v1fzs7v3Yk6WRY7XbJFYvKWr8A7_txUgwPCFaR0eH1HpiCbldw4X6Y690O75ksoSepbyYwmdi5u2JqX1lz3a2O5taYdBYC0pO6gaNfgT-lYSf4Ws5CAZND3qhMLD8Cnby4n0Hxj6xnpr8ODAnVNbWQ0JECthfjolCI026t87kC7S5hHSnd2DFvM4arHE7TRj__3SrBKzcJZxM70ApNkAwytOUgLbHKmL9x2IXW5x650mqloaR0ZHiizD9vjvzFm42D9OqDYcAaywsZotQ"
var int2 = "17074681210596346338238930741604719911471739263388201349161845562864656504969182747504982503280340095360421241540209107445393214665406042841653086840197654417437871915910700267737069351329285696844640299283651360376405207418980525942229586349408618823167696115135581627673750603170073872274583421669922746451810309004616254464214848244623392288415427808688666932715607147074268328966971346040049513796636811220067839144851245646757079953966882044326791197178356185865975824396146836926034001801190192888799824010789817258551225187590168318741899180377721337496032847013690934283071678384120201592932662085385131026613"

func TestEncLongString(t *testing.T) {
	bi, err := base64x.Base64UrlUint.Decode(s2)
	if err != nil {
		t.Errorf("want no error: %v", err)
		t.Failed()
	}
	if reflect.DeepEqual(bi, int2) {
		t.Errorf("result is incorrect")
		t.Failed()
	}

	it, ok := new(big.Int).SetString(int2, 10)
	if !ok {
		t.Errorf("create big.int failed: %v", ok)
	} else {
		enc := base64x.Base64UrlUint.Encode(it)
		if enc != s2 {
			t.Errorf("result is incorrect")
			t.Failed()
		}
	}
}

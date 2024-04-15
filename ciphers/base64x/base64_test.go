package base64x_test

import (
	"fmt"
	"github.com/gophero/goal/ciphers/base64x"
	"reflect"
	"testing"
)

var s = "abcdefghijjklmnopqrstuvwxyz0123456789`~-_=+[]\\{}|;':\",./<>?"
var sw = "YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB+LV89K1tdXHt9fDsnOiIsLi88Pj8="
var rw = "YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB+LV89K1tdXHt9fDsnOiIsLi88Pj8"
var s1 = "hello, 这是中文!"
var sw1 = "aGVsbG8sIOi/meaYr+S4reaWhyE="
var rw1 = "aGVsbG8sIOi/meaYr+S4reaWhyE"

func TestBase(t *testing.T) {
	// YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB+LV89K1tdXHt9fDsnOiIsLi88Pj8=
	fmt.Println(base64x.StdEncoding.Encode([]byte(s)))
	// YWJjZGVmZ2hpamprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OWB+LV89K1tdXHt9fDsnOiIsLi88Pj8
	fmt.Println(base64x.RawStdEncoding.Encode([]byte(s)))

	// aGVsbG8sIOi/meaYr+S4reaWhyE=
	fmt.Println(base64x.StdEncoding.Encode([]byte(s1)))
	// aGVsbG8sIOi/meaYr+S4reaWhyE
	fmt.Println(base64x.RawStdEncoding.Encode([]byte(s1)))
}

func Test_base64raw_Decode(t *testing.T) {
	type args struct {
		str    string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "case1", args: args{str: rw, strict: []bool{true}}, want: []byte(s), wantErr: false},
		{name: "case2", args: args{str: rw1, strict: []bool{true}}, want: []byte(s1), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.RawStdEncoding
			got, err := ba.Decode(tt.args.str, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64raw_Encode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{b: []byte(s)}, want: rw},
		{name: "case2", args: args{b: []byte(s1)}, want: rw1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.RawStdEncoding
			if got := ba.Encode(tt.args.b); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64std_Decode(t *testing.T) {
	type args struct {
		str    string
		strict []bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "case1", args: args{str: sw, strict: []bool{true}}, want: []byte(s), wantErr: false},
		{name: "case2", args: args{str: sw1, strict: []bool{true}}, want: []byte(s1), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.StdEncoding
			got, err := ba.Decode(tt.args.str, tt.args.strict...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64std_Encode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1", args: args{b: []byte(s)}, want: sw},
		{name: "case2", args: args{b: []byte(s1)}, want: sw1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := base64x.StdEncoding
			if got := ba.Encode(tt.args.b); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

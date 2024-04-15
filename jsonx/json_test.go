package jsonx_test

import (
	"fmt"
	"github.com/gophero/goal/jsonx"
	"github.com/gophero/goal/testx"
	"testing"
	"time"
)

func TestToJson(t *testing.T) {
	var mp = make(map[string]any)
	var smp = make(map[string]any)
	mp["a"] = 1
	mp["b"] = 2

	smp["s1"] = 11
	smp["s2"] = 22
	mp["c"] = smp

	println(jsonx.ToJson(mp))
}

type User struct {
	Name      string    // `json:"name"`
	Age       int       // `json:"age"`
	BirthDate time.Time // `json:"birthDate"`
	Other     float32   // 其他字段，科学计数法的情况
}

func TestToJson1(t *testing.T) {
	var user = User{
		Name:      "张三",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		Other:     171.65,
	}
	println(jsonx.ToJson(user))
}

func TestToJsonf(t *testing.T) {
	var user = User{
		Name:      "张三",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
		Other:     171.65,
	}
	println(jsonx.ToJson(user, jsonx.Indent))
}

func TestParse(t *testing.T) {
	s := "{\"Name\":\"张三\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00+08:00\",\"Other\":171.65}"
	u, _ := jsonx.Parse([]byte(s), &User{})
	fmt.Printf("user: %v\n", u)
}

func TestParse1(t *testing.T) {
	logger := testx.Wrap(t)

	logger.Case("parse user json")
	s := "{\n  \"Name\": \"张三\",\n  \"Age\": 20,\n  \"BirthDate\": \"2000-01-01T00:00:00+08:00\",\n  \"Other\":17165123123}"
	u, _ := jsonx.Parse([]byte(s), &User{})
	fmt.Printf("user: %v\n", u)
}

func TestParse2(t *testing.T) {
	s := "{\"Name\":\"张三\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00+08:00\",\"Other\":17165123123}"
	var mp = make(map[string]any)
	jsonx.Parse([]byte(s), &mp)
	// jsoner.Parse(s, &mp, jsoner.UseNumber) // 使用参数将 Other 处理为 number，而不是 float64
	fmt.Printf("user: %v\n", mp)
	fmt.Printf("name: %v\n", mp["Name"])
	fmt.Printf("age: %v\n", mp["Age"])
	fmt.Printf("date: %v\n", mp["BirthDate"])
	fmt.Printf("other: %v\n", mp["Other"]) // 不处理number，则为：1.7165123123e+10，处理后为 17165123123
}

func TestParseDisallowUnknownFields(t *testing.T) {
	logger := testx.Wrap(t)

	logger.Case("add a gender field which will cause panic")
	// 增加 gender 属性
	s := "{\"Name\":\"张三\",\"gender\":1,\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00+08:00\",\"Other\":171.65}"
	defer func() {
		if err := recover(); err != nil {
			logger.Pass("should panic")
		} else {
			logger.Fail("should panic")
		}
	}()
	u, _ := jsonx.Parse([]byte(s), &User{}, jsonx.DisallowUnknownFields)
	fmt.Printf("user: %v\n", u)
}

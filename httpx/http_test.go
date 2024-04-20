package httpx_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gophero/goal/httpx"
	"github.com/stretchr/testify/assert"
)

func TestRespEmpty(t *testing.T) {
	resp := &http.Response{StatusCode: http.StatusOK, Status: "200 ok", Body: http.NoBody}
	r := httpx.Resp(resp)
	s := r.Str()
	assert.True(t, s == "")
	assert.Nil(t, r.Err())
}

func TestRespStringBody(t *testing.T) {
	bodystr := "hello, body"
	sr := strings.NewReader(bodystr)
	resp := &http.Response{StatusCode: http.StatusOK, Status: "200 ok", Body: io.NopCloser(sr)}
	r := httpx.Resp(resp)
	s := r.Str()
	assert.True(t, s == bodystr)
	assert.Nil(t, r.Err())
}

func TestRespObjBody(t *testing.T) {
	bd := TestBody{F: "field", A: 1024}
	bs, _ := json.Marshal(&bd)
	br := bytes.NewReader(bs)
	resp := &http.Response{StatusCode: http.StatusOK, Status: "200 ok", Body: io.NopCloser(br)}
	r := httpx.Resp(resp)
	s := r.Str()
	fmt.Println(s)
	assert.True(t, s == string(bs))
	assert.Nil(t, r.Err())
	var bdr TestBody
	obj := r.JsonObj(&bdr).(*TestBody)
	assert.Equal(t, obj, &bdr)
}

func TestRespSliceBody(t *testing.T) {
	bds := []TestBody{
		{F: "field1", A: 1023},
		{F: "field2", A: 1024},
	}
	bs, _ := json.Marshal(&bds)
	br := bytes.NewReader(bs)
	resp := &http.Response{StatusCode: http.StatusOK, Status: "200 ok", Body: io.NopCloser(br)}
	r := httpx.Resp(resp)

	var bdrs []TestBody
	ss := r.JsonObj(&bdrs).(*[]TestBody)
	assert.Equal(t, &bdrs, ss)
	assert.Nil(t, r.Err())
}

type TestBody struct {
	F string `json:"f"`
	A int32  `json:"a"`
}

func TestNotSuccess(t *testing.T) {
	br := bytes.NewReader([]byte("test not found"))
	resp := &http.Response{Status: "404 Not Found", Body: io.NopCloser(br)}
	r := httpx.Resp(resp)
	s := r.Str()
	assert.True(t, s == "")
	assert.Equal(t, r.Err().Error(), resp.Status)
}

func TestWrapErr(t *testing.T) {
	br := bytes.NewReader([]byte("test not found"))
	resp := &http.Response{Status: "404 Not Found", Body: io.NopCloser(br)}
	r := httpx.Resp(resp)
	s := r.Str()
	assert.True(t, s == "")
	err0 := r.Err()
	assert.Equal(t, err0.Error(), resp.Status)
	err1 := fmt.Errorf("test error")
	httpx.WrapErr(r, err1)
	werr0 := r.Err()
	err2 := fmt.Errorf("read error")
	httpx.WrapErr(r, err2)
	werr1 := r.Err()
	err3 := fmt.Errorf("json error")
	httpx.WrapErr(r, err3)
	werr2 := r.Err()
	fmt.Println(r.Err())
	assert.True(t, errors.Is(werr2, err3))
	assert.True(t, errors.Is(werr1, err2))
	assert.True(t, errors.Is(werr0, err1))
}

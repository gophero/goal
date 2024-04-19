package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/gophero/goal/errorx"
)

var defaultClient = &http.Client{
	Timeout: time.Second * 60,
}

type H map[string]string

type ContentType string

type builder struct {
	client      *http.Client
	url         string
	params      []any
	method      string
	contentType ContentType
	headers     http.Header
	body        io.Reader
}

type R struct {
	*http.Response

	err  error
	body []byte
	read bool
}

func Resp(r *http.Response) *R {
	return &R{Response: r}
}

func RespErr(r *http.Response, err error) *R {
	return &R{Response: r, err: err}
}

func (r *R) ok() bool {
	if !(r.StatusCode >= http.StatusOK && r.StatusCode < http.StatusMultipleChoices) {
		r.wrapErr(fmt.Errorf("%s", r.Status))
		return false
	}
	return true
}

func (r *R) wrapErr(err error) {
	var e = err
	if r.err != nil {
		e = errorx.Wrapf(err, r.err.Error())
	}
	r.err = e
}

func (r *R) readAll() *R {
	if !r.read {
		if r.Body != nil {
			if bs, err := io.ReadAll(r.Body); err != nil {
				r.wrapErr(fmt.Errorf("read body error: %v", err))
				r.body = []byte{}
			} else {
				r.read = true
				r.body = bs
			}
		}
	}
	return r
}

func (r *R) Err() error {
	return r.err
}

func (r *R) Str() string {
	if r.ok() {
		return string(r.readAll().body)
	}
	return ""
}

func (r *R) Bytes() []byte {
	return r.readAll().body
}

func (r *R) JsonObj(v any) any {
	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		panic(fmt.Errorf("param v should be a pointer"))
	}
	if r.ok() {
		if err := json.Unmarshal(r.readAll().body, v); err != nil {
			r.wrapErr(fmt.Errorf("unmarshal json error: %v", err))
		}
	}
	return v
}

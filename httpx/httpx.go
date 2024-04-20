package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gophero/goal/errorx"
)

type Handler func(resp *http.Response)

type ErrHandler func(err error)

type httpBuilder struct {
	builder
	callback   Handler
	errHandler ErrHandler
}

func NewBuilderClient(c *http.Client, url string, param ...any) *httpBuilder {
	var h = http.Header{}
	var b = &httpBuilder{}
	b.client = c
	b.url = url
	b.params = param
	b.headers = h
	return b
}

func NewBuilder(url string, param ...any) *httpBuilder {
	return NewBuilderClient(defaultClient, url, param...)
}

func (b *httpBuilder) Client(c *http.Client) *httpBuilder {
	b.client = c
	return b
}

func (b *httpBuilder) ContentType(contentType ContentType) *httpBuilder {
	b.contentType = contentType
	return b
}

func (b *httpBuilder) Header(key string, v string) *httpBuilder {
	b.headers.Add(key, v)
	return b
}

func (b *httpBuilder) Headers(headers ...H) *httpBuilder {
	for _, h := range headers {
		for k, v := range h {
			b.headers.Add(k, v)
		}
	}
	return b
}

func (b *httpBuilder) Body(body io.Reader) *httpBuilder {
	b.body = body
	return b
}

func (b *httpBuilder) BodyStr(body string) *httpBuilder {
	b.body = strings.NewReader(body)
	return b
}

func (b *httpBuilder) getClient() *http.Client {
	if b.client == nil {
		return defaultClient
	}
	return b.client
}

func (b *httpBuilder) Request(method string) {
	b.method = method
	req, err := http.NewRequest(b.method, b.url, b.body)
	if len(b.headers) > 0 {
		req.Header = b.headers
	}
	resp, err := b.getClient().Do(req)

	if err != nil {
		b.errHandler(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
	return
}

func (b *httpBuilder) Get() {
	b.method = http.MethodGet
	var resp *http.Response
	var err error
	if len(b.headers) > 0 {
		req, err := http.NewRequest(b.method, b.url, nil)
		errorx.Throw(err)
		req.Header = b.headers
		resp, err = b.getClient().Do(req)
	} else {
		resp, err = b.getClient().Get(b.url)
	}

	if err != nil {
		b.errHandler(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
	return
}

func (b *httpBuilder) Post() {
	b.method = http.MethodPost
	var resp *http.Response
	var err error
	if len(b.headers) > 0 {
		req, err := http.NewRequest(b.method, b.url, b.body)
		errorx.Throw(err)

		b.mergeHeaders()

		req.Header = b.headers
		resp, err = b.getClient().Do(req)
	} else {
		resp, err = b.getClient().Post(b.url, string(b.contentType), b.body)
	}

	if err != nil {
		b.errHandler(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
	return
}

func (b *httpBuilder) mergeHeaders() *httpBuilder {
	if b.contentType != "" {
		b.headers.Add("Content-Type", string(b.contentType))
		// } else {
		// 	if len(b.headers) == 0 {
		// 		return b
		// 	}
		// 	var ct, ok = b.headers["Content-Type"]
		// 	if ok {
		// 		b.contentType = ContentType(ct[0]) // 取第一个
		// 	}
	}
	return b
}

func (b *httpBuilder) WhenSuccess(handler Handler) *httpBuilder {
	b.callback = handler
	return b
}

func (b *httpBuilder) WhenFailed(handler ErrHandler) *httpBuilder {
	b.errHandler = handler
	return b
}

// convenient GET methods

func MustGetString(url string, headers ...H) string {
	return GetString(url, func(err error) {
		panic(err)
	}, headers...)
}

func GetString(url string, errHandler ErrHandler, headers ...H) string {
	var s string
	NewBuilder(url).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		s = string(bs)
	}).WhenFailed(errHandler).Headers(headers...).Get()
	return s
}

func MustGetBytes(url string, headers ...H) []byte {
	return GetBytes(url, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, headers...)
}

func GetBytes(url string, errHandler ErrHandler, headers ...H) []byte {
	var ret []byte
	NewBuilder(url).WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read reponse data error: %v", err))
		}
		ret = bytes
	}).WhenFailed(errHandler).Headers(headers...).Get()
	return ret
}

func MustGet(url string, handler Handler, headers ...H) {
	Get(url, handler, func(err error) {
		panic(err)
	}, headers...)
}

func Get(url string, handler Handler, errHandler ErrHandler, headers ...H) {
	NewBuilder(url).WhenSuccess(handler).WhenFailed(errHandler).Headers(headers...).Get()
}

func MustGetJson[T any](url string, t T, headers ...H) T {
	GetJson(url, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, t, headers...)
	return t
}

func GetJson[T any](url string, errHandler ErrHandler, t T, headers ...H) T {
	NewBuilder(url).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		err = json.Unmarshal(bs, t)
		if err != nil {
			panic(fmt.Sprintf("unmarshal error: %v", err))
		}
	}).WhenFailed(errHandler).Headers(headers...).Get()
	return t
}

// convenient POST methods

func MustPostJson(url string, body io.Reader, headers ...H) string {
	return PostJson(url, body, func(err error) {
		panic(err)
	}, headers...)
}

func PostJson(url string, body io.Reader, errHandler ErrHandler, headers ...H) string {
	var s string
	NewBuilder(url).ContentType(ContentTypeApplicationJson).Body(body).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		s = string(bs)
	}).WhenFailed(errHandler).Headers(headers...).Post()
	return s
}

func MustPostForm(url string, body io.Reader, headers ...H) string {
	return PostForm(url, body, func(err error) {
		panic(err)
	}, headers...)
}

func PostForm(url string, body io.Reader, errHandler ErrHandler, headers ...H) string {
	var s string
	NewBuilder(url).ContentType(ContentTypeApplicationFormUrlencoded).Body(body).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		s = string(bs)
	}).WhenFailed(errHandler).Headers(headers...).Post()
	return s
}

func MustPostBytes(url string, ct ContentType, body io.Reader, headers ...H) []byte {
	return PostBytes(url, ct, body, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, headers...)
}

func PostBytes(url string, ct ContentType, body io.Reader, errHandler ErrHandler, headers ...H) []byte {
	var ret []byte
	NewBuilder(url).ContentType(ct).Body(body).WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read reponse data error: %v", err))
		}
		ret = bytes
	}).WhenFailed(errHandler).Headers(headers...).Post()
	return ret
}

func MustPost(url string, ct ContentType, body io.Reader, handler Handler, headers ...H) {
	Post(url, ct, body, handler, func(err error) {
		panic(err)
	}, headers...)
}

func Post(url string, ct ContentType, body io.Reader, handler Handler, errHandler ErrHandler, headers ...H) {
	NewBuilder(url).ContentType(ct).Body(body).WhenSuccess(handler).WhenFailed(errHandler).Headers(headers...).Post()
}

func MustPostJsonObject[T any](url string, ct ContentType, body io.Reader, t T, headers ...H) T {
	PostJsonObject(url, ct, body, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, t, headers...)
	return t
}

func PostJsonObject[T any](url string, ct ContentType, body io.Reader, errHandler ErrHandler, t T, headers ...H) T {
	NewBuilder(url).ContentType(ct).Body(body).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		err = json.Unmarshal(bs, t)
		if err != nil {
			panic(fmt.Sprintf("unmarshal error: %v", err))
		}
	}).WhenFailed(errHandler).Headers(headers...).Post()
	return t
}

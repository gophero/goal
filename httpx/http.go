package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gophero/goal/errorx"
)

var defaultClient = &http.Client{
	Timeout: time.Second * 60,
}

type H map[string]string

type Handler func(resp *http.Response)

type ErrHandler func(err error)

type ContentType string

type builder struct {
	client      *http.Client
	url         string
	params      []any
	method      string
	contentType ContentType
	headers     http.Header
	callback    Handler
	errHandler  ErrHandler
	body        io.Reader
}

func NewBuilderClient(c *http.Client, url string, param ...any) *builder {
	var h = http.Header{}
	return &builder{client: c, url: url, params: param, headers: h}
}

func NewBuilder(url string, param ...any) *builder {
	var h = http.Header{}
	return &builder{url: url, params: param, headers: h}
}

func (b *builder) Client(c *http.Client) *builder {
	b.client = c
	return b
}

func (b *builder) ContentType(contentType ContentType) *builder {
	b.contentType = contentType
	return b
}

func (b *builder) Header(key string, v string) *builder {
	b.headers.Add(key, v)
	return b
}

func (b *builder) Headers(headers ...H) *builder {
	for _, h := range headers {
		for k, v := range h {
			b.headers.Add(k, v)
		}
	}
	return b
}

func (b *builder) Body(body io.Reader) *builder {
	b.body = body
	return b
}

func (b *builder) BodyStr(body string) *builder {
	b.body = strings.NewReader(body)
	return b
}

func (b *builder) getClient() *http.Client {
	if b.client == nil {
		return defaultClient
	}
	return b.client
}

func (b *builder) Get() *builder {
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
		return b
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
	return b
}

func (b *builder) Post() *builder {
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
		return b
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		b.errHandler(err)
	} else {
		b.callback(resp)
	}
	return b
}

func (b *builder) mergeHeaders() *builder {
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

func (b *builder) WhenSuccess(handler Handler) *builder {
	b.callback = handler
	return b
}

func (b *builder) WhenFailed(handler ErrHandler) *builder {
	b.errHandler = handler
	return b
}

// convenient GET methods

func MustGetString(url string) string {
	return GetString(url, func(err error) {
		panic(err)
	})
}

func GetString(url string, errHandler ErrHandler) string {
	var s string
	NewBuilder(url).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		s = string(bs)
	}).WhenFailed(errHandler).Get()
	return s
}

func MustGetBytes(url string) []byte {
	return GetBytes(url, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	})
}

func GetBytes(url string, errHandler ErrHandler) []byte {
	var ret []byte
	NewBuilder(url).WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read reponse data error: %v", err))
		}
		ret = bytes
	}).WhenFailed(errHandler).Get()
	return ret
}

func MustGet(url string, handler Handler) {
	Get(url, handler, func(err error) {
		panic(err)
	})
}

func Get(url string, handler Handler, errHandler ErrHandler) {
	NewBuilder(url).WhenSuccess(handler).WhenFailed(errHandler).Get()
}

func MustGetJson[T any](url string, t T) T {
	GetJson(url, func(err error) {
		panic(fmt.Sprintf("request failed: %v", err))
	}, t)
	return t
}

func GetJson[T any](url string, errHandler ErrHandler, t T) T {
	NewBuilder(url).WhenSuccess(func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Sprintf("read response data error: %v", err))
		}
		err = json.Unmarshal(bs, t)
		if err != nil {
			panic(fmt.Sprintf("unmarshal error: %v", err))
		}
	}).WhenFailed(errHandler).Get()
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

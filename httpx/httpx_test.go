package httpx_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/gophero/goal/httpx"
	"github.com/gophero/goal/testx"
)

const (
	html    = `<html><head>test</head><body><h1>test page!</h1></body></html>`
	jsonstr = `{"name":"张三","age":20,"height":70.5}`
)

var srv *http.Server

// startServer 启动模拟 api 的 http 服务
func startServer() *http.Server {
	if srv != nil {
		return srv
	}
	port := "1234"
	log.Println("starting http server on port: " + port)
	srv = &http.Server{Addr: ":" + port}
	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(html))
		if err != nil {
			return
		}
	})
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(jsonstr))
		if err != nil {
			return
		}
	})
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Println("http server shutdown")
			} else {
				log.Printf("start http server error: %v", err)
			}
		} else {
			fmt.Println("http server started on port: " + port)
		}
	}()
	return srv
}

// ========== GET method tests ==========

func TestHttp_BuilderGet(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Title("test send http GET method with builder api")
	logger.Case("GET html from baidu home page.")
	httpx.NewBuilder("http://www.baidu.com").WhenSuccess(func(resp *http.Response) {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fail("should can read response data, but got error: %v", err)
		} else {
			logger.Pass("should has no error")
		}
		if body == nil {
			logger.Fail("should has body but not found")
		} else {
			logger.Pass("should has body")
		}
	}).WhenFailed(func(err error) {
		logger.Fail("should has no error but found: %v", err)
	}).Get()

	logger.Case("GET html from localhost:1234.")
	httpx.NewBuilder("http://localhost:1234/html").WhenSuccess(func(resp *http.Response) {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fail("response body should be readable but not: %v", err)
		} else {
			logger.Pass("response body should be readable")
		}
		if string(body) != html {
			logger.Fail("should return the correct html")
		} else {
			logger.Pass("should return the correct html")
		}
	}).WhenFailed(func(err error) {
		logger.Fail("should has not error but found: %v", err)
	}).Get()

	logger.Case("GET json from localhost:1234.")
	httpx.NewBuilder("http://localhost:1234/json").WhenSuccess(func(resp *http.Response) {
		jsonToUser(resp.Body, logger)
	}).WhenFailed(func(err error) {
		logger.Fail("should has not error but found: %v", err)
	}).Get()

	// NewBuilder("http://localhost:1234/json").ReadResp().WhenFailed(func(err error, resp *http.Response) {
	// 	t.Fatal(err)
	// }).Get()

	logger.Case("GET unknown url from localhost:1234.")
	httpx.NewBuilder("http://localhost:1234/unknown").WhenSuccess(func(resp *http.Response) {
		if resp != nil {
			logger.Fail("should be nil response but not")
		}
	}).WhenFailed(func(err error) {
		if err == nil {
			logger.Fail("expect occur an error but no one")
		} else {
			logger.Pass("get an expected error: %v", err)
		}
	}).Get()
}

func jsonToUser(r io.Reader, logger *testx.Logger) {
	body, err := io.ReadAll(r)
	if err != nil {
		logger.Fail("response body should be readable but not: %v", err)
	} else {
		logger.Pass("response body should be readable")
	}
	type user struct {
		Name   string  `json:"name"`
		Age    int     `json:"age"`
		Height float32 `json:"height"`
	}
	var u = new(user)
	err = json.Unmarshal(body, u)
	if err != nil {
		logger.Fail("response body should be able to unmarshal to struct but failed: %v", err)
	} else {
		logger.Pass("response body should be able to unmarshal to struct")
	}
	if u.Age == 0 || u.Name == "" {
		logger.Fail("response body unmarshalled to struct successfully but some fields have incorrect value")
	} else {
		logger.Pass("response body unmarshalled to struct successfully")
	}
}

func TestHttp_GetString(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Case("using GetString method to get html from localhost:1234")
	s := httpx.GetString("http://localhost:1234/html", func(err error) {
		if err != nil {
			logger.Fail("should has no error but found: %v", err)
		} else {
			logger.Pass("request should be successful")
		}
	})
	logger.Require(s == html, "response html should match the dest")

	logger.Case("using GetString method to get string from unknown url, should be failed")
	httpx.GetString("http://localhost:1234/unknown", func(err error) {
		logger.Require(err != nil, "request should be unsuccessful: %v", err)
	})
}

func TestHttp_MustGetString(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Title("test MustGetString method")

	logger.Case("positive case: get json from localhost:1234")
	s := httpx.MustGetString("http://localhost:1234/json")
	logger.Require(s == jsonstr, "response string should be correct")

	logger.Case("negative case: get json from unknown url, should occur and error")
	func() {
		defer func() {
			err := recover()
			logger.Require(err != nil, "should occur an error: %v", err)
		}()
		_ = httpx.MustGetString("http://localhost:1234/unknown")
	}()
}

func TestHttp_MustGet(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Case("using MustGet method to get json string from localhost:1234")
	httpx.MustGet("http://localhost:1234/json", func(resp *http.Response) {
		jsonToUser(resp.Body, logger)
	})

	logger.Case("using MustGet method to get unknown url from localhost:1234, should be failed")
	// 包装为匿名函数，以便能够捕获异常
	func() {
		// 捕获异常
		defer func() {
			err := recover()
			logger.Require(err != nil, "request should be failed: %v", err)
		}()
		httpx.MustGet("http://localhost:1234/unknown", func(resp *http.Response) {})
	}()
}

func TestHttp_Get(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Case("using Get method to get json string from localhost:1234")
	httpx.Get("http://localhost:1234/json", func(resp *http.Response) {
		jsonToUser(resp.Body, logger)
	}, func(err error) {
		logger.Require(err == nil, "should not occur an error")
	})
	logger.Pass("request should be successful")

	logger.Case("using Get method to get unknown url from localhost:1234, should be failed")
	func() {
		defer func() {
			err := recover()
			logger.Require(err != nil, "request should be failed: %v", err)
		}()
		httpx.Get("http://localhost:1234/unknown", func(resp *http.Response) {}, func(err error) {
			panic(err) // if has an error, panic it
		})
	}()
}

func TestHttp_GetBytes(t *testing.T) {
	logger := testx.Wrap(t)
	logger.Title("using GetBytes or MustGetBytes method to get bytes")
	logger.Case("request an image from baidu")
	url := "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"
	bs := httpx.GetBytes(url, func(err error) {
		logger.Require(err == nil, "request should be successful with no error")
	})
	logger.Pass("request should be successful")
	logger.Require(len(bs) > 0, "response data should has content")

	logger.Case("request an image from baidu")
	bs = httpx.MustGetBytes(url)
	logger.Pass("request should be successful")
	logger.Require(len(bs) > 0, "response data should has content")
}

func TestHttp_GetJsonObject(t *testing.T) {
	startServer()

	type user struct {
		Name   string  `json:"name"`
		Age    int     `json:"age"`
		Height float32 `json:"height"`
	}

	logger := testx.Wrap(t)
	logger.Title("using GetJsonObject or MustGetJsonObject method to get object from json response")

	logger.Case("GetJsonObject: request json from localhost:1234 and unmarshal it to user struct")
	u := httpx.GetJson("http://localhost:1234/json", func(err error) {
		logger.Require(err == nil, "request should be successful")
	}, &user{})
	logger.Require(u != nil, "request should be successful")
	bs, _ := json.Marshal(u)
	logger.Require(string(bs) == jsonstr, "return user should be correct")

	logger.Case("MustGetJsonObject: request json from localhost:1234 and unmarshal it to user struct")
	u = httpx.MustGetJson("http://localhost:1234/json", &user{})
	bs, _ = json.Marshal(u)
	logger.Require(u != nil, "request should be successful")
	logger.Require(string(bs) == jsonstr, "return user should be correct")
}

// ========== POST method tests ==========

func TestSimplePost(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Title("test basic post request")

	// Post请求用于向服务器发送资源，这里的查询并不符合restful规范

	logger.Case("simplest post html")
	var s string
	httpx.NewBuilder("http://localhost:1234/html").WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fail("should can read response data, but got error: %v", err)
		} else {
			logger.Pass("should has no error")
		}
		if bytes == nil {
			logger.Fail("should has body but not found")
		} else {
			logger.Pass("should has body")
		}
		s = string(bytes)
	}).WhenFailed(func(err error) {
		logger.Require(err == nil, "request should be successful")
	}).Post()
	logger.Require(html == s, "post result with string should be correct")

	logger.Case("simplest post json")
	var js string
	httpx.NewBuilder("http://localhost:1234/json").WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fail("should can read response data, but got error: %v", err)
		} else {
			logger.Pass("should has no error")
		}
		if bytes == nil {
			logger.Fail("should has body but not found")
		} else {
			logger.Pass("should has body")
		}
		js = string(bytes)
	}).WhenFailed(func(err error) {
		logger.Require(err == nil, "request should be successful")
	}).Post()
	logger.Require(js == jsonstr, "post result with string should be correct")
}

func TestRequest(t *testing.T) {
	startServer()

	logger := testx.Wrap(t)
	logger.Title("test request")

	logger.Case("test post html")
	var s string
	httpx.NewBuilder("http://localhost:1234/html").WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fail("should can read response data, but got error: %v", err)
		} else {
			logger.Pass("should has no error")
		}
		if bytes == nil {
			logger.Fail("should has body but not found")
		} else {
			logger.Pass("should has body")
		}
		s = string(bytes)
	}).WhenFailed(func(err error) {
		logger.Require(err == nil, "request should be successful")
	}).ContentType(httpx.ContentTypeApplicationFormUrlencoded).Request(http.MethodGet)
	logger.Require(html == s, "post result with string should be correct")

	logger.Case("post json")
	var js string
	httpx.NewBuilder("http://localhost:1234/json").WhenSuccess(func(resp *http.Response) {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Fail("should can read response data, but got error: %v", err)
		} else {
			logger.Pass("should has no error")
		}
		if bytes == nil {
			logger.Fail("should has body but not found")
		} else {
			logger.Pass("should has body")
		}
		js = string(bytes)
	}).WhenFailed(func(err error) {
		logger.Require(err == nil, "request should be successful")
	}).ContentType(httpx.ContentTypeApplicationJson).Request(http.MethodPost)

	logger.Require(js == jsonstr, "post result with string should be correct")
}

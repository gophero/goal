package httpx_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gophero/goal/httpx"
)

// 运行测试时会执行 example 代码，示例方法必须使用 fmt.println 输出，最后约定用 Output 输出一致的结果
// 打印结果，如后边的 Output 输出结果必须一致，否则测试时会失败

func ExampleNewBuilder() {
	// 启动示例 http 服务
	startServer()

	var ret string
	// 构建 get 请求
	httpx.NewBuilder("http://localhost:1234/html").WhenSuccess(func(resp *http.Response) { // 请求成功处理
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("response body should be readable but not: %v", err)
		}
		ret = string(body)
	}).WhenFailed(func(err error) { // 请求失败处理
		panic(err)
	}).Get() // 请求完成，清理资源

	fmt.Println(ret)
	// Output:
	// <html><head>test</head><body><h1>test page!</h1></body></html>
}

func ExampleGetString() {
	startServer()

	s := httpx.GetString("http://localhost:1234/json", func(err error) {
		if err != nil {
			log.Fatalf("should has no error but found: %v", err)
		}
	})
	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGetString() {
	startServer()

	s := httpx.MustGetString("http://localhost:1234/json")
	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGet() {
	startServer()

	var s string
	httpx.MustGet("http://localhost:1234/json", func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		s = string(bs)
	})
	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleGet() {
	startServer()

	var s string
	httpx.Get("http://localhost:1234/json", func(resp *http.Response) {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		s = string(bs)
	}, func(err error) {
		panic(err)
	})

	fmt.Println(s)
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleGetBytes() {
	startServer()

	bs := httpx.GetBytes("http://localhost:1234/json", func(err error) {
		panic(err)
	})
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGetBytes() {
	startServer()

	bs := httpx.MustGetBytes("http://localhost:1234/json")
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

type user struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
}

func ExampleGetJson() {
	startServer()

	u := httpx.GetJson("http://localhost:1234/json", func(err error) {
		panic(err)
	}, &user{})

	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

func ExampleMustGetJson() {
	startServer()

	u := httpx.MustGetJson("http://localhost:1234/json", &user{})

	bs, _ := json.Marshal(u)
	fmt.Println(string(bs))
	// Output:
	// {"name":"张三","age":20,"height":70.5}
}

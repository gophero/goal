package errorx

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDo(t *testing.T) {
	h := NewHandler()
	defer h.Done(func(err error) {
		if h.HasErr() {
			fmt.Printf("done, error: %v\n", h.Err())
		} else {
			fmt.Println("done, no error")
		}
	})
	h.Do(func() error {
		if rand.Intn(100)%3 == 0 {
			return fmt.Errorf("make and error1")
		}
		return nil
	}).Do(func() error {
		if rand.Intn(200)%3 == 0 {
			return fmt.Errorf("make an error2")
		}
		return nil
	})
	fmt.Println("has error:", h.HasErr())
}

func ExampleNewHandler() {
	h := NewHandler()
	fmt.Println(h.HasErr())
	fmt.Println(h.Err())
}

func ExampleErrorHandler_Do() {
	// 重复处理 error
	i, err := f1()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	j, err := f2()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("i= %d, j = %d\n", i, j)

	// 使用 ErrorHandler
	var x, y int
	h := NewHandler()
	h.Do(func() error {
		x, err = f1()
		return err
	})
	h.Do(func() error {
		y, err = f2()
		return err
	})
	if h.HasErr() {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("x = %d, y = %d\n", x, y)
}

func f1() (int, error) {
	n := rand.Intn(100)
	if n%3 == 0 {
		return -1, fmt.Errorf("make an error1")
	}
	return n, nil
}

func f2() (int, error) {
	n := rand.Intn(200)
	if n%3 == 0 {
		return -2, fmt.Errorf("make an error2")
	}
	return n, nil
}

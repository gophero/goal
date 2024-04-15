package random_test

import (
	"fmt"
	"github.com/gophero/goal/random"
	"github.com/gophero/goal/testx"
	"testing"
)

func TestInt(t *testing.T) {
	tr := testx.Wrap(t)
	tr.Case("give 1, should get 0")
	r := random.Int(1)
	tr.Require(r == 0, "result should be 0")

	for _, i := range []int{-1, 0} {
		tr.Case("give %d, should panic", i)
		func() {
			defer func() {
				err := recover()
				tr.Require(err != nil, "should panic")
			}()
			random.Int(i)
		}()
	}

	for _, i := range []int{2, 10, 99, 999, 9999} {
		tr.Case("give %d, should get a random number which < %d", i, i)
		tr.Require(random.Int(i) < i, "result should < %d", i)
	}
}

// 使用 fuzzing 测试随机数
// go test -run=FuzzInt 只测试提供的种子而不进行模糊测试
// go test -fuzz=FuzzInt 进行模糊测试，通过 -timeout=10s 指定执行时间
func FuzzInt(f *testing.F) {
	// 添加一个种子集，让模糊测试可以根据种子生成数据
	seeds := []int{1, 100, 999}
	for _, s := range seeds {
		f.Add(s)
	}
	// 执行模糊测试，第一个参数为 *testing.T，后边的参数为由种子生成的参数
	f.Fuzz(func(t *testing.T, max int) {
		if max > 0 { // 过滤掉非法的参数
			fmt.Println(max)
			i := random.Int(max)
			if !(i >= 0 && i < max) {
				f.Errorf("test failed, result should be >= 0 and <= %d, but is %d", max, i)
			}
		}
	})
}

func TestBetween(t *testing.T) {
	tr := testx.Wrap(t)
	tr.Case("give 0,0, get 0")
	r := random.Between(0, 0)
	tr.Require(r == 0, "should get 0")

	cases := [][]int{{-1, -1}, {-1, 0}, {0, -1}}
	for _, c := range cases {
		tr.Case("give invalid param: %#v, will panic", c)
		func() {
			defer func() {
				err := recover()
				tr.Require(err != nil, "should panic")
			}()
			random.Between(c[0], c[1])
		}()
	}

	cases = [][]int{{0, 1}, {1, 100}, {100, 9999}}
	for _, c := range cases {
		tr.Case("give param: %#v, will get correct result", c)
		r := random.Between(c[0], c[1])
		tr.Require(r >= c[0] && r < c[1], "result should >= %d < %d", c[0], c[1])
	}
}

func FuzzBetween(f *testing.F) {
	f.Add(-1, 0)
	f.Add(0, -1)
	f.Add(0, 0)
	f.Add(1, 1)
	f.Add(1, 100)
	f.Add(100, 999)
	f.Fuzz(func(t *testing.T, min int, max int) {
		if min >= 0 && max >= 0 && min <= max {
			i := random.Between(min, min)
			if min == max {
				if i != min {
					t.Errorf("test failed, result should = %d, but is %d", min, i)
				}
			} else if !(i >= min && i < max) {
				t.Errorf("test failed, result should be >= %d and < %d, but is %d", min, max, i)
			}
		}
	})
}

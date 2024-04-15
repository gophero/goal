package slicex_test

import (
	"fmt"
	"github.com/gophero/goal/collection/slicex"
	"github.com/gophero/goal/testx"
	"reflect"
	"testing"
)

var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestNew(t *testing.T) {
	s := slicex.New[int]()
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	// s = append(s, 5)
	fmt.Println(len(s), cap(s))

	s = slicex.NewSize[int](5)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	fmt.Println(len(s), cap(s))
}

func TestRetain(t *testing.T) {
	logger := testx.Wrap(t)

	ret := slicex.Wrap(testSlice).Retain(func(a int) bool {
		return a > 5
	}).Raw()
	fmt.Println(testSlice)
	expect := []int{6, 7, 8, 9}
	logger.Require(reflect.DeepEqual(expect, ret), "expect %#v, actual %#v", expect, ret)

	// ret1 := slice.Wrap([]any{"a", 1, 3.14}).Retain(func(a any) bool {
	// 	switch a.(type) {
	// 	case int:
	// 		return a.(int) > 1
	// 	default:
	// 		return true
	// 	}
	// })
	// expect1 := []any{"a", 3.14}
	// logger.Require(reflect.DeepEqual(expect1, ret1), "expect %v, actual %v", expect1, ret1)
}

func TestJoin(t *testing.T) {
	logger := testx.Wrap(t)

	s := slicex.Wrap(testSlice).Join(",")
	logger.Require(s == "1,2,3,4,5,6,7,8,9", "%v join result should be %s", testSlice, s)

	join := slicex.Wrap([]string{"a", "b"}).Join(".")
	logger.Require(join == "a.b", "a join b should be %s", join)

	// join = slice.Wrap([]TestData{"a", 1, "3.14"}).Join(",")
	// logger.Require(join == "a,1,3.14", "%s join %d join %.2f should be %s", "a", 1, 3.14, join)
}

func TestUnion(t *testing.T) {
	var before = testSlice

	logger := testx.Wrap(t)
	sl := []int{1, 2, 3, 4, 5, 6, 10, 11}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ret := slicex.Wrap(testSlice).Union(sl).Raw()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestIntersect(t *testing.T) {
	var before = testSlice

	logger := testx.Wrap(t)
	sl := []int{1, 2, 3, 10, 11}
	want := []int{1, 2, 3}
	ret := slicex.Wrap(testSlice).Intersect(sl).Raw()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestRemove(t *testing.T) {
	var before = testSlice

	logger := testx.Wrap(t)
	sl := []int{1, 2, 3, 10, 11}
	want := []int{4, 5, 6, 7, 8, 9}
	ret := slicex.Wrap(testSlice).Remove(sl).Raw()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestDiff(t *testing.T) {
	var before = testSlice

	logger := testx.Wrap(t)
	sl := []int{1, 2, 3, 10, 11}
	want := []int{4, 5, 6, 7, 8, 9, 10, 11}
	ret := slicex.Wrap(testSlice).Diff(sl).Raw()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestDelete(t *testing.T) {
	raw := testSlice

	logger := testx.Wrap(t)
	want := []int{4, 5, 6, 7, 8, 9}
	var ret = slicex.Wrap(testSlice).Delete(1, 2, 3).Raw()
	logger.Require(reflect.DeepEqual(testSlice, raw), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")

	want = []int{2, 4, 6, 8}
	ret = slicex.Wrap(testSlice).Delete(1, 3, 5, 7, 9).Raw()
	logger.Require(reflect.DeepEqual(testSlice, raw), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")
}

func TestRemoveDuplicate(t *testing.T) {
	logger := testx.Wrap(t)

	raw := []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7}
	ret := slicex.Wrap(raw).RemoveDuplicate().Raw()
	want := []int{1, 2, 3, 4, 5, 6, 7}
	logger.Require(reflect.DeepEqual(raw, []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7}), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")
}

func TestChainInvoke(t *testing.T) {
	logger := testx.Wrap(t)

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// not use Raw()
	ret := slicex.Wrap(testSlice).Union([]int{1, 2, 3, 4, 10}).RemoveDuplicate()
	logger.Require(slicex.Equal(testSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}), "raw slice should not be changed")
	logger.Require(slicex.Equal(ret, want), "correct result")

	want = []int{1, 2, 3, 10}
	// use Raw()
	ret1 := slicex.Wrap(testSlice).Union([]int{1, 2, 3, 4, 10}).Intersect([]int{1, 2, 3, 10, 11, 12}).Raw()
	logger.Require(reflect.DeepEqual(testSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret1, want), "correct result")
}

func TestSort(t *testing.T) {
	var ss = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	s := slicex.Wrap(ss)
	sr := s.Sort(func(i, j int) bool {
		if i < j {
			return true
		}
		return false
	})
	fmt.Printf("%v\n", ss)
	sr.Reverse()
	fmt.Printf("%v\n", ss)
}

type user struct {
	name  string
	age   int
	score float32
}

func TestSortObj(t *testing.T) {
	users := []user{
		{name: "huzhou", age: 18, score: 99.5},
		{name: "huzhou", age: 16, score: 100},
		{name: "zhangsan", age: 17, score: 99.5},
		{name: "abbc", age: 17, score: 99.5},
	}

	s := slicex.Wrap(users)
	sr := s.Sort(func(a, b user) bool {
		if a.score != b.score {
			return a.score > b.score
		}
		if a.age != b.age {
			return a.age < b.age
		}
		if a.name != b.name {
			return a.name < b.name
		}
		return false
	})
	fmt.Printf("%v\n", s)
	sr.Reverse()
	fmt.Printf("%v\n", s)
}

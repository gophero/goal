package mapx_test

import (
	"github.com/gophero/goal/collection/mapx"
	"github.com/gophero/goal/testx"
	"testing"
)

func TestMap(t *testing.T) {
	logger := testx.Wrap(t)

	intMap := mapx.New[string, int]()
	intMap.Put("a", 1).Put("b", 2).Put("c", 3)
	v := intMap.Get("a")
	logger.Require(v == 1, "value of map should be %d", v)
	keys1 := intMap.Keys()
	logger.Require(Equal(keys1, []string{"a", "b", "c"}), "keys of map should be equals to %v", keys1)
	val1 := intMap.Values()
	logger.Require(Equal(val1, []int{1, 2, 3}), "values of map should be equals to %v", val1)
	del1 := intMap.Del("a")
	logger.Require(del1 == 1, "del value of map should be %v", del1)
	intMap.Put("b", 3)
	upv1 := intMap.Get("b")
	logger.Require(upv1 == 3, "update value of map should be %v", upv1)

	intMap1 := mapx.New[int, int]()
	intMap1.Put(1, 1).Put(2, 2).Put(3, 3)
	vv := intMap1.Get(1)
	logger.Require(vv == 1, "value of map should be %d", vv)
	keys2 := intMap1.Keys()
	logger.Require(Equal(keys2, []int{1, 2, 3}), "keys of map should be equals to %v", keys2)
	val2 := intMap1.Values()
	logger.Require(Equal(val2, []int{1, 2, 3}), "values of map should be equals to %v", val2)

	stringMap := mapx.New[string, string]()
	stringMap.Put("a", "1").Put("b", "2").Put("c", "3")
	v1 := stringMap.Get("a")
	logger.Require(v1 == "1", "value of map should be %s", v1)
	keys3 := stringMap.Keys()
	logger.Require(Equal(keys3, []string{"a", "b", "c"}), "keys of map should be be equals to %v", keys3)
	val3 := stringMap.Values()
	logger.Require(Equal(val3, []string{"1", "2", "3"}), "values of map should be equals to %v", val3)

	anyMap := mapx.New[string, any]()
	anyMap.Put("a", 1).Put("b", "b").Put("c", 3.14)
	v2 := anyMap.Get("a").(int)
	logger.Require(v2 == 1, "value of map should be %s", v2)
	keys4 := anyMap.Keys()
	logger.Require(Equal(keys4, []string{"a", "b", "c"}), "keys of map should be be equals to %v", keys4)
	val4 := anyMap.Values()
	// Equal(val4, []any{1, "b", 3.14}) 这里不能用==比较了，any没有实现comparable
	logger.Require(len(val4) == 3, "values of map should be equals to %v", val4)
}

func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, c := range a {
		var find bool
		for _, c2 := range b {
			if c == c2 {
				find = true
			}
		}
		if !find {
			return false
		}
	}
	return true
}

package prob_test

import (
	"fmt"
	"github.com/gophero/goal/prob"
	"math"
	"testing"
)

func TestPercentProb0(t *testing.T) {
	for i := 0; i < 10000; i++ {
		if prob.Percent(0) {
			t.Fatal("test failed, want always not hit, but hit")
		}
	}
}

func TestPercentProb100(t *testing.T) {
	for i := 0; i < 10000; i++ {
		if !prob.Percent(100) {
			t.Fatal("test failed, want always hit, but does not hit")
		}
	}
}

func TestPercentProb(t *testing.T) {
	var x, y int
	var cnt = 100000 // 执行次数
	var want = 30    // 30%的概率
	for i := 0; i < cnt; i++ {
		if prob.Percent(want) {
			x++
		} else {
			y++
		}
	}
	// 误差
	var p = 0.2
	var wantProb = float64(want) / 100
	var actProb = float64(x) / float64(cnt)
	fmt.Printf("final hit prob: %.2f\n", actProb)
	fmt.Printf("final not hit prob: %.2f\n", float64(y)/float64(cnt))
	assertResult(wantProb, actProb, p, t)
}

func TestHalfProb(t *testing.T) {
	var x, y, cnt = 0, 0, 100000
	for i := 0; i < cnt; i++ {
		if prob.Half() {
			x++
		} else {
			y++
		}
	}

	var p = 0.2
	var wantProb = 0.5
	var actProb = float64(x) / float64(cnt)
	fmt.Printf("final hit prob: %.2f\n", actProb)
	fmt.Printf("final not hit prob: %.2f\n", float64(y)/float64(cnt))
	assertResult(wantProb, actProb, p, t)
}

func TestSelectProb0(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Errorf("test failed, error is excepted but no error")
		} else {
			fmt.Printf("expected error: %v\n", e)
			t.Skip()
		}
	}()
	idx := prob.Select([]int{})
	fmt.Printf("select index is %v", idx)
}

func TestSelectProb(t *testing.T) {
	createSelectProbCase([]int{0, 1, 0, 0}, t)
	println("====")
	createSelectProbCase([]int{4, 1, 0, 0}, t)
	println("====")
	createSelectProbCase([]int{1, 1, 1, 1, 1}, t)
	println("====")
	createSelectProbCase([]int{2, 2, 2, 1, 1, 1, 1}, t)
	println("====")
	createSelectProbCase([]int{30, 20, 10, 40}, t)
	println("====")
	createSelectProbCase([]int{20, 100, 200, 300, 400}, t)
}

func ExampleSelect() {
	prob.Select([]int{4, 1, 0, 0})
	prob.Select([]int{4, 1, 0, 0, 1, 2})
	prob.Select([]int{20, 30, 0, 50})
}

func createSelectProbCase(data []int, t *testing.T) {
	var length = len(data)
	ret := make([]int, length)
	cnt := 1000000
	var num int
	for _, v := range data {
		num += v
	}
	for i := 0; i < cnt; i++ {
		idx := prob.Select(data)
		ret[idx]++
	}

	var p = 0.1
	var wantProbs = make([]float64, length)
	var actProbs = make([]float64, length)
	for idx, _ := range data {
		wantProbs[idx] = float64(data[idx]) / float64(num)
	}
	for idx, _ := range ret {
		actProbs[idx] = float64(ret[idx]) / float64(cnt)
		fmt.Printf("final hit prob: %.2f, v: %v\n", actProbs[idx], data[idx])
		assertResult(wantProbs[idx], actProbs[idx], p, t)
	}
}

func assertResult(want, act, er float64, t *testing.T) {
	if math.Abs(want-act) > er {
		t.Errorf("test failed, want prob: %v, but is: %v", want, act)
	}
}

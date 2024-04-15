package prob

import (
	"fmt"
	"github.com/gophero/goal/assert"
	"math/rand"
	"time"
)

// Percent method calculates the percentage probability. Define a number within 100,
// then calculate the percentage probability, return true if hit, otherwise return false.
func Percent(r int) bool {
	assert.Require(r >= 0 && r <= 100, "invalid parameter")
	if r == 0 {
		return false
	}
	if r == 100 {
		return true
	}
	rd := rand.NewSource(time.Now().UnixNano())
	n := rand.New(rd).Intn(100)
	return n <= r
}

// Half method calculates the percentage probability that both of which hit or not hit is 50%.
func Half() bool {
	rd := rand.NewSource(time.Now().UnixNano())
	return rand.New(rd).Intn(2) == 1
}

// Select is a method which calculates the proportion of each element in the given int slice in entire of it, the order
// of the given slice is unconcerned.
//
// This method first calculates the total value of the entire slice, and then calculates a random number based on it,
// then it matches the random number with each element in the slice in order.
// If the match is successful, it returns the index of the slice, otherwise it will panic.
func Select(is []int) int {
	var num int
	for _, v := range is {
		if v < 0 {
			panic(fmt.Sprintf("invalid parameter, not negative value is required: %d", v))
		}
		num += v
	}
	if num == 0 {
		panic("invalid parameter, total value is zero")
	}
	rd := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rd).Intn(num)
	rate := 0
	for i, v := range is {
		rate += v
		if r < rate {
			return i
		}
	}
	panic("unknown error")
}

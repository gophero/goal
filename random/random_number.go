package random

import (
	"math/rand"
	"time"
)

// Int is a convenient method for rand.Intn(n int), but every invoking it will set time.Now().UnixNano()
// as it's seed. Obviously, the max parameter will be great than 0.
func Int(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

// Between will generate a random number between min(include) and max(exclude), and min must less than or equals max,
// otherwise it will panic.
func Between(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch {
	case min > max:
		panic("min must be less than or equal to max")
	case min < 0 || max < 0:
		panic("invalid params: both min and max must be >= 0")
	case min == max:
		return min
	}
	return min + r.Intn(max-min)
}

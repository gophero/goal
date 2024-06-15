package mathx

import (
	"math/rand"
	"testing"

	"github.com/gophero/goal/constraintx"
	"github.com/stretchr/testify/assert"
)

func TestMulBigFloat(t *testing.T) {
	// m1 := web3.ParseBigFloat()
	// fmt.Println(conv.StringToFloat64("0.1"))
	// m2 := web3.ParseBigFloat(conv.FloatToString(0.2))
	// fmt.Println(MulBigFloat(m1, m2))
}

func TestMaxn(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, maxn(ints...), Maxn(ints...))
	for i := 0; i < 1000; i++ {
		ints = randomInts()
		assert.Equal(t, maxn(ints...), Maxn(ints...))
	}

	ints = []int{1, 1, 1, 1, 1}
	assert.Equal(t, 1, Maxn(ints...))

	ints = []int{-1, -1, -1, -1}
	assert.Equal(t, -1, Maxn(ints...))
}

func maxn[T constraintx.Number](ns ...T) T {
	var max, tmp T
	for i := 0; i < len(ns); i++ {
		tmp = ns[i]
		if max < tmp {
			max = tmp
		}
	}
	return max
}

// go test -bench="Maxn(1)?$" -count=10 -cpu=2 -benchmem
func BenchmarkMaxn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			Maxn(randomInts()...)
		}
	}
}

func BenchmarkMaxn1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			maxn(randomInts()...)
		}
	}
}

func randomInts() []int {
	var ints []int
	for i := 0; i < 1000; i++ {
		ints = append(ints, rand.Int())
	}
	return ints
}

func TestMinn(t *testing.T) {
	ints := []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, Minn(ints...), -1)
	assert.Equal(t, minn(ints...), -1)
	assert.Equal(t, minn(ints...), Minn(ints...))
	for i := 0; i < 1000; i++ {
		ints = randomInts()
		assert.Equal(t, minn(ints...), Minn(ints...))
	}

	ints = []int{1, 1, 1, 1, 1}
	assert.Equal(t, 1, Minn(ints...))

	ints = []int{-1, -1, -1, -1}
	assert.Equal(t, -1, Minn(ints...))
}

func minn[T constraintx.Number](ns ...T) T {
	var min, tmp T = ns[0], ns[0]
	for i := 0; i < len(ns); i++ {
		tmp = ns[i]
		if min > tmp {
			min = tmp
		}
	}
	return min
}

package random_test

import (
	"github.com/gophero/goal/random"
	"github.com/gophero/goal/testx"
	"math/rand"
	"testing"
	"time"
)

func TestRandomAlphabetic(t *testing.T) {
	tl := testx.Wrap(t)
	tl.Case("loop 10 times to generate random string")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Alphabetic(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomNumber(t *testing.T) {
	tl := testx.Wrap(t)
	tl.Case("loop 10 times to generate random number as string")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Numeric(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomAlphanumeric(t *testing.T) {
	tl := testx.Wrap(t)
	tl.Case("loop 10 times to generate random Alphanumeric")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Alphanumeric(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

func TestRandomHex(t *testing.T) {
	tl := testx.Wrap(t)
	tl.Case("loop 10 times to generate random hex string")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(32) + 1
		s := random.Hex(n)
		tl.Require(n == len(s), "length of generated string should be %d", n)
	}
}

package errorx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	fmt.Println(ServerBusy)
	fmt.Println(NewPreferredErr(ServerBusy))
}

func TestIsPreferred2(t *testing.T) {
	err := fmt.Errorf("demo error")
	b := IsPreferred(err)
	fmt.Println(b) // false
	err = NewPreferredErr(err)
	fmt.Println(IsPreferred(err)) // true
}

func TestIsPreferred(t *testing.T) {
	err := fmt.Errorf("haha")
	b := IsPreferred(err) // false
	assert.True(t, !b)

	err = NewPreferredErr(fmt.Errorf("haha"))
	b = IsPreferred(err)
	assert.True(t, b)

	err = ServerBusy
	b = IsPreferred(err)
	assert.True(t, b)
}

func ExampleIsPreferred() {
	err := fmt.Errorf("demo error")
	b := IsPreferred(err)
	fmt.Println(b) // false
	err = NewPreferredErr(err)
	fmt.Println(IsPreferred(err)) // true
}

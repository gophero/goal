package reflectool_test

import (
	"fmt"
	"github.com/gophero/goal/reflectool"
	"testing"
)

type Inter interface {
	M1()
	M2()
}

type InterImpl struct {
}

// 非指针接收者
func (r InterImpl) M1() {

}

// 指针接收者
func (r *InterImpl) M2() {

}

func TestPrintMethodSet(t *testing.T) {
	var ty InterImpl
	var pty *InterImpl
	reflectool.PrintMethodSet(&ty)
	reflectool.PrintMethodSet(&pty)
	// nil interface
	reflectool.PrintMethodSet((*Inter)(nil))
}

type Inter2 interface {
	M3()
}

type GenericInter[T any] interface {
	Inter
	Print(t T)
}

type GenericImpl[T any] struct {
	InterImpl
}

func (r GenericImpl[T]) M3() {
}

func (r *GenericImpl[T]) Print(t T) {
	fmt.Println(t)
}

func TestPrintMethodSet1(t *testing.T) {
	var ty GenericImpl[int]
	var pty *GenericImpl[int]
	reflectool.PrintMethodSet(&ty)
	reflectool.PrintMethodSet(&pty)
	// nil interface
	reflectool.PrintMethodSet((*GenericInter[int])(nil))
}

package reflectool

import (
	"fmt"
	"reflect"
)

func PrintMethodSet(i interface{}) {
	var t = reflect.TypeOf(i)
	var elem = t.Elem()

	var nm = elem.NumMethod()
	if nm == 0 {
		fmt.Printf("%s's method set is empty\n", elem)
		return
	}
	fmt.Printf("%s's method set:\n", elem)
	for i := 0; i < nm; i++ {
		var m = elem.Method(i)
		fmt.Printf("  - %s\n", m.Name)
	}
	fmt.Println()
}

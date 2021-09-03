package test

import (
	"fmt"
	"reflect"
	"testing"
)

type c struct {
	aa int
}

func (cc *c) aaa(b int64) {
	fmt.Println(b)
}

func TestArrary(t *testing.T) {
	// _, err := fmt.Printf("%03d-%s-%04d", 1, "030201", 5)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	var a c = c{
		aa: 1,
	}
	fmt.Println(a)
	value := reflect.ValueOf(a)
	fmt.Println(value.IsZero())
	fmt.Println(value.IsValid())
	// fmt.Println(value.IsNil())
}

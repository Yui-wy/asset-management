package test

import (
	"testing"
)

func TestArrary(t *testing.T) {
	var us []int32 = nil
	println(len(us))
	for _, u := range us {
		println(u)
	}
	println("123")
}

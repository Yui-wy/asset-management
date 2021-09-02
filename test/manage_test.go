package test

import (
	"fmt"
	"testing"
)

func TestArrary(t *testing.T) {
	_, err := fmt.Printf("%03d-%s-%04d", 1, "030201", 5)
	if err != nil {
		fmt.Print(err)
	}

}

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Yui-wy/asset-management/pkg/util/snowflake"
)

func TestArrary(t *testing.T) {
	a, err := snowflake.NewSnowflake(1, 1)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 100; i++ {
		time.Sleep(2)
		id, err := a.NextVal()
		fmt.Println(id)
		if err != nil {
			fmt.Println(err)
		}
	}
}

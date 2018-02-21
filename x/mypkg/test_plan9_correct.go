// +build ignore

package mypkg

import (
	"syscall"
	"fmt"
)

func init() {
	FillSystemInfo = func(sys interface{}) {
		if os, ok := sys.(*syscall.Dir); ok {
			fmt.Println(os)
		}
	}
}

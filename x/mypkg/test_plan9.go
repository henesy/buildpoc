// +build ignore

package mypkg

import (
	"syscall"
	"time"
	"fmt"
)

func init() {
	FillSystemInfo = func(sys interface{}) {
		if os, ok := sys.(*syscall.Dir); ok {
			fmt.Println(os)
		}
	}
}

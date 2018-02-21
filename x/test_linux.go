// +build linux
package mypkg

import (
	"syscall"
	"time"
	"fmt"
)

func init() {
	FillSystemInfo = func(sys interface{}) {
		if os, ok := sys.(*syscall.Stat_t); ok {
			fmt.Println(time.Unix(int64(os.Ctim.Sec), int64(os.Ctim.Nsec)))
			fmt.Println(uint32(os.Dev))
			fmt.Println(uint32(os.Ino))
			fmt.Println(os.Gid)
			fmt.Println(os.Uid)
		}
	}
}

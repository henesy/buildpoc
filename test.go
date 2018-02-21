package main

import (
	"github.com/henesy/buildpoc/x/mypkg"
	"fmt"
	"os"
)

func main() {
	fmt.Println("test follows")
	fi, err := os.Lstat("test.go")
	mypkg.FillSystemInfo(fi.Sys())
	fmt.Println(err)
}


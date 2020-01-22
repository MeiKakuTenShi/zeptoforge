package main

import (
	"runtime"

	"github.com/MeiKakuTenShi/zeptoforge/sandbox"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// clArgs := os.Args
	// fmt.Println(clArgs)

	sandbox.CreateApplication()
	defer sandbox.Close()
	sandbox.Run()
}

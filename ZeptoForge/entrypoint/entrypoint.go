package main

import (
	"runtime"

	"github.com/MeiKakuTenShi/zeptoforge/Sandbox/sandbox"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
)

func init() {
	runtime.LockOSThread()

}

func main() {
	// clArgs := os.Args
	// fmt.Println(clArgs)

	logsys.Init()

	sb := sandbox.CreateApplication()
	defer sb.Close()
	sb.App.Init()
	sb.Run()
}

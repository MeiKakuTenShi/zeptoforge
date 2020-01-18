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

	app := sandbox.CreateApplication()
	defer app.Close()
	app.Init()
	app.Run()
}

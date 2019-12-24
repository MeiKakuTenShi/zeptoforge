package main

import (
	"fmt"
	"os"

	"github.com/MeiKakuTenShi/zeptoforge/Sandbox/sandbox"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
)

func main() {
	clArgs := os.Args
	fmt.Println(clArgs)

	logsys.Init()
	// logsys.ZF_CORE_WARN("Corelog initialized")
	// logsys.ZF_INFO("Clientlog initialized", os.Args)

	sb := sandbox.CreateApplication()
	sb.App.Init()
	sb.Run()
}

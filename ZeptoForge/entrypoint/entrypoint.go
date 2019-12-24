package main

import (
	"github.com/MeiKakuTenShi/zeptoforge/Sandbox/sandbox"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
)

func main() {
	// clArgs := os.Args
	// fmt.Println(clArgs)

	logsys.Init()
	logsys.ZF_CORE_WARN("Corelog initialized")
	logsys.ZF_INFO("Clientlog initialized")

	logsys.PrintLog(logsys.Lcore)
	logsys.PrintLog(logsys.Lclient)

	sb := sandbox.CreateApplication()
	sb.App.Init()
	sb.Run()
}

package sandbox

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/application"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
)

type Sandbox struct {
}

func (sb *Sandbox) Init() {
	logsys.ZF_INFO("Sandbox Application Initialized")
}

func (sb *Sandbox) Run() {
	for x := 0; ; x++ {
		fmt.Println(x % 3)
	}
}

func CreateApplication() *application.Application {
	result := application.NewApplication()
	result.App = &Sandbox{}
	result.Name = "Sandbox"
	return result
}

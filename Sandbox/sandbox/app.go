package sandbox

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"

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

func (sb *Sandbox) OnEvent(e *event.Eventum) {
}

func CreateApplication() *application.Application {
	result := application.NewApplication(&Sandbox{}, "Sandbox")
	return result
}

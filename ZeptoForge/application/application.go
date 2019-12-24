package application

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/appEvent"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
)

type App interface {
	Init()
	Run()
}

type Application struct {
	App  App
	Name string
}

func (Application) Run() {
	winResize := appEvent.NewWindowResizeEvent(1280, 720)

	if winResize.IsInCategory(event.EventCategoryApplication) {
		logsys.ZF_CORE_INFO(winResize.ToString())
	}
	if winResize.IsInCategory(event.EventCategoryInput) {
		logsys.ZF_CORE_INFO(winResize.ToString())
	}

	for x := 0; ; x++ {
		// fmt.Println(x % 2)
	}
}

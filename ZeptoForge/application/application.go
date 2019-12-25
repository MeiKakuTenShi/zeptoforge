package application

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/appEvent"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/window"
)

type App interface {
	Init()
	Run()
	CreateApplication() *Application
}

type Application struct {
	App     App
	window  *Windows.WinWindow
	running bool
	Name    string
}

func NewApplication() *Application {
	win := Windows.WinWindow{}.Create(window.NewWindowProps("", 0, 0))
	return &Application{window: win, running: false}
}

func (app Application) Run() {
	app.running = true
	winResize := appEvent.NewWindowResizeEvent(1280, 720)

	if winResize.IsInCategory(event.EventCategoryApplication) {
		logsys.ZF_CORE_INFO(winResize.ToString())
	}
	if winResize.IsInCategory(event.EventCategoryInput) {
		logsys.ZF_CORE_INFO(winResize.ToString())
	}

	for app.running {
		// fmt.Println(x % 2)
		app.window.OnUpdate()
	}
}

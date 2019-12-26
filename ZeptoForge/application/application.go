package application

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/window"
)

type App interface {
	Init()
	Run()
}

type Application struct {
	App     App
	window  Windows.WinWindow
	running bool
	Name    string
}

func NewApplication() *Application {
	// TODO: make this platform independent
	win := Windows.Create(window.NewWindowProps("", 0, 0))
	return &Application{window: win, running: true}
}

func (app *Application) Close() {
	app.window.Destruct()
}

func (app *Application) Run() {
	logsys.PrintLog(logsys.Lcore)
	logsys.PrintLog(logsys.Lclient)

	for app.running {
		//fmt.Println(x % 2)
		app.window.OnUpdate()
	}
}

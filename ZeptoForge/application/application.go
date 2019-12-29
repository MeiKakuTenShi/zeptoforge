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
	OnEvent(*event.Eventum)
}

type Application struct {
	App     App
	window  Windows.WinWindow
	running bool
	Name    string
}

func NewApplication(app App, name string) *Application {
	// TODO: make this platform independent
	result := &Application{App: app, window: Windows.Create(window.NewWindowProps(name, 0, 0)), running: true, Name: name}
	result.window.SetEventCallback(window.EventCallBackFn{CallbackFn: result.OnEvent})
	return result
}

func (app *Application) Close() {
	app.window.Destruct()
}

func (app *Application) Run() {
	logsys.PrintLog(logsys.Lcore)
	logsys.PrintLog(logsys.Lclient)

	for app.running {
		app.window.OnUpdate()
	}
}

func (app *Application) OnEvent(e *event.Eventum) {
	logsys.ZF_CORE_INFO(e.ToString())

	dispatcher := event.NewEventDispatcher(e)
	dispatcher.Dispatch(event.EventFn{Event: &appEvent.WindowCloseEvent{}, Fn: app.OnWindowClose})
}

func (app *Application) OnWindowClose(w event.Event) bool {
	app.running = false
	return true
}

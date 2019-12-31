package application

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/appEvent"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/layerstack"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/window"
)

type Application struct {
	window  Windows.WinWindow
	stack   layerstack.LayerStack
	running bool
	Name    string
}

func NewApplication(name string) *Application {
	// TODO: make this platform independent
	result := &Application{window: Windows.Create(window.NewWindowProps(name, 0, 0)), stack: layerstack.NewLayerStack(), running: true, Name: name}
	result.window.SetEventCallback(window.EventCallBackFn{CallbackFn: result.OnEvent})
	return result
}

func (app *Application) Close() {
	app.window.Destruct()
}

func (app *Application) PushLayer(layer *layerstack.Layem) {
	app.stack.PushLayer(layer)
}

func (app *Application) PushOverlay(layer *layerstack.Layem) {
	app.stack.PushOverlay(layer)
}

func (app *Application) OnEvent(e *event.Eventum) {
	logsys.ZF_CORE_INFO(e.ToString())

	dispatcher := event.NewEventDispatcher(e)
	dispatcher.Dispatch(event.EventFn{Event: &appEvent.WindowCloseEvent{}, Fn: app.OnWindowClose})

	for _, v := range app.stack.GetStack() {
		v.Layer.OnEvent(e)
		if e.Done() {
			break
		}
	}
}

func (app *Application) Run() {
	logsys.PrintLog(logsys.Lcore)
	logsys.PrintLog(logsys.Lclient)

	for app.running {
		for _, v := range app.stack.GetStack() {
			v.Layer.OnUpdate()
		}
		app.window.OnUpdate()
	}
}

func (app *Application) OnWindowClose(w event.Event) bool {
	app.running = false
	return true
}

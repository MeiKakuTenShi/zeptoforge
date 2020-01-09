package application

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/appEvent"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/layerstack"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/window"
)

var app_SINGLETON *Application

type Application struct {
	window  window.Window
	stack   layerstack.LayerStack
	running bool
	Name    string
}

func NewApplication(name string) *Application {
	// TODO: make this platform independent
	if app_SINGLETON == nil {
		result := new(Application)
		result.window = Windows.NewWinWindow(window.NewWindowProps(name, 0, 0))
		result.stack = layerstack.NewLayerStack()
		result.running = true
		result.Name = name

		result.window.SetEventCallback(window.EventCallBackFn{CallbackFn: result.OnEvent})
		app_SINGLETON = result
	}

	return app_SINGLETON
}

func Get() *Application {
	return app_SINGLETON
}

func DisplaySize() [2]float32 {
	return [2]float32{float32(app_SINGLETON.window.GetWidth()), float32(app_SINGLETON.window.GetHeight())}
}

func FrameBufferSize() [2]float32 {
	return app_SINGLETON.window.FramebufferSize()
}

func GetWindow() window.Window {
	return app_SINGLETON.window
}

func (app *Application) Close() {
	for _, v := range app.stack.GetStack() {
		v.Layer.OnDetach()
	}
	app.window.Destruct()
}

func (app *Application) PushLayer(layer *layerstack.Layem) {
	app.stack.PushLayer(layer)
	layer.Layer.OnAttach()
}

func (app *Application) PushOverlay(layer *layerstack.Layem) {
	app.stack.PushOverlay(layer)
	layer.Layer.OnAttach()
}

func (app *Application) OnEvent(e *event.Eventum) {
	logsys.ZF_CORE_INFO(e.String())

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
	for app.running {
		for _, v := range app.stack.GetStack() {
			v.Layer.OnUpdate()
		}

		app.window.OnUpdate()
	}
}

func (app *Application) OnWindowClose(w event.Eventum) bool {
	app.running = false
	return true
}

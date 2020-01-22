package zforge

import (
	"fmt"
	"runtime"

	"github.com/MeiKakuTenShi/zeptoforge/zforge/Platform/Windows"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/Platform/Windows/winInput"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
)

var (
	appInstance *appSingleton
)

type appSingleton struct {
	window  Window
	gui     *ImGuiLayer
	stack   *LayerStack
	running bool
	Name    string
}

func InitApp(name string) {
	if appInstance == nil {
		initLogSys()
		appInstance = new(appSingleton)
		// Windows specific build
		if runtime.GOOS == "windows" {
			p := &Windows.WinData{Title: name, Width: 0, Height: 0}
			appInstance.window = Windows.NewWinWindow(p)
			appInstance.gui = NewImGuiLayer(appInstance.window.GetWindow())
			initInputSingleton(&winInput.WindowsInput{}) // singleton can only be set once
		} else {
			panic("Platform currently not supported")
		}

		//-----------------------
		appInstance.stack = newLayerStack()
		appInstance.running = true
		appInstance.Name = name

		PushOverlayOnApp(NewLayem(appInstance.gui, "ImGuiLayer"))
		appInstance.window.SetEventCallback(AppOnEvent)

		ZF_CORE_INFO(fmt.Sprintf("Created App: %s (Width %v, Height %v)", appInstance.Name, appInstance.window.GetWidth(), appInstance.window.GetHeight()))
	}
}

func DisplaySize() [2]float32 {
	return [2]float32{float32(appInstance.window.GetWidth()), float32(appInstance.window.GetHeight())}
}

func FrameBufferSize() [2]float32 {
	return appInstance.window.FramebufferSize()
}

func GetApplication() appSingleton {
	return *appInstance
}

func GetWindow() Window {
	return appInstance.window
}

func PushLayerOnApp(l *Layem) {
	appInstance.stack.PushLayer(l)
	l.layer.OnAttach()
}

func PushOverlayOnApp(l *Layem) {
	if appInstance == nil {
		fmt.Println("application::PushOverlayOnApp:app is nil")
	}
	if l == nil {
		fmt.Println("application::PushOverlayOnApp:layem is nil")
	}
	appInstance.stack.PushOverlay(l)
	l.layer.OnAttach()
}

func AppOnEvent(e *event.Eventum) {
	ZF_CORE_INFO(e.String())

	dispatcher := event.NewEventDispatcher(e)
	dispatcher.Dispatch(event.EventFn{Event: &event.WindowCloseEvent{}, Fn: onWindowClose})

	for _, v := range appInstance.stack.layers {
		v.layer.OnEvent(e)
		if e.Done() {
			break
		}
	}
}

func RunApp() {
	for appInstance.running {
		for _, v := range appInstance.stack.layers {
			v.layer.OnUpdate()
		}

		appInstance.gui.Begin()
		for _, v := range appInstance.stack.layers {
			v.layer.OnImGuiRender()
		}
		appInstance.gui.End()
		appInstance.window.OnUpdate()
	}
}
func CloseApp() {
	for _, v := range appInstance.stack.GetStack() {
		v.layer.OnDetach()
	}
	appInstance.window.Destruct()
}

func onWindowClose(e event.Eventum) bool {
	appInstance.running = false
	return true
}

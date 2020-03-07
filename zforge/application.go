package zforge

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows/winInput"
	"github.com/MeiKakuTenShi/zeptoforge/platform/windows"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
)

var (
	appInstance appSingleton
)

type appSingleton struct {
	window  Window
	gui     *ImGuiLayer
	stack   *LayerStack
	running bool
	Name    string
	// timeStep      TimeStep
	lastFrameTime time.Time
}

func appInstanceNil() bool {
	return reflect.DeepEqual(appInstance, appSingleton{})
}

func InitApp(name string) {
	if appInstanceNil() {
		initLogSys()
		// ------- Windows specific build ----------------------------------------------
		if runtime.GOOS == "windows" {
			p := &windows.WinData{Title: name, Width: 0, Height: 0}
			appInstance.window = windows.NewWinWindow(p)
			appInstance.gui = NewImGuiLayer(appInstance.window.GetWindow())
			initInputSingleton(&winInput.WindowsInput{}) // singleton can only be set once
		} else {
			panic("Platform currently not supported")
		}

		//-------------------------------------------------------------------------------
		appInstance.stack = newLayerStack()
		appInstance.running = true
		appInstance.Name = name

		PushOverlayOnApp(NewLayem(appInstance.gui, "ImGuiLayer"))
		appInstance.window.SetEventCallback(AppOnEvent)

		core_INFO(fmt.Sprintf("Created App: %s (Width %v, Height %v)", appInstance.Name, appInstance.window.GetWidth(), appInstance.window.GetHeight()))
	}
}

func DisplaySize() [2]float32 {
	return [2]float32{float32(appInstance.window.GetWidth()), float32(appInstance.window.GetHeight())}
}

func FrameBufferSize() [2]float32 {
	return appInstance.window.FramebufferSize()
}

func GetApplication() (appSingleton, error) {
	if appInstanceNil() {
		return appSingleton{}, errors.New("application has not been created")
	}
	return appInstance, nil
}

func GetWindow() (Window, error) {
	if appInstanceNil() {
		return nil, errors.New("cannot obtain window, application is nil")
	}
	return appInstance.window, nil
}

func PushLayerOnApp(l *Layem) {
	appInstance.stack.PushLayer(l)
	l.layer.OnAttach()
}

func PushOverlayOnApp(l *Layem) {
	appInstance.stack.PushOverlay(l)
	l.layer.OnAttach()
}

func AppOnEvent(e *event.Eventum) {
	core_INFO(e.String())

	dispatcher := event.NewEventDispatcher(e)
	dispatcher.Dispatch(event.EventFn{Event: event.WindowClose, Fn: onWindowClose})

	for _, v := range appInstance.stack.layers {
		v.layer.OnEvent(e)
		if e.Done() {
			break
		}
	}
}

func RunApp() {
	defer closeApp()
	for appInstance.running {
		t := time.Now()
		timeStep := NewTimeStep(time.Since(appInstance.lastFrameTime))
		appInstance.lastFrameTime = t

		for _, v := range appInstance.stack.layers {
			v.layer.OnUpdate(timeStep)
		}

		appInstance.gui.Begin()
		for _, v := range appInstance.stack.layers {
			v.layer.OnImGuiRender()
		}
		appInstance.gui.End()

		appInstance.window.OnUpdate()
	}
}
func closeApp() {
	for _, v := range appInstance.stack.layers {
		v.layer.OnDetach()
	}

	appInstance.window.Destruct()
}

func onWindowClose(rec interface{}, e event.Eventum) bool {
	appInstance.running = false
	return true
}

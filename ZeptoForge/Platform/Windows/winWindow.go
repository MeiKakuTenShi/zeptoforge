package Windows

import (
	"fmt"
	"unsafe"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type winData struct {
	title         string
	width, height int
	vsync         bool
	callback      window.EventCallBackFn
}

type WinWindow struct {
	win    window.Window
	window *glfw.Window
	data   *winData
}

var (
	glfwInitialized = false
	default_title   = "ZeptoForge Application"
	default_width   = 1024
	default_height  = 720
)

func Create(props *window.WindowProps) WinWindow {
	result := WinWindow{data: &winData{}}

	if props.Title == "" {
		props.Title = default_title
	} else {

	}
	if props.Width == 0 {
		props.Width = default_width
	} else {

	}
	if props.Height == 0 {
		props.Height = default_height
	} else {

	}

	result.init(props)

	return result
}

func (win *WinWindow) Destruct() {
	win.Shutdown()
}
func (win *WinWindow) Shutdown() {
	win.window.Destroy()
}
func (win *WinWindow) OnUpdate() {
	glfw.PollEvents()
	win.window.SwapBuffers()
}
func (win *WinWindow) GetWidth() int {
	return win.data.width
}
func (win *WinWindow) GetHeight() int {
	return win.data.height
}
func (win *WinWindow) SetEventCallback(callback *window.EventCallBackFn) {
	win.data.callback = *callback
}
func (win *WinWindow) SetVSync(enabled bool) {
	if enabled {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}
	win.data.vsync = enabled
}
func (win *WinWindow) IsVSync() bool {
	return win.data.vsync
}

func (win *WinWindow) init(props *window.WindowProps) {
	win.data.title = props.Title
	win.data.width = props.Width
	win.data.height = props.Height

	logsys.ZF_CORE_INFO(fmt.Sprintf("Creating window %s (%v, %v)", props.Title, props.Width, props.Height))

	if !glfwInitialized {
		if err := glfw.Init(); err != nil {
			logsys.ZF_CORE_ERROR(err)
		}
		//defer glfw.Terminate()
		glfwInitialized = true
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}

	var err error
	win.window, err = glfw.CreateWindow(props.Width, props.Height, props.Title, nil, nil)
	if err != nil {
		logsys.ZF_CORE_ERROR(err)
	}
	win.window.MakeContextCurrent()
	win.window.SetUserPointer(unsafe.Pointer(win.data))
	win.SetVSync(true)
}

package windows

import (
	"fmt"
	"unsafe"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/renderer"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	default_title  = "ZeptoForge Application"
	default_width  = 1024
	default_height = 720
)

var (
	glfwInitialized = false
)

type WinData struct {
	Title         string
	Width, Height int
	Vsync         bool
	Callback      func(*event.Eventum)
}

type WinWindow struct {
	window  *glfw.Window
	data    *WinData
	context renderer.GraphicsContext
}

func NewWinWindow(props *WinData) *WinWindow {
	win := new(WinWindow)
	win.data = new(WinData)
	win.init(props)
	return win
}

//----------------- Window interface -------------------
func (win *WinWindow) Destruct() {
	win.window.Destroy()
	glfw.Terminate()
}
func (win *WinWindow) OnUpdate() {
	glfw.PollEvents()
	win.context.SwapBuffers()
}
func (win *WinWindow) GetWindow() *glfw.Window {
	return win.window
}
func (win *WinWindow) IsVSync() bool {
	return win.data.Vsync
}
func (win *WinWindow) GetHeight() int {
	_, h := win.window.GetSize()
	return h
}
func (win *WinWindow) GetWidth() int {
	w, _ := win.window.GetSize()
	return w
}
func (win *WinWindow) FramebufferSize() [2]float32 {
	w, h := win.window.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}
func (win *WinWindow) SetEventCallback(callback func(*event.Eventum)) {
	win.data.Callback = callback
}
func (win *WinWindow) SetVSync(enabled bool) {
	if enabled {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}
	win.data.Vsync = enabled
}

func (win *WinWindow) init(props *WinData) {
	if props.Title == "" {
		win.data.Title = default_title
	} else {
		win.data.Title = props.Title
	}
	if props.Width == 0 {
		win.data.Width = default_width
	} else {
		win.data.Width = props.Width
	}
	if props.Height == 0 {
		win.data.Height = default_height
	} else {
		win.data.Height = props.Height
	}

	var err error
	// Initialize GLFW
	if !glfwInitialized {
		if err = glfw.Init(); err != nil {
			panic(err)
		}
		glfwInitialized = true
		glfw.WindowHint(glfw.Resizable, glfw.True)
		glfw.WindowHint(glfw.ContextVersionMajor, 4)
		glfw.WindowHint(glfw.ContextVersionMinor, 5)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.False)
	}
	// Create and Initialize Window
	win.window, err = glfw.CreateWindow(win.data.Width, win.data.Height, win.data.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	win.context = &opengl.OpenGLContext{}
	fmt.Println(win.context.Init(win.window))

	win.window.SetUserPointer(unsafe.Pointer(win.data))
	win.SetVSync(true)

	// Set GLFW callbacks
	// Window
	win.window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		data := *(*WinData)(w.GetUserPointer())
		data.Width = width
		data.Height = height

		data.Callback(event.NewWindowResizeEvent(width, height))
	})

	win.window.SetCloseCallback(func(w *glfw.Window) {
		data := *(*WinData)(w.GetUserPointer())

		data.Callback(event.NewWindowCloseEvent())
	})

	// Keys
	win.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		data := *(*WinData)(w.GetUserPointer())

		switch action {
		case glfw.Press:
			{
				data.Callback(event.NewKeyPressedEvent(int(key), 0))
				break
			}
		case glfw.Release:
			{
				data.Callback(event.NewKeyReleasedEvent(int(key)))
				break
			}
		case glfw.Repeat:
			{
				data.Callback(event.NewKeyPressedEvent(int(key), 1))
				break
			}
		}
	})

	win.window.SetCharCallback(func(w *glfw.Window, char rune) {
		data := *(*WinData)(w.GetUserPointer())
		data.Callback(event.NewKeyTypedEvent(char))
	})

	// Mouse
	win.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		data := *(*WinData)(w.GetUserPointer())

		switch action {
		case glfw.Press:
			{
				data.Callback(event.NewMouseButtonPressedEvent(int(button)))
				break
			}
		case glfw.Release:
			{
				data.Callback(event.NewMouseButtonReleasedEvent(int(button)))
				break
			}
		}
	})

	win.window.SetScrollCallback(func(w *glfw.Window, xOff, yOff float64) {
		data := *(*WinData)(w.GetUserPointer())
		data.Callback(event.NewMouseScrolledEvent(xOff, yOff))
	})

	win.window.SetCursorPosCallback(func(w *glfw.Window, xPos, yPos float64) {
		data := *(*WinData)(w.GetUserPointer())
		data.Callback(event.NewMouseMovedEvent(xPos, yPos))
	})
}

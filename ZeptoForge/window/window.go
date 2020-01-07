package window

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowProps struct {
	Title  string
	Width  int
	Height int
}

func NewWindowProps(title string, width, height int) *WindowProps {
	return &WindowProps{Title: title, Width: width, Height: height}
}

type EventCallBackFn struct {
	CallbackFn func(*event.Eventum)
}

type Window interface {
	Destruct()
	OnUpdate()
	// Getters
	GetWindow() *glfw.Window
	GetWidth() int
	GetHeight() int
	FramebufferSize() [2]float32
	// Window attributes
	SetEventCallback(EventCallBackFn)
	SetVSync(bool)
	IsVSync() bool

	Create(*WindowProps) Window
}

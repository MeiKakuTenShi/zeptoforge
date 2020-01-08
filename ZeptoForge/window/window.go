package window

import (
	"unsafe"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
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
	GetNativeWindow() unsafe.Pointer
	GetWidth() int
	GetHeight() int
	FramebufferSize() [2]float32
	// Window attributes
	SetEventCallback(EventCallBackFn)
	SetVSync(bool)
	IsVSync() bool

	Create(*WindowProps) Window
}

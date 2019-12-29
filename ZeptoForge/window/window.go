package window

import (
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
	GetWidth()
	GetHeight()
	// Window attributes
	SetEventCallback(EventCallBackFn)
	SetVSync(bool)
	IsVSync()

	Create(*WindowProps) *Window
}

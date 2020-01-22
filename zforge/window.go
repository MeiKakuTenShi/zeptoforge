package zforge

import (
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window interface {
	Destruct()
	OnUpdate()
	// Getters
	GetWindow() *glfw.Window
	IsVSync() bool
	GetWidth() int
	GetHeight() int
	FramebufferSize() [2]float32
	// Setters
	SetEventCallback(func(*event.Eventum))
	SetVSync(bool)
}

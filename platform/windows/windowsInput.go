package windows

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type WindowsInput struct {
}

func (WindowsInput) IsKeyPressedImpl(keycode int) bool {
	win := glfw.GetCurrentContext()
	// win := *(*glfw.Window)(application.GetWindow().GetNativeWindow())
	state := win.GetKey(glfw.Key(keycode))

	return state == glfw.Press || state == glfw.Repeat
}
func (WindowsInput) IsMouseButtonPressedImpl(button int) bool {
	win := glfw.GetCurrentContext()
	// win := *(*glfw.Window)(application.GetWindow().GetNativeWindow())
	state := win.GetMouseButton(glfw.MouseButton(button))

	return state == glfw.Press
}
func (WindowsInput) GetMousePositionImpl() (x, y float32) {
	win := glfw.GetCurrentContext()
	// win := *(*glfw.Window)(application.GetWindow().GetNativeWindow())
	xPos, yPos := win.GetCursorPos()

	return float32(xPos), float32(yPos)
}
func (WindowsInput) GetMouseXImpl() float32 {
	x, _ := WindowsInput{}.GetMousePositionImpl()

	return x
}
func (WindowsInput) GetMouseYImpl() float32 {
	_, y := WindowsInput{}.GetMousePositionImpl()

	return y
}

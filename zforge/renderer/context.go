package renderer

import "github.com/go-gl/glfw/v3.3/glfw"

type GraphicsContext interface {
	Init(*glfw.Window) string
	SwapBuffers()
}

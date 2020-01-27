package opengl

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type OpenGLContext struct {
	windowHandle *glfw.Window
}

func (con *OpenGLContext) Init(handle *glfw.Window) string {
	con.windowHandle = handle
	con.windowHandle.MakeContextCurrent()
	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	return fmt.Sprintf("OpenGL|\tVendor (%s)\n\tRenderer: (%s)\n\tVersion: (%s)",
		gl.GoStr(gl.GetString(gl.VENDOR)),
		gl.GoStr(gl.GetString(gl.RENDERER)),
		gl.GoStr(gl.GetString(gl.VERSION)),
	)
}

func (con *OpenGLContext) SwapBuffers() {
	con.windowHandle.SwapBuffers()
}

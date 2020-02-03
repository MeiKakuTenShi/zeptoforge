package opengl

import (
	"fmt"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type OpenGLVertexArray struct {
	rendererID uint32
}

func (va *OpenGLVertexArray) Init() {
	gl.CreateVertexArrays(1, &va.rendererID)
}

func (va *OpenGLVertexArray) Bind() {
	gl.BindVertexArray(va.rendererID)
}
func (va *OpenGLVertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (va *OpenGLVertexArray) String() string {
	return fmt.Sprintf("OpenGL Vertex Array {ID: %v}", va.rendererID)
}

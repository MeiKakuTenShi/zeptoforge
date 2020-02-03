package opengl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

type OpenGLVertexBuffer struct {
	rendererID uint32
}

func (vb *OpenGLVertexBuffer) Init(vertices []float32, size int) {
	// fmt.Println("Initializing OpenGL Vertex Buffer ", vertices, " ", size)

	gl.GenBuffers(1, &vb.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(vertices), gl.STATIC_DRAW)

}
func (vb *OpenGLVertexBuffer) Destruct() {
	gl.DeleteBuffers(1, &vb.rendererID)
}
func (vb *OpenGLVertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.rendererID)
}
func (vb *OpenGLVertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

type OpenGLIndexBuffer struct {
	rendererID uint32
	count      int
}

func (ib *OpenGLIndexBuffer) Init(indices []uint32, count int) {
	// fmt.Printf("Initializing OpenGL Index Buffer: indices(%v) count(%v)\n", indices, count)

	ib.count = count
	gl.GenBuffers(1, &ib.rendererID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.rendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, count*4, gl.Ptr(indices), gl.STATIC_DRAW)
}
func (ib *OpenGLIndexBuffer) GetCount() int32 {
	return int32(ib.count)
}
func (ib *OpenGLIndexBuffer) Destruct() {
	gl.DeleteBuffers(1, &ib.rendererID)
}
func (ib *OpenGLIndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.rendererID)
}
func (ib *OpenGLIndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

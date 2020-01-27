package opengl

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type OpenGLVertexBuffer struct {
	rendererID uint32
}

func (vb *OpenGLVertexBuffer) Init(vertices []float32, size int) {
	gl.CreateBuffers(1, &vb.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, size, gl.Ptr(vertices), gl.STATIC_DRAW)

}

func (vb OpenGLVertexBuffer) Destuct() {
	gl.DeleteBuffers(1, &vb.rendererID)
}

func (vb OpenGLVertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vb.rendererID)
}

func (vb OpenGLVertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

type OpenGLIndexBuffer struct {
	rendererID uint32
	count      int
}

func (ib *OpenGLIndexBuffer) Init(indices []float32, c int) {
	ib.count = c
	gl.CreateBuffers(1, &ib.rendererID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.rendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, c*4, gl.Ptr(indices), gl.STATIC_DRAW)

}

func (ib OpenGLIndexBuffer) GetCount() int {
	return ib.count
}

func (ib OpenGLIndexBuffer) Destuct() {
	gl.DeleteBuffers(1, &ib.rendererID)
}

func (ib OpenGLIndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ib.rendererID)
}

func (ib OpenGLIndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

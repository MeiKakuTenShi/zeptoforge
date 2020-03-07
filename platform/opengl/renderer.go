package opengl

import (
	"github.com/go-gl/gl/v4.5-core/gl"
)

type Renderer struct {
}

func (Renderer) SetClearColor(color []float32) {
	gl.ClearColor(color[0], color[1], color[2], color[3])
}

func (Renderer) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (Renderer) DrawIndexed(count int32) {
	gl.DrawElements(gl.TRIANGLES, count, gl.UNSIGNED_INT, nil)
}

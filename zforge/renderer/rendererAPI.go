package renderer

import (
	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
)

type API uint8

const (
	NoneRenderer API = iota
	OpenGL
)

// Will be set dynamically in the future depending on the context
var (
	sAPI = ZFrenderer{api: OpenGL, renderer: opengl.Renderer{}}
)

type Renderer interface {
	SetClearColor([]float32)
	Clear()
	DrawIndexed(int32)
}

type ZFrenderer struct {
	api      API
	renderer Renderer
}

func GetAPI() API {
	return sAPI.api
}

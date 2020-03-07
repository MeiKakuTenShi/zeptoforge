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
	static_API = RendererAPI{api: OpenGL, renderer: opengl.Renderer{}}
)

type RendAPI interface {
	SetClearColor([]float32)
	Clear()
	DrawIndexed(int32)
}

type RendererAPI struct {
	api      API
	renderer RendAPI
}

func GetAPI() API {
	return static_API.api
}

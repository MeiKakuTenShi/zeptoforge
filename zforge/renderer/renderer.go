package renderer

type RendererAPI uint32

const (
	NoneRenderer RendererAPI = iota
	OpenGL
)

func GetAPI() RendererAPI {
	return sRendererAPI
}

var sRendererAPI = OpenGL

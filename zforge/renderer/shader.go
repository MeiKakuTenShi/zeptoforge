package renderer

import (
	"errors"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
)

type Shader interface {
	Bind()
	Unbind()
	Dispose()
}

type ZFshader struct {
	shader Shader
}

func NewShader(vSrc, fSrc string) (*ZFshader, error) {
	switch sAPI.api {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLShader)
		r.Init(vSrc, fSrc)
		return &ZFshader{shader: r}, nil
	default:
		return nil, errors.New("could not create index buffer; unkown api")
	}
}

func (s ZFshader) Bind() {
	s.shader.Bind()
}

func (s ZFshader) Unbind() {
	s.shader.Unbind()
}

func (s ZFshader) GetShader() (Shader, error) {
	if s.shader != nil {
		return s.shader, nil
	}

	return nil, errors.New("shader is empty")
}

package renderer

import (
	"errors"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
)

type Texture interface {
	Init(path string) error
	GetWidth() int32
	GetHeight() int32
	Bind(uint32)
	Destruct()
}

type Texture2D struct {
	texture Texture
}

func NewTexture2D(filepath string) (*Texture2D, error) {
	switch sAPI.api {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLTexture2D)
		err := r.Init(filepath)
		if err != nil {
			return &Texture2D{}, err
		}
		return &Texture2D{texture: r}, nil
	default:
		return &Texture2D{}, errors.New("renderer api unknown")
	}
}

func (tex *Texture2D) Bind() {
	tex.texture.Bind(0)
}

func (tex *Texture2D) BindInSlot(slot uint32) {
	tex.texture.Bind(slot)
}

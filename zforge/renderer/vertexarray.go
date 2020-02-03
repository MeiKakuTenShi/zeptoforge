package renderer

import (
	"errors"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type VertexArray interface {
	Bind()
	Unbind()

	Init()
	String() string
}

type VertArray struct {
	va  VertexArray
	vbs []VertBuff
	ib  IndBuff
}

func NewVertexArray() (*VertArray, error) {
	switch sRendererAPI {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLVertexArray)
		r.Init()
		return &VertArray{va: r}, nil
	default:
		return &VertArray{}, errors.New("renderer api unknown")
	}
}
func (va *VertArray) Bind() {
	va.va.Bind()
}
func (va *VertArray) GetVertexBuffers() []VertBuff {
	return va.vbs
}
func (va *VertArray) GetIndexBuffer() IndBuff {
	return va.ib
}
func (va *VertArray) AddVertexBuffer(vb VertBuff) error {
	if len(vb.layout.elements) == 0 {
		return errors.New("vertex buffer has no layout")
	}

	switch sRendererAPI {
	case NoneRenderer:
		panic("AddVertexBuffer::RendererAPI::None - currently not supported")
	case OpenGL:
		va.va.Bind()
		vb.vb.Bind()

		index := uint32(0)
		for _, v := range vb.layout.elements {
			gl.EnableVertexAttribArray(index)
			gl.VertexAttribPointer(index,
				v.GetComponentCount(),
				ShaderDataTypeToOpenGLBaseType(v.DataType),
				v.Normalized,
				vb.layout.stride,
				gl.PtrOffset(int(v.Offset)))
			index++
		}
		va.vbs = append(va.vbs, vb)
		return nil
	default:
		return errors.New("unknown api")
	}
}
func (va *VertArray) SetIndexBuffer(ib IndBuff) {
	switch sRendererAPI {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		va.va.Bind()
		ib.ib.Bind()
		va.ib = ib
	}
}

func ShaderDataTypeToOpenGLBaseType(t ShaderDataType) uint32 {
	switch t {
	case Float:
		return gl.FLOAT
	case Float2:
		return gl.FLOAT
	case Float3:
		return gl.FLOAT
	case Float4:
		return gl.FLOAT
	case Mat3:
		return gl.FLOAT
	case Mat4:
		return gl.FLOAT
	case Int:
		return gl.INT
	case Int2:
		return gl.INT
	case Int3:
		return gl.INT
	case Int4:
		return gl.INT
	case Bool:
		return gl.BOOL
	default:
		return 0
	}
}

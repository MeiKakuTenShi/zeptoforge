package renderer

import (
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
)

type ShaderDataType uint8

const (
	NoneShaderType ShaderDataType = iota
	Float
	Float2
	Float3
	Float4
	Mat3
	Mat4
	Int
	Int2
	Int3
	Int4
	Bool
)

func dataTypeSize(t ShaderDataType) int32 {
	switch t {
	case Float:
		return 4
	case Float2:
		return 4 * 2
	case Float3:
		return 4 * 3
	case Float4:
		return 4 * 4
	case Mat3:
		return 4 * 3 * 3
	case Mat4:
		return 4 * 4 * 4
	case Int:
		return 4
	case Int2:
		return 4 * 2
	case Int3:
		return 4 * 3
	case Int4:
		return 4 * 4
	case Bool:
		return 1
	default:
		return 0
	}
}

type Buffem struct {
	name     string
	offset   int32
	size     int32
	dataType ShaderDataType
}

func NewBuffem(dt ShaderDataType, n string) *Buffem {
	r := new(Buffem)
	r.name = n
	r.dataType = dt
	r.size = dataTypeSize(dt)
	return r
}

func (bm Buffem) String() string {
	return fmt.Sprintf("{name: %s, type: %v, size: %v, offset: %v}",
		bm.name,
		bm.dataType,
		bm.size,
		bm.offset,
	)
}

var 

type BufferLayout struct {
	elements []*Buffem
	stride   int32
}

func NewBufferLayout(bm ...*Buffem) BufferLayout {
	r := new(BufferLayout)

	for _, v := range bm {
		r.elements = append(r.elements, v)
	}

	r.calcOffSetStride()

	return *r
}

func (bl *BufferLayout) calcOffSetStride() {
	off := int32(0)
	bl.stride = 0

	for _, v := range bl.elements {
		v.offset = off
		off += v.size
		bl.stride += v.size
	}
}

func (bl BufferLayout) GetElements() []*Buffem {
	return bl.elements
}

func (bl BufferLayout) String() string {
	r := fmt.Sprintf("BufferLayout: stride(%v), size(%v)\n\telements:\n\t",
		bl.stride,
		len(bl.elements),
	)

	for _, v := range bl.elements {
		r += fmt.Sprint(v, "\n\t")
	}

	return r
}

type VertexBuffer interface {
	Bind()
	Unbind()

	Init([]float32, int)
}

func CreateVertexBuffer(vertices []float32, size int) VertexBuffer {
	switch sRendererAPI {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLVertexBuffer)
		r.Init(vertices, size)
		sVertexBuffer = r
		return r
	default:
		return nil
	}
}

func (vb VertexBuffer) SetLayout(BufferLayout) {

}
func (vb GetLayout() BufferLayout

type IndexBuffer interface {
	Bind()
	Unbind()
	GetCount() int
	Init([]float32, int)
}

func CreateIndexBuffer(indices []float32, count int) IndexBuffer {
	switch sRendererAPI {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLIndexBuffer)
		r.Init(indices, count)
		return r
	default:
		return nil
	}
}

package renderer

import (
	"errors"
	"fmt"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
)

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

type BuffElem struct {
	Name       string
	Offset     int
	Size       int
	Normalized bool
	DataType   ShaderDataType
}

func NewBuffElem(dt ShaderDataType, n string, norm bool) *BuffElem {
	r := new(BuffElem)
	r.Name = n
	r.DataType = dt
	r.Size = dataTypeSize(dt)
	r.Normalized = norm
	return r
}
func (bm BuffElem) String() string {
	return fmt.Sprintf("{name: %s, type: %s, size: %v, normalized: %v, offset: %v}",
		bm.Name,
		bm.DataType,
		bm.Size,
		bm.Normalized,
		bm.Offset,
	)
}
func (bm *BuffElem) GetComponentCount() int32 {
	switch bm.DataType {
	case Float:
		return 1
	case Float2:
		return 2
	case Float3:
		return 3
	case Float4:
		return 4
	case Mat3:
		return 3 * 3
	case Mat4:
		return 4 * 4
	case Int:
		return 1
	case Int2:
		return 2
	case Int3:
		return 3
	case Int4:
		return 4
	case Bool:
		return 1
	default:
		return 0
	}
}

type BufferLayout struct {
	elements []*BuffElem
	stride   int32
}

func NewBufferLayout(bm ...*BuffElem) BufferLayout {
	r := new(BufferLayout)

	for _, v := range bm {
		r.elements = append(r.elements, v)
	}
	r.calcOffSetStride()

	return *r
}
func (bl *BufferLayout) calcOffSetStride() {
	off := 0
	bl.stride = 0

	for _, v := range bl.elements {
		v.Offset = off
		off += v.Size
		bl.stride += int32(v.Size)
	}
}
func (bl BufferLayout) GetElements() []*BuffElem {
	return bl.elements
}
func (bl BufferLayout) GetStride() int32 {
	return bl.stride
}
func (bl BufferLayout) String() string {
	r := fmt.Sprintf("BufferLayout: stride(%v), size(%v)\n\telements:\n\t", bl.stride, len(bl.elements))

	for _, v := range bl.elements {
		r += fmt.Sprint("\t", v, "\n\t")
	}

	return r
}

type VertexBuffer interface {
	Bind()
	Unbind()

	Init([]float32, int)
	Destruct()
}

type VertBuff struct {
	vb     VertexBuffer
	layout BufferLayout
}

func NewVertexBuffer(vertices []float32, size int) (*VertBuff, error) {
	switch sRendererAPI {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLVertexBuffer)
		r.Init(vertices, size)
		return &VertBuff{vb: r}, nil
	default:
		return &VertBuff{}, errors.New("could not create vertex buffer; unknown api")
	}
}
func (vb *VertBuff) SetLayout(layout BufferLayout) {
	vb.layout = layout
}
func (vb *VertBuff) GetLayout() BufferLayout {
	return vb.layout
}
func (vb *VertBuff) Remove() {
	vb.vb.Destruct()
}

type IndexBuffer interface {
	Bind()
	Unbind()

	GetCount() int32

	Init([]uint32, int)
	Destruct()
}

type IndBuff struct {
	ib IndexBuffer
}

func NewIndexBuffer(indices []uint32, count int) (*IndBuff, error) {
	switch sRendererAPI {
	case NoneRenderer:
		panic("RendererAPI::None - currently not supported")
	case OpenGL:
		r := new(opengl.OpenGLIndexBuffer)
		r.Init(indices, count)
		return &IndBuff{ib: r}, nil
	default:
		return &IndBuff{}, errors.New("could not create index buffer; unkown api")
	}
}
func (ib IndBuff) GetCount() int32 {
	return ib.ib.GetCount()
}
func (ib *IndBuff) Remove() {
	ib.ib.Destruct()
}

type ShaderDataType uint8

func dataTypeSize(t ShaderDataType) int {
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
func (dt ShaderDataType) String() string {
	switch dt {
	case Float:
		return "Float"
	case Float2:
		return "Float2"
	case Float3:
		return "Float3"
	case Float4:
		return "Float4"
	case Mat3:
		return "Mat3"
	case Mat4:
		return "Mat4"
	case Int:
		return "Int"
	case Int2:
		return "Int2"
	case Int3:
		return "Int3"
	case Int4:
		return "Int4"
	case Bool:
		return "Bool"
	default:
		return "Unknown"
	}
}

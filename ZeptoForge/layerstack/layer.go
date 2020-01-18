package layerstack

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

type Layer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnImGuiRender()
	OnEvent(*event.Eventum)
}

type Layem struct {
	Layer     Layer
	debugName string
}

func NewLayem(lay Layer, name string) *Layem {
	l := new(Layem)
	l.Layer = lay
	l.debugName = name

	return l
}

func (l Layem) GetName() string {
	return l.debugName
}

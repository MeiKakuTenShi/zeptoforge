package layerstack

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
)

type Layer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnEvent(*event.Eventum)
}

type Layem struct {
	Layer     Layer
	debugName string
}

func NewLayem(lay Layer, name string) *Layem {
	return &Layem{Layer: lay, debugName: name}
}

func (l Layem) GetName() string {
	return l.debugName
}

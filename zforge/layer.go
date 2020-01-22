package zforge

import (
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
)

type Layer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnImGuiRender()
	OnEvent(*event.Eventum)
}

type Layem struct {
	layer     Layer
	debugName string
}

func NewLayem(lay Layer, name string) *Layem {
	l := new(Layem)
	l.layer = lay
	l.debugName = name

	return l
}

func (l Layem) GetName() string {
	return l.debugName
}

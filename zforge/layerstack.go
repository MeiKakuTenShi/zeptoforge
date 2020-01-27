package zforge

import "github.com/MeiKakuTenShi/zeptoforge/zforge/event"

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

type LayerStack struct {
	layers      []*Layem
	layerInsert int
}

func newLayerStack() *LayerStack {
	r := new(LayerStack)

	r.layers = []*Layem{}
	r.layerInsert = 0

	return r
}
func (ls *LayerStack) PushLayer(layer *Layem) {
	ls.layers = append(ls.layers, layer)
	copy(ls.layers[ls.layerInsert+1:], ls.layers[ls.layerInsert:])
	ls.layers[ls.layerInsert] = layer
}
func (ls *LayerStack) PushOverlay(overlay *Layem) {
	ls.layers = append(ls.layers, overlay)
}
func (ls LayerStack) PopLayer(layer *Layem) {
	for k, v := range ls.layers {
		if v == layer {
			copy(ls.layers[k:], ls.layers[k+1:])
			ls.layers = ls.layers[:len(ls.layers)-1]
			ls.layerInsert--
		}
	}
}
func (ls LayerStack) PopOverlay(overlay *Layem) {
	for k, v := range ls.layers {
		if v == overlay {
			copy(ls.layers[k:], ls.layers[k+1:])
			ls.layers = ls.layers[:len(ls.layers)-1]
		}
	}
}

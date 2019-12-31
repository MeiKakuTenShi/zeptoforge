package layerstack

type LayerStack struct {
	layers      []*Layem
	layerInsert int
}

func NewLayerStack() LayerStack {
	return LayerStack{layers: []*Layem{}, layerInsert: 0}
}

func (ls *LayerStack) PushLayer(layer *Layem) {
	// fmt.Printf("Pushing %s onto stack\n", layer.debugName)
	// if ls.layers == nil {
	// 	fmt.Println("layerstack is nil")
	// } else if len(ls.layers) == 0 {
	// 	fmt.Println("layerstack empty")
	// }

	ls.layers = append(ls.layers, layer)
	copy(ls.layers[ls.layerInsert+1:], ls.layers[ls.layerInsert:])
	ls.layers[ls.layerInsert] = layer

	// for _, v := range ls.layers {
	// 	fmt.Printf("%s is on the stack\n", v.debugName)
	// }

	// if len(ls.layers) > 0 {
	// 	fmt.Println("layerstack not empty")
	// }
}

func (ls LayerStack) PushOverlay(overlay *Layem) {
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

func (ls LayerStack) GetStack() []*Layem {
	return ls.layers
}

func (ls LayerStack) Length() int {
	return len(ls.layers)
}

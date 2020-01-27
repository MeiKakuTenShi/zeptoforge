package sandbox

import (
	"github.com/MeiKakuTenShi/zeptoforge/zforge"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
)

func CreateApplication() {
	zforge.InitApp("Sandbox")
	initSB()
}

func initSB() {
	zforge.PushLayerOnApp(zforge.NewLayem(&ExLayer{}, "ExampleLayer"))
	zforge.ZF_INFO("Sandbox Application Initialized")
}

func Run() {
	zforge.RunApp()
}

func Close() {
}

type ExLayer struct {
}

func (l ExLayer) OnAttach() {
}

func (l ExLayer) OnDetach() {
}

func (l ExLayer) OnImGuiRender() {
}

func (l ExLayer) OnUpdate() {

	if zforge.IsKeyPressed(int(zforge.ZF_KeyTab)) {
		zforge.ZF_INFO("Tab key is pressed!")
	}
}

func (l ExLayer) OnEvent(e *event.Eventum) {
	// keyDis := event.NewEventDispatcher(e)
	// keyDis.Dispatch(event.EventFn{Event: event.MouseMoved, Fn: WhoopsEx})
}

func WhoopsEx(e event.Eventum) bool {
	if mme, ok := e.GetEvent().(*event.MouseMovedEvent); ok {
		xPos := mme.GetX()
		yPos := mme.GetY()

		switch xPos < 512 {
		case true:
			if yPos < 360 {
				zforge.ZF_INFO("Mouse in quadrant II")
			} else {
				zforge.ZF_INFO("Mouse in quadrant III")
			}
		case false:
			if yPos < 360 {
				zforge.ZF_INFO("Mouse in quadrant I")
			} else {
				zforge.ZF_INFO("Mouse in quadrant IV")
			}
		}
	}

	return true
}

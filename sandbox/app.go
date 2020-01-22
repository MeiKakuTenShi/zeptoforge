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
	zforge.ZF_INFO("Sandbox Application Initialized")
	zforge.PushLayerOnApp(zforge.NewLayem(&ExLayer{}, "ExampleLayer"))
}

func Run() {
	zforge.RunApp()
}

func Close() {
	zforge.CloseApp()
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
	// logsys.ZF_INFO("ExLayer::Update")

	if zforge.IsKeyPressed(int(zforge.ZF_KeyTab)) {
		zforge.ZF_INFO("Tab key is pressed!")
	}
}

func (l ExLayer) OnEvent(e *event.Eventum) {
	// logsys.ZF_TRACE(e.String())
}

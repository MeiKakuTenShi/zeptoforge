package sandbox

import (
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/application"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/imgui"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/input"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/layerstack"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/logsys"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/keycodes"

	// This is imported to initialize Windows specific event polling, this may vary pending platform
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows/winInput"
)

type Sandbox struct {
	app *application.Application
}

func (sb *Sandbox) Init() {
	logsys.ZF_INFO("Sandbox Application Initialized")
	sb.app.PushLayer(layerstack.NewLayem(&ExLayer{}, "ExampleLayer"))
	sb.app.PushOverlay(layerstack.NewLayem(&imgui.ImGuiLayer{}, "ImGuiLayer"))
	// Possibly temporary
	winInput.SetWindowsInput()
}

func (sb *Sandbox) Run() {
	sb.app.Run()
}

func (sb *Sandbox) Close() {
	sb.app.Close()
}

func CreateApplication() *Sandbox {
	result := application.NewApplication("Sandbox")
	return &Sandbox{app: result}
}

type ExLayer struct {
}

func (l ExLayer) OnAttach() {
}

func (l ExLayer) OnDetach() {
}

func (l ExLayer) OnUpdate() {
	// logsys.ZF_INFO("ExLayer::Update")

	if input.IsKeyPressed(int(keycodes.ZF_KeyTab)) {
		logsys.ZF_INFO("Tab key is pressed!")
	}
}

func (l ExLayer) OnEvent(e *event.Eventum) {
	// logsys.ZF_TRACE(e.String())
}

package imguiLayer

import (
	IG "github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/imgui"
	"github.com/inkyblackness/imgui-go"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/window"
)

var (
	showDemoWindow = true
)

type ImGuiLayer struct {
	imgui *IG.Imgui
}

func NewImGuiLayer(win window.Window) *ImGuiLayer {
	result := new(ImGuiLayer)
	result.imgui = IG.NewImgui(win)
	return result
}

func (gui *ImGuiLayer) OnAttach() {

}

func (gui *ImGuiLayer) OnDetach() {
	gui.imgui.Destruct()
}

func (gui *ImGuiLayer) Begin() {
	gui.imgui.Begin()
}
func (gui *ImGuiLayer) End() {
	gui.imgui.End()
}

func (gui *ImGuiLayer) OnUpdate() {

}

func (gui *ImGuiLayer) OnImGuiRender() {
	if showDemoWindow {
		imgui.ShowDemoWindow(&showDemoWindow)
	}
}

func (gui *ImGuiLayer) OnEvent(e *event.Eventum) {

}

// ClipboardText returns the current clipboard text, if available.
func (gui *ImGuiLayer) ClipboardText() string {
	return gui.imgui.ClipboardText()
}

// SetClipboardText sets the text as the current clipboard text.
func (gui *ImGuiLayer) SetClipboardText(text string) {
	gui.imgui.SetClipboardText(text)
}

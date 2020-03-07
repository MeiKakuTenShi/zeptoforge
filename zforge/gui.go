package zforge

import (
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	IG "github.com/MeiKakuTenShi/zeptoforge/zforge/imgui"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	showDemoWindow = true
)

type ImGuiLayer struct {
	imgui *IG.Imgui
}

func NewImGuiLayer(win *glfw.Window) *ImGuiLayer {
	result := new(ImGuiLayer)
	result.imgui = IG.NewImgui(win)
	return result
}

//------------------ Wrappers --------------------------
func (gui *ImGuiLayer) Begin() {
	gui.imgui.Begin()
}
func (gui *ImGuiLayer) End() {
	gui.imgui.End()
}

//------------------------------------------------------

//------------------ Layer interface -------------------
func (gui *ImGuiLayer) OnAttach() {
}
func (gui *ImGuiLayer) OnDetach() {
	gui.imgui.Destruct()
}
func (gui *ImGuiLayer) OnUpdate(ts TimeStep) {
}
func (gui *ImGuiLayer) OnImGuiRender() {
	if showDemoWindow {
		gui.imgui.ShowDemo(&showDemoWindow)
	}
}
func (gui *ImGuiLayer) OnEvent(e *event.Eventum) {
}

//-----------------------------------------------------

//------------------- Extra ---------------------------
// ClipboardText returns the current clipboard text, if available.
func (gui *ImGuiLayer) ClipboardText() string {
	return gui.imgui.ClipboardText()
}

// SetClipboardText sets the text as the current clipboard text.
func (gui *ImGuiLayer) SetClipboardText(text string) {
	gui.imgui.SetClipboardText(text)
}

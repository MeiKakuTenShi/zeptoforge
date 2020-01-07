package imgui

import (
	imgui "github.com/inkyblackness/imgui-go"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/application"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/appEvent"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/keyEvent"
	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/event/mouseEvent"

	// TEMPORARY
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	showDemoWindow      = true
	clearColor          = [4]float32{0.0, 0.0, 0.0, 1.0}
	glfwButtonIDByIndex = map[int]glfw.MouseButton{
		0: glfw.MouseButton1,
		1: glfw.MouseButton2,
		2: glfw.MouseButton3,
	}
)

type ImGuiLayer struct {
	io               imgui.IO
	context          *imgui.Context
	renderer         *OpenGL3
	time             float64
	mouseJustPressed [3]bool
}

func (gui *ImGuiLayer) OnAttach() {
	gui.context = imgui.CreateContext(nil)

	gui.io = imgui.CurrentIO()
	gui.io.SetBackendFlags(imgui.BackendFlagHasMouseCursors)
	gui.io.SetBackendFlags(imgui.BackendFlagHasSetMousePos)

	// TEMPORARY: Will use ZeptoForge key codes later
	gui.io.KeyMap(imgui.KeyTab, int(glfw.KeyTab))
	gui.io.KeyMap(imgui.KeyLeftArrow, int(glfw.KeyLeft))
	gui.io.KeyMap(imgui.KeyRightArrow, int(glfw.KeyRight))
	gui.io.KeyMap(imgui.KeyUpArrow, int(glfw.KeyUp))
	gui.io.KeyMap(imgui.KeyDownArrow, int(glfw.KeyDown))
	gui.io.KeyMap(imgui.KeyPageUp, int(glfw.KeyPageUp))
	gui.io.KeyMap(imgui.KeyPageDown, int(glfw.KeyPageDown))
	gui.io.KeyMap(imgui.KeyHome, int(glfw.KeyHome))
	gui.io.KeyMap(imgui.KeyEnd, int(glfw.KeyEnd))
	gui.io.KeyMap(imgui.KeyInsert, int(glfw.KeyInsert))
	gui.io.KeyMap(imgui.KeyDelete, int(glfw.KeyDelete))
	gui.io.KeyMap(imgui.KeyBackspace, int(glfw.KeyBackspace))
	gui.io.KeyMap(imgui.KeySpace, int(glfw.KeySpace))
	gui.io.KeyMap(imgui.KeyEnter, int(glfw.KeyEnter))
	gui.io.KeyMap(imgui.KeyEscape, int(glfw.KeyEscape))
	gui.io.KeyMap(imgui.KeyA, int(glfw.KeyA))
	gui.io.KeyMap(imgui.KeyC, int(glfw.KeyC))
	gui.io.KeyMap(imgui.KeyV, int(glfw.KeyV))
	gui.io.KeyMap(imgui.KeyX, int(glfw.KeyX))
	gui.io.KeyMap(imgui.KeyY, int(glfw.KeyY))
	gui.io.KeyMap(imgui.KeyZ, int(glfw.KeyZ))

	gui.renderer = NewOpenGL3(gui.io)
}

func (gui ImGuiLayer) OnDetach() {
	gui.renderer.Dispose()
	gui.context.Destroy()
}

func (gui *ImGuiLayer) OnUpdate() {
	gui.newFrame()
	display := application.DisplaySize()
	gui.io.SetDisplaySize(imgui.Vec2{X: display[0], Y: display[1]})
	imgui.NewFrame()
	imgui.ShowDemoWindow(&showDemoWindow)
	// imgui.Begin("Test ImGui")
	// imgui.End()
	imgui.Render()
	gui.renderer.PreRender(clearColor)
	gui.renderer.Render(display, application.FrameBufferSize(), imgui.RenderedDrawData())
}

func (gui *ImGuiLayer) newFrame() {
	// Setup time step
	currentTime := glfw.GetTime()
	if gui.time > 0 {
		gui.io.SetDeltaTime(float32(currentTime - gui.time))
	}
	gui.time = currentTime

	// Setup inputs
	win := application.GetWindow().GetWindow()

	for i := 0; i < len(gui.mouseJustPressed); i++ {
		down := gui.mouseJustPressed[i] || (win.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		gui.io.SetMouseButtonDown(i, down)
		gui.mouseJustPressed[i] = false
	}
}

func (gui *ImGuiLayer) OnEvent(e *event.Eventum) {
	dispatcher := event.NewEventDispatcher(e)
	dispatcher.Dispatch(event.EventFn{Event: &keyEvent.KeyPressedEvent{}, Fn: gui.onKeyPressedEvent})
	dispatcher.Dispatch(event.EventFn{Event: &keyEvent.KeyReleasedEvent{}, Fn: gui.onKeyReleasedEvent})
	dispatcher.Dispatch(event.EventFn{Event: &keyEvent.KeyTypedEvent{}, Fn: gui.onKeyTypedEvent})
	dispatcher.Dispatch(event.EventFn{Event: &mouseEvent.MouseButtonPressedEvent{}, Fn: gui.onMouseButtonPressedEvent})
	dispatcher.Dispatch(event.EventFn{Event: &mouseEvent.MouseButtonReleasedEvent{}, Fn: gui.onMouseButtonReleasedEvent})
	dispatcher.Dispatch(event.EventFn{Event: &mouseEvent.MouseMovedEvent{}, Fn: gui.onMouseMovedEvent})
	dispatcher.Dispatch(event.EventFn{Event: &mouseEvent.MouseScrolledEvent{}, Fn: gui.onMouseScrolledEvent})
	dispatcher.Dispatch(event.EventFn{Event: &appEvent.WindowResizeEvent{}, Fn: gui.onWindowResizedEvent})
}

func (gui *ImGuiLayer) onMouseButtonPressedEvent(e event.Eventum) bool {
	but := e.GetEvent().(*mouseEvent.MouseButtonPressedEvent).GetButton()

	gui.io.SetMouseButtonDown(but, true)
	return false
}

func (gui *ImGuiLayer) onMouseButtonReleasedEvent(e event.Eventum) bool {
	but := e.GetEvent().(*mouseEvent.MouseButtonReleasedEvent).GetButton()

	gui.io.SetMouseButtonDown(but, false)
	return false
}

func (gui *ImGuiLayer) onMouseMovedEvent(e event.Eventum) bool {
	ev := e.GetEvent().(*mouseEvent.MouseMovedEvent)
	xPos := float32(ev.GetX())
	yPos := float32(ev.GetY())

	gui.io.SetMousePosition(imgui.Vec2{X: xPos, Y: yPos})
	return false
}

func (gui *ImGuiLayer) onMouseScrolledEvent(e event.Eventum) bool {
	ev := e.GetEvent().(*mouseEvent.MouseScrolledEvent)
	xOff := float32(ev.GetXOffset())
	yOff := float32(ev.GetYOffset())

	gui.io.AddMouseWheelDelta(xOff, yOff)
	return false
}

func (gui *ImGuiLayer) onKeyPressedEvent(e event.Eventum) bool {
	key := e.GetEvent().(*keyEvent.KeyPressedEvent).GetKey()
	gui.io.KeyPress(key)

	gui.io.KeyCtrl(int(glfw.KeyLeftControl), int(glfw.KeyRightControl))
	gui.io.KeyShift(int(glfw.KeyLeftShift), int(glfw.KeyRightShift))
	gui.io.KeyAlt(int(glfw.KeyLeftAlt), int(glfw.KeyRightAlt))
	gui.io.KeySuper(int(glfw.KeyLeftSuper), int(glfw.KeyRightSuper))

	return false
}

func (gui *ImGuiLayer) onKeyReleasedEvent(e event.Eventum) bool {
	ev := e.GetEvent().(*keyEvent.KeyReleasedEvent)
	key := ev.GetKey()

	gui.io.KeyRelease(key)
	return false
}

func (gui *ImGuiLayer) onKeyTypedEvent(e event.Eventum) bool {
	ev := e.GetEvent().(*keyEvent.KeyTypedEvent)
	key := ev.GetKey()
	if key > 0 && key < 0x10000 {
		gui.io.AddInputCharacters(string(key))
	}
	return false
}

func (gui *ImGuiLayer) onWindowResizedEvent(e event.Eventum) bool {
	ev := e.GetEvent().(*appEvent.WindowResizeEvent)
	width := float32(ev.GetWidth())
	height := float32(ev.GetHeight())

	gui.io.SetDisplaySize(imgui.Vec2{X: width, Y: height})
	gl.Viewport(0, 0, int32(width), int32(height))
	return false
}
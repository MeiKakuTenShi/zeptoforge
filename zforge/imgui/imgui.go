package imgui

import (
	"math"

	imgui "github.com/inkyblackness/imgui-go"

	// TEMPORARY

	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	clearColor = [4]float32{0.0, 0.0, 0.0, 1.0}
)

type Imgui struct {
	io               imgui.IO
	window           *glfw.Window
	context          *imgui.Context
	renderer         *OpenGL3
	time             float64
	mouseJustPressed [3]bool
}

func NewImgui(window *glfw.Window) *Imgui {
	result := new(Imgui)
	result.context = imgui.CreateContext(nil)
	result.io = imgui.CurrentIO()

	result.io.SetConfigFlags(imgui.ConfigFlagNavEnableKeyboard) // Enable keyboard Controls
	// result.io.SetConfigFlags(imgui.ConfigFlagNavEnableGamepad)	// Enable Gamepad Controls
	// result.io.SetConfigFlags(imgui.ConfigFlagDockingEnable)		// Enable Docking (Not implemented in imgui-go)
	// result.io.SetConfigFlags(imgui.ConfigFlagViewportsEnable)	// Enable Multi-Viewport / PlatformWindows (Not implemented in imgui-go)
	// result.io.SetConfigFlags(imgui.ConfigFlagViewportNoTaskBarIcons)	// (Not implemented in imgui-go)
	// result.io.SetConfigFlags(imgui.ConfigFlagViewportsNoMerge) // (Not implemented in imgui-go)

	result.setKeyMapping()
	result.window = window
	result.renderer = &OpenGL3{glslVersion: "#version 410"}
	result.renderer.createDeviceObjects()

	return result
}

func (gui *Imgui) Destruct() {
	gui.renderer.Dispose()
	gui.context.Destroy()
}

func (gui *Imgui) Begin() {
	gui.newFrame()
	imgui.NewFrame()
}
func (gui *Imgui) End() {
	// Setup display size (every frame to accommodate for window resizing)
	width, height := gui.window.GetSize()
	buffWidth, buffHeight := gui.window.GetFramebufferSize()

	imgui.Render()
	gui.renderer.PreRender(clearColor)
	gui.renderer.Render([2]float32{float32(width), float32(height)}, [2]float32{float32(buffWidth), float32(buffHeight)}, imgui.RenderedDrawData())
}
func (gui *Imgui) ShowDemo(v *bool) {
	imgui.ShowDemoWindow(v)
}

func (gui *Imgui) newFrame() {
	// Setup display size (every frame to accommodate for window resizing)
	width, height := gui.window.GetSize()
	gui.io.SetDisplaySize(imgui.Vec2{X: float32(width), Y: float32(height)})

	// Setup time step
	currentTime := glfw.GetTime()
	if gui.time > 0 {
		gui.io.SetDeltaTime(float32(currentTime - gui.time))
	} else {
		gui.io.SetDeltaTime(float32(1 / 60))
	}
	gui.time = currentTime

	// Setup inputs
	if gui.window.GetAttrib(glfw.Focused) != 0 {
		x, y := gui.window.GetCursorPos()
		gui.io.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
	} else {
		gui.io.SetMousePosition(imgui.Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32})
	}

	for i := 0; i < len(gui.mouseJustPressed); i++ {
		down := gui.mouseJustPressed[i] || (gui.window.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		gui.io.SetMouseButtonDown(i, down)
		gui.mouseJustPressed[i] = false
	}
}

func (gui *Imgui) setKeyMapping() {
	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
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
}

// ClipboardText returns the current clipboard text, if available.
func (gui *Imgui) ClipboardText() string {
	return gui.window.GetClipboardString()
}

// SetClipboardText sets the text as the current clipboard text.
func (gui *Imgui) SetClipboardText(text string) {
	gui.window.SetClipboardString(text)
}

var glfwButtonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: 0,
	glfw.MouseButton2: 1,
	glfw.MouseButton3: 2,
}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	0: glfw.MouseButton1,
	1: glfw.MouseButton2,
	2: glfw.MouseButton3,
}

package imgui

// import (
// 	"github.com/go-gl/glfw/v3.3/glfw"
// )

// func  installCallbacks() {
// 	platform.window.SetMouseButtonCallback(platform.mouseButtonChange)
// 	platform.window.SetScrollCallback(platform.mouseScrollChange)
// 	platform.window.SetKeyCallback(platform.keyChange)
// 	platform.window.SetCharCallback(platform.charChange)
// }

// func (platform *GLFW) mouseButtonChange(window *glfw.Window, rawButton glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
// 	buttonIndex, known := glfwButtonIndexByID[rawButton]

// 	if known && (action == glfw.Press) {
// 		platform.mouseJustPressed[buttonIndex] = true
// 	}
// }

// func (platform *GLFW) mouseScrollChange(window *glfw.Window, x, y float64) {
// 	platform.imguiIO.AddMouseWheelDelta(float32(x), float32(y))
// }

// func (platform *GLFW) keyChange(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
// 	if action == glfw.Press {
// 		platform.imguiIO.KeyPress(int(key))
// 	}
// 	if action == glfw.Release {
// 		platform.imguiIO.KeyRelease(int(key))
// 	}

// 	// Modifiers are not reliable across systems
// 	platform.imguiIO.KeyCtrl(int(glfw.KeyLeftControl), int(glfw.KeyRightControl))
// 	platform.imguiIO.KeyShift(int(glfw.KeyLeftShift), int(glfw.KeyRightShift))
// 	platform.imguiIO.KeyAlt(int(glfw.KeyLeftAlt), int(glfw.KeyRightAlt))
// 	platform.imguiIO.KeySuper(int(glfw.KeyLeftSuper), int(glfw.KeyRightSuper))
// }

// func (platform *GLFW) charChange(window *glfw.Window, char rune) {
// 	platform.imguiIO.AddInputCharacters(string(char))
// }

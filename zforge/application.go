package zforge

import (
	"fmt"
	"runtime"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows/winInput"
	"github.com/MeiKakuTenShi/zeptoforge/platform/windows"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/renderer"
	"github.com/go-gl/gl/v4.6-core/gl"
)

var (
	appInstance *appSingleton
)

type appSingleton struct {
	window  Window
	gui     *ImGuiLayer
	stack   *LayerStack
	shader  renderer.Shader
	running bool
	Name    string

	vbo renderer.VertexBuffer
	ibo renderer.IndexBuffer
	vao uint32
}

func InitApp(name string) {
	if appInstance == nil {
		initLogSys()
		appInstance = &appSingleton{}
		// Windows specific build
		if runtime.GOOS == "windows" {
			p := &windows.WinData{Title: name, Width: 0, Height: 0}

			appInstance.window = windows.NewWinWindow(p)
			appInstance.gui = NewImGuiLayer(appInstance.window.GetWindow())
			initInputSingleton(&winInput.WindowsInput{}) // singleton can only be set once
		} else {
			panic("Platform currently not supported")
		}

		//-----------------------
		appInstance.stack = newLayerStack()
		appInstance.running = true
		appInstance.Name = name

		PushOverlayOnApp(NewLayem(appInstance.gui, "ImGuiLayer"))
		appInstance.window.SetEventCallback(AppOnEvent)

		shader, err := renderer.NewShader(vertexShader, fragmentShader)
		if err != nil {
			core_ERROR(err)
		}
		appInstance.shader = shader
		appInstance.shader.Bind()

		gl.GenVertexArrays(1, &appInstance.vao)
		gl.BindVertexArray(appInstance.vao)

		appInstance.vbo = renderer.CreateVertexBuffer(vertices, len(vertices)*4)
		appInstance.ibo = renderer.CreateIndexBuffer(indices, len(indices)*4)

		vertPos := renderer.NewBuffem(renderer.Float3, "a_Position")
		col := renderer.NewBuffem(renderer.Float4, "a_Color")
		norm := renderer.NewBuffem(renderer.Float3, "a_Normal")

		layout := renderer.NewBufferLayout(vertPos, col, norm)
		fmt.Println(layout)
		// appInstance.vbo.SetLayout(layout)

		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)

		core_INFO(fmt.Sprintf("Created App: %s (Width %v, Height %v)", appInstance.Name, appInstance.window.GetWidth(), appInstance.window.GetHeight()))
	}
}

func DisplaySize() [2]float32 {
	return [2]float32{float32(appInstance.window.GetWidth()), float32(appInstance.window.GetHeight())}
}

func FrameBufferSize() [2]float32 {
	return appInstance.window.FramebufferSize()
}

func GetApplication() appSingleton {
	return *appInstance
}

func GetWindow() Window {
	return appInstance.window
}

func PushLayerOnApp(l *Layem) {
	appInstance.stack.PushLayer(l)
	l.layer.OnAttach()
}

func PushOverlayOnApp(l *Layem) {
	appInstance.stack.PushOverlay(l)
	l.layer.OnAttach()
}

func AppOnEvent(e *event.Eventum) {
	core_INFO(e.String())

	dispatcher := event.NewEventDispatcher(e)
	dispatcher.Dispatch(event.EventFn{Event: event.WindowClose, Fn: onWindowClose})

	for _, v := range appInstance.stack.layers {
		v.layer.OnEvent(e)
		if e.Done() {
			break
		}
	}
}

func RunApp() {
	defer closeApp()
	for appInstance.running {
		gl.ClearColor(0.0, 0.0, 0.0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Render
		appInstance.shader.Bind()
		gl.BindVertexArray(appInstance.vao)
		gl.DrawElements(gl.TRIANGLES, int32(appInstance.ibo.GetCount()), gl.UNSIGNED_INT, nil)

		for _, v := range appInstance.stack.layers {
			v.layer.OnUpdate()
		}

		appInstance.gui.Begin()
		for _, v := range appInstance.stack.layers {
			v.layer.OnImGuiRender()
		}
		appInstance.gui.End()
		appInstance.window.OnUpdate()
	}
}
func closeApp() {
	for _, v := range appInstance.stack.layers {
		v.layer.OnDetach()
	}

	appInstance.window.Destruct()
	appInstance.shader.Unbind()
}

func onWindowClose(e event.Eventum) bool {
	appInstance.running = false
	return true
}

var vertices = []float32{
	-0.5, -0.5, 0.0,
	0.5, -0.5, 0.0,
	0.0, 0.5, 0.0,
}

var indices = []float32{0, 1, 2}

var vertexShader = `
#version 330 core

layout(location = 0) in vec3 aPos;

out vec3 vPos;

void main() {
	vPos = aPos;
    gl_Position = vec4(aPos, 1);
}` + "\x00"

var fragmentShader = `
#version 330 core

layout(location = 0) out vec4 color;

in vec3 vPos;

void main() {
    color = vec4(vPos * 0.5 + 0.5, 1.0);
}` + "\x00"

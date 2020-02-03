package zforge

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"

	"github.com/MeiKakuTenShi/zeptoforge/ZeptoForge/Platform/Windows/winInput"
	"github.com/MeiKakuTenShi/zeptoforge/platform/windows"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/renderer"
	"github.com/go-gl/gl/v4.5-core/gl"
)

var (
	appInstance appSingleton
)

type appSingleton struct {
	window  Window
	gui     *ImGuiLayer
	stack   *LayerStack
	running bool
	Name    string

	shader renderer.Shader
	vbo    *renderer.VertBuff
	ibo    *renderer.IndBuff
	vao    *renderer.VertArray

	shader2    renderer.Shader
	square_vao *renderer.VertArray
}

func appInstanceNil() bool {
	return reflect.DeepEqual(appInstance, appSingleton{})
}

func InitApp(name string) {
	if appInstanceNil() {
		initLogSys()
		// ------- Windows specific build ----------------------------------------------
		if runtime.GOOS == "windows" {
			p := &windows.WinData{Title: name, Width: 0, Height: 0}
			appInstance.window = windows.NewWinWindow(p)
			appInstance.gui = NewImGuiLayer(appInstance.window.GetWindow())
			initInputSingleton(&winInput.WindowsInput{}) // singleton can only be set once
		} else {
			panic("Platform currently not supported")
		}

		//-------------------------------------------------------------------------------
		appInstance.stack = newLayerStack()
		appInstance.running = true
		appInstance.Name = name

		PushOverlayOnApp(NewLayem(appInstance.gui, "ImGuiLayer"))
		appInstance.window.SetEventCallback(AppOnEvent)

		////////////////////////////////////////////////
		////////////////////////////////////////////////
		//////////// TESTING TRIANGLES /////////////////
		///////////////////BEGIN////////////////////////
		////////////////////////////////////////////////

		shader, err := renderer.NewShader(vertexShader, fragmentShader)
		if err != nil {
			core_ERROR("Failed to create Shader Program - ", err)
		}
		appInstance.shader = shader
		appInstance.shader.Bind()

		appInstance.vao, err = renderer.NewVertexArray()
		if err != nil {
			core_ERROR("Failed to create Vertex Array - ", err)
		}

		appInstance.vbo, err = renderer.NewVertexBuffer(vertices, len(vertices)*4)
		if err != nil {
			core_ERROR("Failed to create Vertex Buffer - ", err)
		}

		pos := renderer.NewBuffElem(renderer.Float3, "aPosition", false)
		col := renderer.NewBuffElem(renderer.Float4, "aColor", false)

		layout := renderer.NewBufferLayout(pos, col)
		// fmt.Println(layout)
		appInstance.vbo.SetLayout(layout)

		err = appInstance.vao.AddVertexBuffer(*appInstance.vbo)
		if err != nil {
			core_INFO(err)
		}

		appInstance.ibo, err = renderer.NewIndexBuffer(indices, len(indices))
		if err != nil {
			core_ERROR("Failed to create Index Buffer - ", err)
		}
		appInstance.vao.SetIndexBuffer(*appInstance.ibo)

		appInstance.shader2, err = renderer.NewShader(vertexShader2, fragmentShader2)
		if err != nil {
			core_ERROR("Failed to create Shader Program - ", err)
		}

		appInstance.square_vao, err = renderer.NewVertexArray()
		if err != nil {
			core_INFO("Failed to create Vertex Array - ", err)
		}
		squareVB, err := renderer.NewVertexBuffer(vertices2, len(vertices2)*4)
		if err != nil {
			core_INFO("Failed to create Vertex Buffer - ", err)
		}

		squareVBlayout := renderer.NewBufferLayout(pos)
		squareVB.SetLayout(squareVBlayout)

		err = appInstance.square_vao.AddVertexBuffer(*squareVB)
		if err != nil {
			core_INFO(err)
		}

		squareIB, err := renderer.NewIndexBuffer(square_indices, len(square_indices)*4)
		if err != nil {
			core_INFO(err)
		}
		appInstance.square_vao.SetIndexBuffer(*squareIB)
		////////////////////////////////////////////////
		////////////////////////////////////////////////
		//////////// TESTING TRIANGLES /////////////////
		////////////////// END /////////////////////////
		////////////////////////////////////////////////

		core_INFO(fmt.Sprintf("Created App: %s (Width %v, Height %v)", appInstance.Name, appInstance.window.GetWidth(), appInstance.window.GetHeight()))
	}
}

func DisplaySize() [2]float32 {
	return [2]float32{float32(appInstance.window.GetWidth()), float32(appInstance.window.GetHeight())}
}

func FrameBufferSize() [2]float32 {
	return appInstance.window.FramebufferSize()
}

func GetApplication() (appSingleton, error) {
	if appInstanceNil() {
		return appSingleton{}, errors.New("application has not been created")
	}
	return appInstance, nil
}

func GetWindow() (Window, error) {
	if appInstanceNil() {
		return nil, errors.New("cannot obtain window, application is nil")
	}
	return appInstance.window, nil
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
		appInstance.shader2.Bind()
		appInstance.square_vao.Bind()
		gl.DrawElements(gl.TRIANGLES, appInstance.square_vao.GetIndexBuffer().GetCount(), gl.UNSIGNED_INT, nil)

		appInstance.shader.Bind()
		appInstance.vao.Bind()
		gl.DrawElements(gl.TRIANGLES, appInstance.ibo.GetCount(), gl.UNSIGNED_INT, nil)

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
	appInstance.ibo.Remove()
	appInstance.vbo.Remove()
}

func onWindowClose(e event.Eventum) bool {
	appInstance.running = false
	return true
}

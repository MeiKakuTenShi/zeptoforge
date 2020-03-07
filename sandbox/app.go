package sandbox

import (
	"github.com/MeiKakuTenShi/zeptoforge/zforge"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/renderer"
	"github.com/go-gl/mathgl/mgl32"
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
	shader renderer.Shader
	vao    *renderer.VertArray

	shader2    renderer.Shader
	square_vao *renderer.VertArray

	camera      *renderer.OrthoCam
	cameraPos   mgl32.Vec3
	cameraSpeed float32

	cameraRotation      float64
	cameraRotationSpeed float64

	squarePos   mgl32.Vec3
	squareSpeed float64
}

func (ex *ExLayer) OnAttach() {
	////////////////////////////////////////////////
	////////////////////////////////////////////////
	//////////// TESTING TRIANGLES /////////////////
	///////////////////BEGIN////////////////////////
	////////////////////////////////////////////////
	var err error

	ex.cameraPos = mgl32.Vec3{0, 0, 0}
	ex.cameraRotation = 0
	ex.cameraSpeed = 2
	ex.cameraRotationSpeed = 10

	ex.squarePos = mgl32.Vec3{0, 0, 0}
	ex.squareSpeed = 1

	ex.camera = renderer.NewOrtho2DCamera(-1.6, 1.6, -0.9, 0.9)

	pos := renderer.NewBuffElem(renderer.Float3, "aPosition", false)
	col := renderer.NewBuffElem(renderer.Float4, "aColor", false)

	// Triangle
	ex.shader, err = renderer.NewShader(renderer.VertexShader, renderer.FragmentShader)
	if err != nil {
		zforge.ZF_INFO("Failed to create Shader Program - ", err)
	}
	ex.shader.Bind()

	ex.vao, err = renderer.NewVertexArray()
	if err != nil {
		zforge.ZF_INFO("Failed to create Vertex Array - ", err)
	}

	vbo, err := renderer.NewVertexBuffer(renderer.Vertices, len(renderer.Vertices)*4)
	if err != nil {
		zforge.ZF_INFO("Failed to create Vertex Buffer - ", err)
	}

	layout := renderer.NewBufferLayout(pos, col)
	vbo.SetLayout(layout)

	err = ex.vao.AddVertexBuffer(*vbo)
	if err != nil {
		zforge.ZF_INFO(err)
	}

	ibo, err := renderer.NewIndexBuffer(renderer.Indices, len(renderer.Indices))
	if err != nil {
		zforge.ZF_INFO("Failed to create Index Buffer - ", err)
	}
	ex.vao.SetIndexBuffer(*ibo)

	// Square

	ex.shader2, err = renderer.NewShader(renderer.VertexShader2, renderer.FragmentShader2)
	if err != nil {
		zforge.ZF_INFO("Failed to create Shader Program - ", err)
	}

	ex.square_vao, err = renderer.NewVertexArray()
	if err != nil {
		zforge.ZF_INFO("Failed to create Vertex Array - ", err)
	}
	squareVB, err := renderer.NewVertexBuffer(renderer.Vertices2, len(renderer.Vertices2)*4)
	if err != nil {
		zforge.ZF_INFO("Failed to create Vertex Buffer - ", err)
	}

	squareVBlayout := renderer.NewBufferLayout(pos)
	squareVB.SetLayout(squareVBlayout)

	err = ex.square_vao.AddVertexBuffer(*squareVB)
	if err != nil {
		zforge.ZF_INFO(err)
	}

	squareIB, err := renderer.NewIndexBuffer(renderer.Square_indices, len(renderer.Square_indices)*4)
	if err != nil {
		zforge.ZF_INFO(err)
	}
	ex.square_vao.SetIndexBuffer(*squareIB)
	////////////////////////////////////////////////
	////////////////////////////////////////////////
	//////////// TESTING TRIANGLES /////////////////
	////////////////// END /////////////////////////
	////////////////////////////////////////////////
}

func (l ExLayer) OnDetach() {
	l.shader.Unbind()
	l.shader2.Unbind()

	l.vao.Destruct()
	l.square_vao.Destruct()
}

func (l ExLayer) OnImGuiRender() {
}

func (l *ExLayer) OnUpdate(ts zforge.TimeStep) {
	step := ts.GetSeconds()

	if zforge.IsKeyPressed(int(zforge.ZF_KeyLeft)) {
		l.cameraPos = l.cameraPos.Sub(mgl32.Vec3{l.cameraSpeed * float32(step), 0, 0})
	} else if zforge.IsKeyPressed(int(zforge.ZF_KeyRight)) {
		l.cameraPos = l.cameraPos.Add(mgl32.Vec3{l.cameraSpeed * float32(step), 0, 0})
	}
	if zforge.IsKeyPressed(int(zforge.ZF_KeyUp)) {
		l.cameraPos = l.cameraPos.Add(mgl32.Vec3{0, l.cameraSpeed * float32(step), 0})
	} else if zforge.IsKeyPressed(int(zforge.ZF_KeyDown)) {
		l.cameraPos = l.cameraPos.Sub(mgl32.Vec3{0, l.cameraSpeed * float32(step), 0})
	}

	if zforge.IsKeyPressed(int(zforge.ZF_KeyJ)) {
		l.squarePos = l.squarePos.Sub(mgl32.Vec3{float32(l.squareSpeed * step), 0, 0})
	} else if zforge.IsKeyPressed(int(zforge.ZF_KeyL)) {
		l.squarePos = l.squarePos.Add(mgl32.Vec3{float32(l.squareSpeed * step), 0, 0})
	}
	if zforge.IsKeyPressed(int(zforge.ZF_KeyI)) {
		l.squarePos = l.squarePos.Add(mgl32.Vec3{0, float32(l.squareSpeed * step), 0})
	} else if zforge.IsKeyPressed(int(zforge.ZF_KeyK)) {
		l.squarePos = l.squarePos.Sub(mgl32.Vec3{0, float32(l.squareSpeed * step), 0})
	}

	if zforge.IsKeyPressed(int(zforge.ZF_KeyA)) {
		l.cameraRotation += l.cameraRotationSpeed * step
	}
	if zforge.IsKeyPressed(int(zforge.ZF_KeyD)) {
		l.cameraRotation -= l.cameraRotationSpeed * step
	}

	renderer.SetClearColor([]float32{0, 0, 0, 1})
	renderer.Clear()

	l.camera.SetPosition(l.cameraPos)
	l.camera.SetRotation(l.cameraRotation)

	// Render
	renderer.BeginScene(*l.camera)

	scale := mgl32.Scale3D(0.1, 0.1, 0.1)

	for x := float32(0); x < 20; x++ {
		for y := float32(0); y < 20; y++ {
			pos := mgl32.Vec3{x * 0.11, y * 0.11, 0}
			transform := mgl32.Translate3D(pos.Elem()).Mul4(scale)
			renderer.Submit(l.square_vao, l.shader2, transform)
		}
	}

	// renderer.Submit(l.vao, l.shader)

	renderer.EndScene()
}

func (l *ExLayer) OnEvent(e *event.Eventum) {

}

package sandbox

import (
	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
	"github.com/MeiKakuTenShi/zeptoforge/zforge"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/event"
	"github.com/MeiKakuTenShi/zeptoforge/zforge/renderer"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/inkyblackness/imgui-go"
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
	shader *renderer.ZFshader
	vao    *renderer.VertArray

	shader2, textureShader *renderer.ZFshader
	square_vao             *renderer.VertArray

	testTexture, alphaTexture *renderer.Texture2D

	camera      *renderer.OrthoCam
	cameraPos   mgl32.Vec3
	cameraSpeed float32

	cameraRotation      float64
	cameraRotationSpeed float64

	squarePos   mgl32.Vec3
	squareSpeed float64

	squareColor [3]float32
}

func (ex *ExLayer) OnAttach() {
	////////////////////////////////////////////////
	////////////////////////////////////////////////
	//////////// TESTING TRIANGLES /////////////////
	///////////////////BEGIN////////////////////////
	////////////////////////////////////////////////
	var err error

	ex.squareColor = [3]float32{0.2, 0.3, 0.8}

	ex.cameraPos = mgl32.Vec3{0, 0, 0}
	ex.cameraRotation = 0
	ex.cameraSpeed = 2
	ex.cameraRotationSpeed = 10

	ex.squarePos = mgl32.Vec3{0, 0, 0}
	ex.squareSpeed = 1

	ex.camera = renderer.NewOrtho2DCamera(-1.6, 1.6, -0.9, 0.9)

	pos := renderer.NewBuffElem(renderer.Float3, "aPosition", false)
	tex := renderer.NewBuffElem(renderer.Float2, "aTexCoord", false)
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
		zforge.ZF_INFO("Failed to create Shader - ", err)
	}

	ex.textureShader, err = renderer.NewShader(renderer.TextureShaderVert, renderer.TextureShaderFrag)
	if err != nil {
		zforge.ZF_INFO("Failed to create Shader - ", err)
	}

	ex.square_vao, err = renderer.NewVertexArray()
	if err != nil {
		zforge.ZF_INFO("Failed to create Vertex Array - ", err)
	}
	squareVB, err := renderer.NewVertexBuffer(renderer.Vertices2, len(renderer.Vertices2)*4)
	if err != nil {
		zforge.ZF_INFO("Failed to create Vertex Buffer - ", err)
	}

	squareVBlayout := renderer.NewBufferLayout(pos, tex)
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

	ex.testTexture, err = renderer.NewTexture2D("test.png")
	if err != nil {
		zforge.ZF_INFO("Failed to create texture - ", err)
	}

	ex.textureShader.Bind()

	shader, err := ex.textureShader.GetShader()
	if err != nil {
		zforge.ZF_INFO("Failed to get shader - ", err)
	}

	glShader, ok := shader.(*opengl.OpenGLShader)
	if !ok {
		zforge.ZF_INFO("Failed to assert shader type")
	}

	glShader.UploadUniformInt("uTexture", 0)
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

func (l *ExLayer) OnImGuiRender() {
	imgui.Begin("Settings")
	imgui.SliderFloat3("Square Color", &l.squareColor, 0, 1)
	imgui.End()
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

	// redColor := mgl32.Vec4{0.8, 0.2, 0.3, 1.0}
	// blueColor := mgl32.Vec4{0.2, 0.3, 0.8, 1.0}

	l.shader2.Bind()

	shader, err := l.shader2.GetShader()
	if err != nil {
		zforge.ZF_INFO("Failed to get shader - ", err)
	}

	glShader, ok := shader.(*opengl.OpenGLShader)
	if !ok {
		zforge.ZF_INFO("Failed to assert shader type")
	}

	glShader.UploadUniformFloat3("uColor", mgl32.Vec3{l.squareColor[0], l.squareColor[1], l.squareColor[2]})

	for x := float32(0); x < 10; x++ {
		for y := float32(0); y < 10; y++ {
			pos := mgl32.Vec3{x * 0.11, y * 0.11, 0}
			transform := mgl32.Translate3D(pos.Elem()).Mul4(scale)

			renderer.Submit(l.square_vao, l.shader2, transform)
		}
	}
	l.testTexture.Bind()
	renderer.Submit(l.square_vao, l.textureShader, mgl32.Diag4(mgl32.Vec4{1.5, 1.5, 1.5, 1}))

	// renderer.Submit(l.vao, l.shader)

	renderer.EndScene()
}

func (l *ExLayer) OnEvent(e *event.Eventum) {

}

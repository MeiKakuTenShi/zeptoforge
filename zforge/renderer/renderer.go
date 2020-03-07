package renderer

import "github.com/go-gl/mathgl/mgl32"

func BeginScene(cam OrthoCam) {
	scene.ViewProjectionMatrix = cam.GetViewProjectionMatrix()
}

func EndScene() {
}

func SetClearColor(color []float32) {
	static_API.renderer.SetClearColor(color)
}

func Clear() {
	static_API.renderer.Clear()
}

func Submit(va *VertArray, s Shader, mats ...mgl32.Mat4) {
	s.Bind()

	if len(mats) == 0 {
		s.UploadUniformMat4("transform", mgl32.Ident4())
	} else {
		s.UploadUniformMat4("transform", mats[0])
	}
	s.UploadUniformMat4("viewProjection", scene.ViewProjectionMatrix)

	va.Bind()
	static_API.renderer.DrawIndexed(va.GetIndexBuffer().GetCount())
}

type sceneData struct {
	ViewProjectionMatrix mgl32.Mat4
}

var scene sceneData

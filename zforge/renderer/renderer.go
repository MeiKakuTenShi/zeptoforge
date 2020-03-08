package renderer

import (
	"errors"

	"github.com/MeiKakuTenShi/zeptoforge/platform/opengl"
	"github.com/go-gl/mathgl/mgl32"
)

func BeginScene(cam OrthoCam) {
	scene.ViewProjectionMatrix = cam.GetViewProjectionMatrix()
}

func EndScene() {
}

func SetClearColor(color []float32) {
	sAPI.renderer.SetClearColor(color)
}

func Clear() {
	sAPI.renderer.Clear()
}

func Submit(va *VertArray, s *ZFshader, mats ...mgl32.Mat4) error {
	shader, err := s.GetShader()
	if err != nil {
		return err
	}

	shader.Bind()

	glShader, ok := shader.(*opengl.OpenGLShader)
	if !ok {
		return errors.New("shader type assertion failure")
	}

	if len(mats) == 0 {
		glShader.UploadUniformMat4("transform", mgl32.Ident4())
	} else {
		glShader.UploadUniformMat4("transform", mats[0])
	}
	glShader.UploadUniformMat4("viewProjection", scene.ViewProjectionMatrix)

	va.Bind()
	sAPI.renderer.DrawIndexed(va.GetIndexBuffer().GetCount())

	return nil
}

type sceneData struct {
	ViewProjectionMatrix mgl32.Mat4
}

var scene sceneData

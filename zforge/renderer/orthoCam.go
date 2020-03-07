package renderer

import (
	"github.com/go-gl/mathgl/mgl32"
)

type OrthoCam struct {
	projectionMatrix     mgl32.Mat4
	viewMatrix           mgl32.Mat4
	viewProjectionMatrix mgl32.Mat4
	position             mgl32.Vec3
	rotation             float64
}

func NewOrtho2DCamera(left, right, bottom, top float32) *OrthoCam {
	r := new(OrthoCam)
	r.projectionMatrix = mgl32.Ortho2D(left, right, bottom, top)
	r.viewMatrix = mgl32.Ident4()
	r.viewProjectionMatrix = r.projectionMatrix.Mul4(r.viewMatrix)

	return r
}

func (cam OrthoCam) GetPosition() mgl32.Vec3 {
	return cam.position
}
func (cam *OrthoCam) SetPosition(position mgl32.Vec3) {
	cam.position = position
	cam.recalculateViewMatrix()
}
func (cam OrthoCam) GetRotation() float64 {
	return cam.rotation
}
func (cam *OrthoCam) SetRotation(rotation float64) {
	cam.rotation = rotation
	cam.recalculateViewMatrix()
}

func (cam OrthoCam) GetProjectionMatrix() mgl32.Mat4 {
	return cam.projectionMatrix
}
func (cam OrthoCam) GetViewMatrix() mgl32.Mat4 {
	return cam.viewMatrix
}
func (cam OrthoCam) GetViewProjectionMatrix() mgl32.Mat4 {
	return cam.viewProjectionMatrix
}

func (cam *OrthoCam) recalculateViewMatrix() {
	transform := mgl32.Translate3D(cam.position[0], cam.position[1], cam.position[2]).Mul4(mgl32.HomogRotate3D(float32(cam.rotation), mgl32.Vec3{0, 0, 1}))

	cam.viewMatrix = transform.Inv()
	cam.viewProjectionMatrix = cam.projectionMatrix.Mul4(cam.viewMatrix)
}

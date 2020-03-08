package opengl

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type OpenGLShader struct {
	rendererID uint32
}

func (s *OpenGLShader) Init(vSrc, fSrc string) error {
	vertexShader, err := compileShader(vSrc, gl.VERTEX_SHADER)
	if err != nil {
		return err
	}

	fragmentShader, err := compileShader(fSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		gl.DeleteProgram(program)

		gl.DeleteShader(vertexShader)
		gl.DeleteShader(fragmentShader)

		return fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	s.rendererID = program

	return nil
}

func (s OpenGLShader) Dispose() {
	gl.DeleteProgram(s.rendererID)
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		gl.DeleteShader(shader)

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (s OpenGLShader) Bind() {
	gl.UseProgram(s.rendererID)
}

func (s OpenGLShader) Unbind() {
	gl.UseProgram(0)
}

func (s OpenGLShader) GetID() uint32 {
	return s.rendererID
}

func (s OpenGLShader) UploadUniformInt(name string, v int32) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.Uniform1i(location, v)
}
func (s OpenGLShader) UploadUniformFloat(name string, v float32) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.Uniform1f(location, v)
}
func (s OpenGLShader) UploadUniformFloat2(name string, v mgl32.Vec2) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.Uniform2f(location, v.X(), v.Y())
}
func (s OpenGLShader) UploadUniformFloat3(name string, v mgl32.Vec3) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.Uniform3f(location, v.X(), v.Y(), v.Z())
}
func (s OpenGLShader) UploadUniformFloat4(name string, v mgl32.Vec4) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.Uniform4f(location, v.X(), v.Y(), v.Z(), v.W())
}
func (s OpenGLShader) UploadUniformMat3(name string, mat mgl32.Mat3) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.UniformMatrix3fv(location, 1, false, &mat[0])
}
func (s OpenGLShader) UploadUniformMat4(name string, mat mgl32.Mat4) {
	location := gl.GetUniformLocation(s.rendererID, gl.Str(fmt.Sprint(name, "\x00")))
	gl.UniformMatrix4fv(location, 1, false, &mat[0])
}

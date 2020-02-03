package renderer

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type Shader struct {
	rendererID uint32
}

func NewShader(vSrc, fSrc string) (Shader, error) {
	r := new(Shader)

	vertexShader, err := compileShader(vSrc, gl.VERTEX_SHADER)
	if err != nil {
		return *r, err
	}

	fragmentShader, err := compileShader(fSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return *r, err
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

		return Shader{}, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	r.rendererID = program

	return *r, nil
}

func (s Shader) Dispose() {
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

func (s Shader) Bind() {
	gl.UseProgram(s.rendererID)
}

func (s Shader) Unbind() {
	gl.UseProgram(0)
}

func (s Shader) GetID() uint32 {
	return s.rendererID
}

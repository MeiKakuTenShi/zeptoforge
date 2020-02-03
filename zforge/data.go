package zforge

var vertices = []float32{
	-0.5, -0.5, 0.0, 0.8, 0.2, 0.8, 1.0,
	0.5, -0.5, 0.0, 0.2, 0.2, 0.8, 1.0,
	0.0, 0.5, 0.0, 0.8, 0.8, 0.2, 1.0,
}

var vertices2 = []float32{
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0.5, 0.5, 0,
	-0.5, 0.5, 0,
}

var indices = []uint32{0, 1, 2}

var square_indices = []uint32{0, 1, 2, 2, 3, 0}

var vertexShader = `
#version 410 core

layout(location = 0) in vec3 aPosition;
layout(location = 1) in vec4 aColor;

out vec4 vColor;

void main() {
	vColor = aColor;
    gl_Position = vec4(aPosition, 1.0);
}` + "\x00"

var fragmentShader = `
#version 410 core

layout(location = 0) out vec4 color;

in vec4 vColor;

void main() {
	color = vColor;
}` + "\x00"

var vertexShader2 = `
#version 410 core

layout(location = 0) in vec3 a_Position;

out vec3 v_Position;

void main() {
	v_Position = a_Position;
    gl_Position = vec4(a_Position, 1.0);
}` + "\x00"

var fragmentShader2 = `
#version 410 core

layout(location = 0) out vec4 color;

in vec3 v_Position;

void main() {
	color = vec4(0.2, 0.3, 0.8, 1.0);
}` + "\x00"

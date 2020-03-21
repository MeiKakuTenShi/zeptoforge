package renderer

var Vertices = []float32{
	-0.5, -0.5, 0.0, 0.8, 0.2, 0.8, 1.0,
	0.5, -0.5, 0.0, 0.2, 0.2, 0.8, 1.0,
	0.0, 0.5, 0.0, 0.8, 0.8, 0.2, 1.0,
}

var Vertices2 = []float32{
	-0.5, -0.5, 0, 0, 0,
	0.5, -0.5, 0, 1, 0,
	0.5, 0.5, 0, 1, 1,
	-0.5, 0.5, 0, 0, 1,
}

var Indices = []uint32{0, 1, 2}

var Square_indices = []uint32{0, 1, 2, 2, 3, 0}

var VertexShader = `
#version 410 core

layout(location = 0) in vec3 aPosition;
layout(location = 1) in vec4 aColor;

uniform mat4 viewProjection;
uniform mat4 transform;

out vec4 vColor;

void main() {
	vColor = aColor;
    gl_Position = viewProjection * transform * vec4(aPosition, 1.0);
}` + "\x00"

var FragmentShader = `
#version 410 core

layout(location = 0) out vec4 color;

in vec4 vColor;

void main() {
	color = vColor;
}` + "\x00"

var VertexShader2 = `
#version 410 core

layout(location = 0) in vec3 aPosition;

uniform mat4 viewProjection;
uniform mat4 transform;

out vec3 vPosition;

void main() {
	vPosition = aPosition;
    gl_Position = viewProjection * transform * vec4(aPosition, 1.0);
}` + "\x00"

var FragmentShader2 = `
#version 410 core

layout(location = 0) out vec4 color;

in vec3 vPosition;

uniform vec3 uColor;

void main() {
	color = vec4(uColor, 1.0);
}` + "\x00"

var TextureShaderVert = `
#version 410 core

layout(location = 0) in vec3 aPosition;
layout(location = 1) in vec2 aTexCoord;

uniform mat4 viewProjection;
uniform mat4 transform;

out vec2 vTexCoord;

void main() {
	vTexCoord = aTexCoord;
    gl_Position = viewProjection * transform * vec4(aPosition, 1.0);
}` + "\x00"

var TextureShaderFrag = `
#version 410 core

layout(location = 0) out vec4 color;

in vec2 vTexCoord;

uniform sampler2D uTexture;

void main() {
	color = texture(uTexture, vTexCoord);
}` + "\x00"

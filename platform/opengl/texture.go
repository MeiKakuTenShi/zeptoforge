package opengl

import (
	"fmt"
	"image"
	_ "image/png"
	"os"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type OpenGLTexture2D struct {
	rendererID uint32

	path          string
	width, height int32
}

func (t *OpenGLTexture2D) Init(filepath string) error {
	t.path = filepath

	infile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		return err
	}

	bounds := src.Bounds()
	t.width, t.height = int32(bounds.Max.X), int32(bounds.Max.Y)

	rgba := image.NewRGBA(bounds)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return fmt.Errorf("unsupported stride")
	}

	for i := 0; i < int(t.height); i++ {
		for j := 0; j < int(t.width); j++ {
			rgba.Set(j, i, src.At(j, int(t.height)-i))
		}
	}
	// draw.Draw(rgba, rgba.Bounds(), src, image.Point{0, 0}, draw.Src)

	gl.GenTextures(1, &t.rendererID)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, t.rendererID)

	// gl.CreateTextures(gl.TEXTURE_2D, 1, &t.rendererID)
	// gl.TextureStorage2D(t.rendererID, 1, gl.RGBA, t.width, t.height)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	// gl.TextureSubImage2D(t.rendererID, 0, 0, 0, t.width, t.height, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, t.width, t.height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))

	return nil
}

func (t OpenGLTexture2D) GetWidth() int32 {
	return t.width
}

func (t OpenGLTexture2D) GetHeight() int32 {
	return t.height
}

func (t *OpenGLTexture2D) Destruct() {
	gl.DeleteTextures(1, &t.rendererID)
}

func (t OpenGLTexture2D) Bind(slot uint32) {
	gl.BindTextureUnit(slot, t.rendererID)
}

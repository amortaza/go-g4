package g4

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	_ "image/png"
	_ "image/jpeg"
	"os"
	"image"
	"image/draw"
	"unsafe"
)

type Texture struct {
	TextureId     uint32

	Width, Height int

	textureUnit   uint32
}

func NewTexture() *Texture {
	t := &Texture{}

	gl.GenTextures(1, &t.TextureId)

	return t
}

func (t *Texture) LoadImage(filename string) {

	rgba, img := loadRGBA(filename)

	draw.Draw(rgba, rgba.Bounds(), *img, image.Pt(0, 0), draw.Src)

	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("wrong stride")
	}

	t.Width  = rgba.Rect.Size().X
	t.Height = rgba.Rect.Size().Y

	dataPtr := gl.Ptr(rgba.Pix)

	t.Activate(gl.TEXTURE0)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, int32(t.Width), int32(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, dataPtr)

	t.Deactivate()
}

func (t *Texture) LoadBytes_RGBA(width, height int, bytes []uint8) {

	t.Width  = width
	t.Height = height

	dataPtr := gl.Ptr(bytes)

	t.Activate(gl.TEXTURE0)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, int32(t.Width), int32(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, dataPtr)

	t.Deactivate()
}

func (t *Texture) Allocate(width, height int) {

	t.Width  = width
	t.Height = height

	t.Activate(gl.TEXTURE0)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, int32(t.Width), int32(t.Height), 0, gl.RGBA, gl.UNSIGNED_BYTE, unsafe.Pointer(nil))

	t.Deactivate()
}

func (t *Texture) Activate(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(gl.TEXTURE_2D, t.TextureId)
	t.textureUnit = texUnit
}

func (t *Texture) Deactivate() {
	t.textureUnit = 0
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (t *Texture) Free() {
	t.Deactivate()
	gl.DeleteTextures(1, &t.TextureId);
	t.TextureId = 0
}

func loadRGBA(filename string) (*image.RGBA, *image.Image) {
	imgFile, err := os.Open(filename)

	if err != nil {
		panic(err.Error())
	}

	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)

	if err != nil {
		panic(err.Error())
	}

	rgba := image.NewRGBA(img.Bounds())

	return rgba, &img
}

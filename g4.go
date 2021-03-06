package g4

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func Init() {
	gl.ClearColor(0.1, 0.4, 0.4, 1.0)

	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.CULL_FACE)

	gl.Enable(gl.FRAMEBUFFER_SRGB);

	// blending is required to be able to render text
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	g_colorRect = NewColorRect()
	g_textureRect = NewTextureRect("github.com/amortaza/go-g4/shader/texture.vertex.txt", "github.com/amortaza/go-g4/shader/texture.fragment.txt")
	g_stringRect = NewTextureRect("github.com/amortaza/go-g4/shader/font.vertex.txt", "github.com/amortaza/go-g4/shader/font.fragment.txt")
	g_canvasRect = NewTextureRect("github.com/amortaza/go-g4/shader/canvas.vertex.txt", "github.com/amortaza/go-g4/shader/canvas.fragment.txt")
}

func Clear(red,green,blue,alpha float32) {
	gl.ClearColor(red,green,blue,alpha)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func Uninit() {
	g_canvasRect.Free()
	g_stringRect.Free()
	g_textureRect.Free()
	g_colorRect.Free()
}

func PushView(width, height int) {
	PushViewport(width, height)
	PushOrtho(width,height)
}

func PopView() {
	PopViewport()
	PopOrtho()
}

func PushViewport(width, height int) {
	g_viewportWidthStack.Push(width)
	g_viewportHeightStack.Push(height)

	gl.Viewport(0, 0, int32(width), int32(height));
}

func PopViewport() {
	g_viewportWidthStack.Pop()
	g_viewportHeightStack.Pop()

	if g_viewportWidthStack.Size != 0 {
		width, _ := g_viewportWidthStack.Top().(int)
		height, _ := g_viewportHeightStack.Top().(int)

		gl.Viewport(0, 0, int32(width), int32(height));
	}
}

func PushOrtho(width, height int) {
	g_projection = mgl32.Ortho2D(0, float32(width), float32(height), 0)
	g_orthoStack.Push(g_projection)
}

func PopOrtho() {
	g_orthoStack.Pop()

	if g_orthoStack.Size != 0 {
		g_projection = g_orthoStack.Top().(mgl32.Mat4)
	}
}

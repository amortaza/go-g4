package g4

func DrawColorRect(	left int, top int,
			width int, height int,
			leftTopColor []float32,
			rightTopColor []float32,
			rightBottomColor []float32,
			leftBottomColor []float32) {

	g_colorRect.Draw(left, top, width, height, leftTopColor, rightTopColor, rightBottomColor, leftBottomColor, &g_projection[0])
}

func DrawTextureRect(	texture *Texture,
			left int32, top int32,
			width int32, height int32,
			rgba []float32,) {

	g_textureRect.Draw(texture, left, top, width, height, rgba, &g_projection[0])
}

func DrawStringRect(	fontTexture *StringTexture,
			left int32, top int32,
			rgb []float32,
			bg []float32,
			alpha float32) {

	g_stringRect.DrawString(fontTexture.Texture, left, top, fontTexture.Texture.Width, fontTexture.Texture.Height, rgb, bg, alpha, &g_projection[0])
}

func ClearRect(	width int32, height int32,
		red, green, blue float32 ) {

	g_colorRect.DrawSolid(0, 0, width, height, red, green, blue, &g_projection[0])
}

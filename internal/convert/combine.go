package convert

import (
	"image"
	"image-combinator/internal/input"
	"image/draw"
)

func Combine(imgs input.Images, screenWidth int, screenHeight int, padding int) *image.RGBA {
	// 背景画像の作成
	screen := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))

	// 背景画像への書き込み
	posX, posY := 0, 0
	for _, img := range imgs {
		rect := image.Rect(posX, posY, posX+img.Width, posY+img.Height)
		draw.Draw(screen, rect, img.Src, image.Point{0, 0}, draw.Over)
		posX += img.Width
	}

	return screen
}

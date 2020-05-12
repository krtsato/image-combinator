package convert

import (
	"image"
	"image-combinator/internal/input"
	"image/draw"
)

func Combine(imgs input.Images, options input.CliOptions) *image.RGBA {
	// 背景画像の作成
	platform := options.Platform
	usecase := options.Usecase
	usecaseArr := input.PlatformMap[platform][usecase]
	screen := image.NewRGBA(image.Rect(0, 0, usecaseArr["width"], usecaseArr["height"]))

	// 背景画像への書き込み
	posX, posY := 0, 0
	for _, img := range imgs {
		rect := image.Rect(posX, posY, posX+img.Width, posY+img.Height)
		draw.Draw(screen, rect, img.Src, image.Point{0, 0}, draw.Over)
		posX += img.Width
	}

	return screen
}

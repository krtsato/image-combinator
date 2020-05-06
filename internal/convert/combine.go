package convert

import (
	"image"
	"image-combinator/internal/calc"
	"image-combinator/internal/input"
	"image/draw"
)

func Combine(imgs input.Images) *image.RGBA {
	// 背景画像の作成
	outImgWidth := calc.SumWidth(imgs)
	outImgHeight := calc.MaxHeight(imgs)
	outImg := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))

	// 背景画像への書き込み
	posX, posY := 0, 0
	for _, img := range imgs {
		rect := image.Rect(posX, posY, posX + img.Width, posY + img.Height)
		draw.Draw(outImg, rect, img.Src, image.Point{0, 0}, draw.Over)
		posX += img.Width
	}

	return outImg
}

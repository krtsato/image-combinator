package convert

import (
	"image"
	"image-combinator/internal/calc"
	"image-combinator/internal/input"
	"image/draw"
)

func Combine(imgs input.Images) *image.RGBA {
	// 背景画像の作成
	outImgWidth := calc.MaxWidth(imgs)
	outImgHeight := calc.SumHeight(imgs)
	outImg := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))

	// 背景画像への書き込み
	pos := 0
	for _, img := range imgs {
		rect := image.Rect(0, pos, img.Width, pos + img.Height)
		draw.Draw(outImg, rect, img.Src, image.Point{0, 0}, draw.Over)
		pos += img.Height
	}

	return outImg
}

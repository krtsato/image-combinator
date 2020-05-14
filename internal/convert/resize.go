package convert

import (
	"image"
	"image-combinator/internal/input"

	"golang.org/x/image/draw"
)

func ResizeImage(img *input.Image) {
	// 要 usecaseMap["width"], measureMap["column"] -> xPadding,
	// imgSideLen := (usecaseMap["width"] - (measureMap["column"]+1)*xPadding) / measureMap["column"]
	imgSideLen := 200
	scaledImg := image.NewRGBA(image.Rect(0, 0, imgSideLen, imgSideLen))
	draw.CatmullRom.Scale(scaledImg, scaledImg.Bounds(), img.Src, img.Src.Bounds(), draw.Over, nil)
	img.Src = scaledImg
}

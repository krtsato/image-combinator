package convert

import (
	"image"
	"image-combinator/internal/input"

	"golang.org/x/image/draw"
)

func ResizeImage(img *input.Image, sideLen int) {
	scaledImg := image.NewRGBA(image.Rect(0, 0, sideLen, sideLen))
	draw.CatmullRom.Scale(scaledImg, scaledImg.Bounds(), img.Src, img.Src.Bounds(), draw.Over, nil)
	img.Src = scaledImg
}

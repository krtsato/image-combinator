package input

import (
	"image"
	"os"
)

type Image struct {
	Src    image.Image
	Width  int
	Height int
}

type Images []Image

// 入力画像のデータを取得する
func InitImage(path string) (*Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	imgBounds := img.Bounds()

	return &Image{img, imgBounds.Dx(), imgBounds.Dy()}, nil
}

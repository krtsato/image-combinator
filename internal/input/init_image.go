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

// 画像データの読み込み
func getImage(file *os.File) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// 画像サイズの読み込み
func getImageConfig(file *os.File) (image.Config, error) {
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return image.Config{}, err
	}

	return config, nil
}

// Image 構造体の初期化
func InitImage(path string) (Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return Image{}, err
	}
	defer file.Close()

	img, err := getImage(file)
	if err != nil {
		return Image{}, err
	}

	config, err := getImageConfig(file)
	if err != nil {
		return Image{}, err
	}

	return Image{img, config.Width, config.Height}, nil
}

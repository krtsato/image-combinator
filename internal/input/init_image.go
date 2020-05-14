package input

import (
	"bytes"
	"image"
	"image/jpeg"
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
func getImageConfig(file *bytes.Buffer) (*image.Config, error) {
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return &image.Config{}, err
	}

	return &config, nil
}

// Image 構造体の初期化
func InitImage(path string, options *CliOptions) (*Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return &Image{}, err
	}
	defer file.Close()

	img, err := getImage(file)
	if err != nil {
		return &Image{}, err
	}

	// config を取得するため buf を用意する
	// decode 後の file を使い回すと unknown format になる
	// 今後 image.Decode や img.Bounds().Dx()/Dy()
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return &Image{}, err
	}

	config, err := getImageConfig(buf)
	if err != nil {
		return &Image{}, err
	}

	/*
		platform := options.Platform
		usecase := options.Usecase
		aspectRatio := options.AspectRatio
		density := strconv.Itoa(options.Density)
		calc.ResizeImage()
	*/

	config.Width = 200
	config.Height = 200

	return &Image{img, config.Width, config.Height}, nil
}

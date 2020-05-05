package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type myImage struct {
	img    image.Image
	width  int
	height int
}

func getImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getImageConfig(path string) (image.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return image.Config{}, err
	}
	defer file.Close()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return image.Config{}, err
	}

	return config, nil
}

func getMyImage(path string) (myImage, error) {
	img, err := getImage(path)
	if err != nil {
		return myImage{}, err
	}

	config, err := getImageConfig(path)
	if err != nil {
		return myImage{}, err
	}

	return myImage{img, config.Width, config.Height}, nil
}

func getMaxWidth(imgs []myImage) int {
	var result int
	for _, img := range imgs {
		if result < img.width {
			result = img.width
		}
	}

	return result
}

func getSumHeight(imgs []myImage) int {
	var result int
	for _, img := range imgs {
		result += img.height
	}

	return result
}

func createConcatImage(imgs []myImage) *image.RGBA {
	outImgWidth := getMaxWidth(imgs)
	outImgHeight := getSumHeight(imgs)

	outImg := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))

	pos := 0
	for _, img := range imgs {
		rect := image.Rect(0, pos, img.width, pos+img.height)
		draw.Draw(outImg, rect, img.img, image.Point{0, 0}, draw.Over)
		pos += img.height
	}

	return outImg
}

func concatImages(paths []string) error {
	if len(paths) < 2 {
		return fmt.Errorf("Error: need more than 2 images")
	}

	var imgs []myImage
	for _, path := range paths {
		img, err := getMyImage(path)
		if err != nil {
			return err
		}

		imgs = append(imgs, img)
	}

	outImg := createConcatImage(imgs)

	out, err := os.Create("assets/output/out.jpg")
	if err != nil {
		return err
	}
	defer out.Close()

	jpegQuality := &jpeg.Options{Quality: 100}
	if err := jpeg.Encode(out, outImg, jpegQuality); err != nil {
		return err
	}

	return nil
}

func main() {
	// input images file path
	paths := []string{"assets/input/test-img1.jpg", "assets/input/test-img2.jpg"}
	if err := concatImages(paths); err != nil {
		log.Fatalln(err)
	}
}

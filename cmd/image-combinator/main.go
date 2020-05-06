package main

import (
	"fmt"
	"image"
	"image-combinator/internal/calc"
	"image-combinator/internal/input"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func createConcatImage(imgs input.Images) *image.RGBA {
	outImgWidth := calc.GetMaxWidth(imgs)
	outImgHeight := calc.GetSumHeight(imgs)

	outImg := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))

	pos := 0
	for _, img := range imgs {
		rect := image.Rect(0, pos, img.Width, pos+img.Height)
		draw.Draw(outImg, rect, img.Src, image.Point{0, 0}, draw.Over)
		pos += img.Height
	}

	return outImg
}

func concatImages(paths []string) error {
	if len(paths) < 2 {
		return fmt.Errorf("Error: need more than 2 images")
	}

	var imgs input.Images
	for _, path := range paths {
		img, err := input.InitImage(path)
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

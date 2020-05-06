package main

import (
	"fmt"
	"image-combinator/internal/convert"
	"image-combinator/internal/input"
	"image-combinator/internal/output"
	"log"
)

func integrateImages(paths []string) error {
	// 入力画像は２枚以上
	if len(paths) < 2 {
		return fmt.Errorf("Error: need two or more images")
	}

	// 全入力画像を配列に格納
	var imgs input.Images
	for _, path := range paths {
		img, err := input.InitImage(path)
		if err != nil {
			return err
		}

		imgs = append(imgs, img)
	}

	// 加工
	outImg := convert.Combine(imgs)

	// 出力
	err := output.Save(outImg)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	paths := input.GetPaths()
	err := integrateImages(paths)
	if err != nil {
		log.Fatalln(err)
	}
}

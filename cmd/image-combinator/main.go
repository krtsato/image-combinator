package main

import (
	"fmt"
	"image-combinator/internal/calc"
	"image-combinator/internal/convert"
	"image-combinator/internal/input"
	"image-combinator/internal/output"
	"log"
)

func integrateImages() error {
	// コマンドオプションを読み込む
	cliOptions, err := input.InitCliOptions()
	if err != nil {
		return err
	}
	density := cliOptions.Density

	// 全入力画像のパスを取得
	paths, err := input.GetPaths("assets/input/*.jpg")
	if err != nil {
		return err
	}

	// 画像の不足分を取得する
	outputQuant, addition := calc.GetOutputQuant(len(paths), density)
	fmt.Println(outputQuant, addition)

	// 全入力画像の情報を格納
	var imgs input.Images
	for _, path := range paths {
		img, err := input.InitImage(path)
		if err != nil {
			return err
		}

		imgs = append(imgs, img)
	}

	// 加工
	screen := convert.Combine(imgs, cliOptions)

	// 出力
	if err := output.Save(screen); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := integrateImages(); err != nil {
		log.Fatalln(err)
	}
}

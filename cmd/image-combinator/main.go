package main

import (
	"image-combinator/internal/convert"
	"image-combinator/internal/input"
	"image-combinator/internal/output"
	"log"
)

func integrateImages() error {
	// 全入力画像のパスを取得
	paths, err := input.GetPaths("assets/input/*.jpg")
	if err != nil {
		return err
	}

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
	screen := convert.Combine(imgs)

	// 出力
	if err := output.Save(screen); err != nil {
		return err
	}

	return nil
}

func main() {
	// コマンドオプションを読み込む
	input.CliFlagParse()

	// 画像処理
	if err := integrateImages(); err != nil {
		log.Fatalln(err)
	}
}

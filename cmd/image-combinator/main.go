package main

import (
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

	// 全入力画像のパス・出力画像枚数を取得する
	paths, outputQuant, err := input.GetPaths("assets/input/*.jpg", density)
	if err != nil {
		return err
	}

	entryIndex := 0
	for outputQuant > 0 {
		var imgs input.Images
		entryPaths := paths[entryIndex:density]

		// 入力画像の情報を格納
		for _, path := range entryPaths {
			img, err := input.InitImage(path, cliOptions)
			if err != nil {
				return err
			}

			imgs = append(imgs, *img)
			entryIndex++
		}
		// リサイズ

		// 加工
		screen := convert.Combine(imgs, cliOptions)

		// 出力
		if err := output.Save(screen); err != nil {
			return err
		}

		// 出力画像の残り枚数を減らす
		outputQuant--
	}

	return nil
}

func main() {
	if err := integrateImages(); err != nil {
		log.Fatalln(err)
	}
}

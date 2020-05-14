package main

import (
	"image-combinator/internal/calc"
	"image-combinator/internal/convert"
	"image-combinator/internal/input"
	"image-combinator/internal/output"
	"log"
	"strconv"
)

func integrateImages() error {
	// コマンドオプションを読み込む
	cliOptions, err := input.InitCliOptions()
	if err != nil {
		return err
	}
	density := cliOptions.Density
	densityCol := input.AspectMap[cliOptions.AspectRatio][strconv.Itoa(density)]["column"]
	screenWidth := input.PlatformMap[cliOptions.Platform][cliOptions.Usecase]["width"]
	screenHeight := input.PlatformMap[cliOptions.Platform][cliOptions.Usecase]["Height"]

	// 全入力画像のパス・出力画像枚数を取得する
	paths, outputQuant, err := input.GetPaths("assets/input/*.jpg", density)
	if err != nil {
		return err
	}

	// 構成画像のサイズと余白を取得する
	sideLen, padding := calc.MaterialSize(screenWidth, densityCol)

	entryIndex := 0
	for outputQuant > 0 {
		var imgs input.Images
		entryPaths := paths[entryIndex:density]

		// 入力画像の情報を格納
		for _, path := range entryPaths {
			img, err := input.InitImage(path)
			if err != nil {
				return err
			}

			// リサイズ
			convert.ResizeImage(img, sideLen)

			imgs = append(imgs, *img)
			entryIndex++
		}

		// 加工
		screen := convert.Combine(imgs, screenWidth, screenHeight, padding)

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

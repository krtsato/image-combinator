package main

import (
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
	measureMap := input.AspectMap[cliOptions.AspectRatio][density]
	densityCol := measureMap["column"]
	densityRow := measureMap["row"]
	screenMap := input.PlatformMap[cliOptions.Platform][cliOptions.Usecase]
	screenW := screenMap["width"]
	screenH := screenMap["height"]

	// 全入力画像のパス・出力画像枚数を取得する
	paths, outputQuant, err := input.GetPaths("assets/input/*.jpg", density)
	if err != nil {
		return err
	}

	// 入力画像１枚あたりのサイズ・余白を取得する
	sideLen, paddingX, paddingY := calc.MaterialSize(screenW, screenH, densityCol, densityRow)

	entryIndex := 0
	for i := 0; i < outputQuant; i++ {
		// 出力画像１枚を構成する入力画像を取得する
		entryPaths := paths[entryIndex : entryIndex+density]

		// 入力画像の情報を格納する
		var imgs input.Images
		for _, path := range entryPaths {
			img, err := input.InitImage(path)
			if err != nil {
				return err
			}

			// 計算結果をもとにリサイズする
			convert.ResizeImage(img, sideLen)

			imgs = append(imgs, *img)
			entryIndex++
		}

		// 画像を連結させる
		screen := convert.Combine(imgs, screenW, screenH, paddingX, paddingY, densityCol, densityRow)

		// 保存する
		if err := output.Save(screen, i+1); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := integrateImages(); err != nil {
		log.Fatalln(err)
	}
}

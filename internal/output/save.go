package output

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func Save(screen *image.RGBA, quantity int) error {
	// 出力パスを指定する
	fileName := fmt.Sprintf("output-%d.jpg", quantity)
	path := fmt.Sprintf("assets/output/%s", fileName)

	// 出力ファイルを生成する
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 出力画像を書き込む
	jpegQuality := &jpeg.Options{Quality: 100}
	if err := jpeg.Encode(file, screen, jpegQuality); err != nil {
		return err
	}

	return nil
}

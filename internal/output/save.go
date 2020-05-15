package output

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// 拡張子に応じてエンコードする
func encode(file io.Writer, screen *image.RGBA, extension string) error {
	switch extension {
	case "jpg":
		jpegQuality := &jpeg.Options{Quality: 100}
		if err := jpeg.Encode(file, screen, jpegQuality); err != nil {
			return err
		}
		return nil
	case "png":
		if err := png.Encode(file, screen); err != nil {
			return err
		}
		return nil
	case "gif":
		if err := gif.Encode(file, screen, nil); err != nil {
			return err
		}
		return nil
	default:
		return errors.New("Error: output file extension is wrong")
	}
}

func Save(screen *image.RGBA, quantity int) error {
	// 出力パスを指定する
	extension := "png"
	fileName := fmt.Sprintf("output-%d.%s", quantity, extension)
	path := fmt.Sprintf("assets/output/%s", fileName)

	// 出力ファイルを生成する
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 出力画像を書き込む
	if err := encode(file, screen, extension); err != nil {
		return err
	}

	return nil
}

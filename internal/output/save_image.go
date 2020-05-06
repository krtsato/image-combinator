package output

import (
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"time"
)

func SaveImage(outImg *image.RGBA) error { 
	// 出力パスの指定
	unixTime := strconv.FormatInt(time.Now().Unix(), 10)
	extension := ".jpg"
	fileName := unixTime + extension
	path := "assets/output/" + fileName

	// 出力画像の生成
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jpegQuality := &jpeg.Options{Quality: 100}
	if err := jpeg.Encode(file, outImg, jpegQuality); err != nil {
		return err
	}

	return nil
}

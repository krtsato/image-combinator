package input

import (
	"image-combinator/internal/calc"
	"path/filepath"
)

// 全入力画像のパス・出力画像枚数を取得する
func GetPaths(dir string, density int) ([]string, int, error) {
	// フォルダ内に用意した入力画像パスを取得する
	files, err := filepath.Glob(dir)
	if err != nil {
		return nil, 0, err
	}
	filesQuant := len(files)

	var paths []string
	outputQuant, addition := calc.GetOutputQuant(filesQuant, density)
	paths = append(paths, files...)

	// 入力画像の不足分を取得する
	for addition > 0 {
		paths = append(paths, "assets/input/default/soundtrackhub-icon.jpg")
		addition--
	}

	return paths, outputQuant, nil
}

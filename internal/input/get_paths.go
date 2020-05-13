package input

import (
	"image-combinator/internal/calc"
	"path/filepath"
)

func GetPaths(dir string, density int) ([]string, int, error) {
	// 入力画像のパスを取得する
	files, err := filepath.Glob(dir)
	if err != nil {
		return nil, 0, err
	}
	filesQuant := len(files)

	// 入力画像の不足分を取得する
	outputQuant, addition := calc.GetOutputQuant(filesQuant, density)
	paths := make([]string, filesQuant+addition)

	paths = append(paths, files...)
	for addition > 0 {
		paths = append(paths, "assets/input/default/soundtrackhub-icon.jpg")
		addition--
	}

	return paths, outputQuant, nil
}

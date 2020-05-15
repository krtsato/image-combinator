package input

import (
	"image-combinator/internal/calc"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

// .DS_Store などの不要ファイルを削除する
func removeNeedlessFile(file string) error {
	if err := os.Remove(file); err != nil {
		if os.IsNotExist(err) {
			return nil // 存在しない場合は何も返さない
		}
		return err
	}
	return nil
}

// 全入力画像のパス・出力画像枚数を取得する
func GetPaths(dir string, density int) ([]string, int, error) {
	// .DS_Store を削除する
	if err := removeNeedlessFile(filepath.Join(dir, ".DS_Store")); err != nil {
		return nil, 0, err
	}

	// フォルダ内に用意した入力画像パスを取得する
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}

	var filesQuant int
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(dir, file.Name())
		paths = append(paths, path)
		filesQuant++
	}

	// 入力画像の不足分を取得する
	outputQuant, addition := calc.GetOutputQuant(filesQuant, density)
	for addition > 0 {
		path := filepath.Join(dir, "default/soundtrackhub-icon.jpg")
		paths = append(paths, path)
		addition--
	}

	return paths, outputQuant, nil
}

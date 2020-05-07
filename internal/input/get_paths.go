package input

import (
	"fmt"
	"path/filepath"
)

func GetPaths(dir string) ([]string, error) {
	var paths []string

	files, err := filepath.Glob(dir)
	if err != nil {
		return paths, err
	}

	paths = append(paths, files...)
	if len(paths) < 2 {
		return paths, fmt.Errorf("Error: need two or more images")
	}
	return paths, nil
}

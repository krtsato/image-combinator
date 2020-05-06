package calc

import (
	"image-combinator/internal/input"
)

func GetMaxWidth(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		if result < img.Width {
			result = img.Width
		}
	}

	return result
}

func GetSumHeight(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		result += img.Height
	}

	return result
}

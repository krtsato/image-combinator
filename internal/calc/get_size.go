package calc

import (
	"image-combinator/internal/input"
)

func MaxWidth(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		if result < img.Width {
			result = img.Width
		}
	}

	return result
}

func SumHeight(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		result += img.Height
	}

	return result
}

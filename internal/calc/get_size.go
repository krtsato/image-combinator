package calc

import "fmt"

func calcGcd(w int, h int) int {
	if h == 0 {
		return w
	}

	return calcGcd(h, w%h)
}

func AspectRatio(width int, height int) string {
	gcd := calcGcd(width, height)
	wRatio := width / gcd
	hRatio := height / gcd
	aspectRatio := fmt.Sprintf("%d:%d", wRatio, hRatio)
	return aspectRatio
}

/*
func MaxWidth(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		if result < img.Width {
			result = img.Width
		}
	}

	return result
}

func MaxHeight(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		if result < img.Height {
			result = img.Height
		}
	}

	return result
}
func SumWidth(imgs input.Images) int {
	var result int
	for _, img := range imgs {
		result += img.Width
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
*/

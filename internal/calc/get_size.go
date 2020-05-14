package calc

import (
	"fmt"
)

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

func MaterialSize(screenSize int, densityCol int) (int, int) {
	padding := screenSize % densityCol
	paddingSum := padding * (densityCol + 1)
	sideLen := (screenSize - paddingSum) / densityCol
	return sideLen, padding
}

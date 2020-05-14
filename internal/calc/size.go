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

func MaterialSize(screenW, screenH, densityCol, densityRow int) (int, int, int) {
	paddingX := screenW % densityCol
	paddingSum := paddingX * (densityCol + 1)
	sideLen := (screenW - paddingSum) / densityCol
	paddingSum = screenH % (sideLen * densityRow)
	paddingY := 0
	if densityRow == 2 {
		paddingY = paddingSum / (densityRow + 2)
	} else {
		paddingY = paddingSum / (densityRow + 1)
	}

	return sideLen, paddingX, paddingY
}

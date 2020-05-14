package calc

import (
	"fmt"
)

// 最大公約数を計算する
func calcGcd(w int, h int) int {
	if h == 0 {
		return w
	}

	return calcGcd(h, w%h)
}

// CLI オプションから出力画像のアスペクト比を計算する
func AspectRatio(width int, height int) string {
	gcd := calcGcd(width, height)
	wRatio := width / gcd
	hRatio := height / gcd
	aspectRatio := fmt.Sprintf("%d:%d", wRatio, hRatio)
	return aspectRatio
}

// 入力画像１枚における１辺の長さ・X 軸方向の余白・Y 軸方向の余白を計算する
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

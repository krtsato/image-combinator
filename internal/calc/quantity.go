package calc

// 出力画像の合計枚数・入力画像の不足枚数を計算する
func GetOutputQuant(priorQuant int, density int) (int, int) {
	outputQuant := 1
	for density < priorQuant {
		priorQuant -= density
		outputQuant++
	}
	addition := density - priorQuant

	return outputQuant, addition
}

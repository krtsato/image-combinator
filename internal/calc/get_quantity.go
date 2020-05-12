package calc

func GetOutputQuant(preparedImg int, density int) (int, int) {
	outputQuant := 1
	for density < preparedImg {
		preparedImg -= density
		outputQuant++
	}
	addition := density - preparedImg

	return outputQuant, addition
}

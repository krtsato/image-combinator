package calc

func GetOutputQuant(priorQuant int, density int) (int, int) {
	outputQuant := 1
	for density < priorQuant {
		priorQuant -= density
		outputQuant++
	}
	addition := density - priorQuant

	return outputQuant, addition
}

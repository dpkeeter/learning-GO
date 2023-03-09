package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodCount := 0
	sauceCount := 0.0
	for i := range layers {
		if layers[i] == "noodles" {
			noodCount++
		}
		if layers[i] == "sauce" {
			sauceCount++
		}
	}
	noodles = 50 * noodCount
	sauce = sauceCount * .2
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fList []string, mList []string) {
	index := len(fList) - 1
	index2 := len(mList) - 1
	mList[index2] = fList[index]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numPortions int) []float64 {
	var newQuant = make([]float64, len(amounts))
	for i := 0; i < len(amounts); i++ {
		newQuant[i] = amounts[i] * float64(numPortions) / 2

	}
	return newQuant

}

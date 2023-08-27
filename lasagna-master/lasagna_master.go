package lasagna

// PreparationTime calculates preperation time
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer <= 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities calculates the amount of noodles and sauce needed based on the layers
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient take the secret ingredient from the friend's recipe and add it to your recipe
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales a recipe based on the amounts needed for 2 portions, and the portions wanted
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var needs []float64
	for _, v := range quantities {
		needs = append(needs, v*float64(portions)/2)
	}

	return needs
}

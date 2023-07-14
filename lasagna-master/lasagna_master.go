package lasagna1

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time > 0 {
		return time * len(layers)
	}
	return 2 * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(otherList []string, myList []string) {
	myList[len(myList)-1] = otherList[len(otherList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := []float64{}
	koef := float64(portions) / 2

	for _, v := range quantities {
		scaled = append(scaled, koef*v)
	}

	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

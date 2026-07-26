package lasagnamaster

// PreparationTime estimates the preparation time in minutes based on the number of layers.
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities calculates the grams of noodles and liters of sauce needed.
func Quantities(layers []string) (int, float64) {
	cnt_sauce := 0
	cnt_noodles := 0
    
	for _, layer := range layers {
		if layer == "sauce" {
			cnt_sauce++
		} else if layer == "noodles" {
			cnt_noodles++
		}
	}    
	return cnt_noodles * 50, float64(cnt_sauce) * 0.2
}

// AddSecretIngredient replaces the last element of your recipe list 
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe calculates the amounts for the desired number of portions.
// The quantities passed in are always for 2 portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
    
	for i, amount := range quantities {
		scaled[i] = amount * (float64(portions) / 2.0)
	}
	return scaled
}
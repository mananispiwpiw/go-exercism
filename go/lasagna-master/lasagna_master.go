package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	var totalTime int
	if time == 0 {
		return len(layers) * 2
	} else {
		totalTime = len(layers) * time
	}
	return totalTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	var totalNoodles int
	var totalSauce float64
	for _, value := range layers {
		if value == "noodles" {
			totalNoodles += 50
		} else if value == "sauce" {
			totalSauce += 0.2
		}
	}
	return totalNoodles, totalSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) []string {
	myList[len(myList)-1] = friendList[len(friendList)-1]
	return myList

}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	//var scaleQuantities = []float64{}
	var result = []float64{}
	totalPortions := float64(portions) / 2.0

	for i := 0; i < len(quantities); i++ {
		dummyValue := quantities[i] * totalPortions
		result = append(result, dummyValue)
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

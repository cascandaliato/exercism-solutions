package lasagna

const (
	defaultAvgPrepTime = 2
	noodlesPerLayer    = 50
	saucePerLayer      = 0.2
	servingsPerRecipe  = 2
)

func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = defaultAvgPrepTime
	}
	return len(layers) * avgPrepTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += noodlesPerLayer
		case "sauce":
			sauce += saucePerLayer
		}
	}
	return
}

func AddSecretIngredient(friendIngredients, ownIngredients []string) {
	ownIngredients[len(ownIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

func ScaleRecipe(recipe []float64, portions int) []float64 {
	quantities := make([]float64, len(recipe))
	factor := float64(portions) / servingsPerRecipe
	for i, q := range recipe {
		quantities[i] = q * factor
	}
	return quantities
}

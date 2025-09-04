package lasagna

func PreparationTime(layers []string, time int) int {
    if 1 > time {
        time = 2
    }

    return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
    noodlesPerLayer := 50
    saucePerLayer := 0.2
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += noodlesPerLayer
        }
        if layer == "sauce" {
            sauce += saucePerLayer
        }
    }
    return
}

func AddSecretIngredient(friendList []string, myList []string) {
    fIndex := len(friendList) - 1
    mIndex := len(myList) - 1
    myList[mIndex] = friendList[fIndex]
}

func ScaleRecipe(quantities []float64, scale int) (scaledQuantities []float64) {
    p := 2.0
    scaledQuantities = make([]float64, len(quantities))
    for i, quantity := range quantities {
        scaledQuantities[i] = quantity * float64(scale) / p
    }
    return
}

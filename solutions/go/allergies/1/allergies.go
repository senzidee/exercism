package allergies

var allergens = []string{
    "eggs", "peanuts", "shellfish", "strawberries",
    "tomatoes", "chocolate", "pollen", "cats",
}

func Allergies(allergies uint) []string {
	var result []string
    for i, name := range allergens {
        if allergies&(1<<i) != 0 {
            result = append(result, name)
        }
    }

    return result
}

func AllergicTo(allergies uint, allergen string) bool {
    for _, name := range Allergies(allergies) {
        if name == allergen {
            return true
        }
    }

    return false
}

package allergies


var allergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}


func Allergies(score uint) []string{
	allergies := []string{}

	for i , allergen := range allergens {

		if 0 != ( score & (1 << uint(i)) ) {
			allergies = append(allergies, allergen)
		}
	}

	return allergies
}

func AllergicTo(score uint, allergen string) bool {

	i := 0
	for i < len(allergens) && allergens[i] != allergen{
		i++
	}

	return 0 != ( score & (1 << uint(i)) )
}
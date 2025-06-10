package exercise

var itemAllergic = map[string]uint{
	"eggs": 1, "peanuts": 2, "shellfish": 4, "strawberries": 8,
	"tomatoes": 16, "chocolate": 32, "pollen": 64, "cats": 128,
}

func Allergies(allergies uint) []string {
	var allergyToList []string
	for allergen, value := range itemAllergic {
		if value&allergies != 0 {
			allergyToList = append(allergyToList, allergen)
		}
	}
	return allergyToList
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&itemAllergic[allergen] != 0
}

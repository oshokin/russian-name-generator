package russian_name_generator

import "math/rand"

// Patronymic generates a patronymic name (middle name) based on the given gender and excludeRareNames flag.
// If gender is Any, it randomly selects male or female gender.
// Rare names are excluded if excludeRareNames is true.
func Patronymic(isFeminine bool, excludeRareNames bool) string {
	return patronymic(globalFaker.Rand, isFeminine, excludeRareNames)
}

// Patronymic generates a patronymic name (middle name) based on the given gender and excludeRareNames flag.
// If gender is Any, it randomly selects male or female gender.
// Rare names are excluded if excludeRareNames is true.
func (f *Faker) Patronymic(isFeminine bool, excludeRareNames bool) string {
	return patronymic(f.Rand, isFeminine, excludeRareNames)
}

func patronymic(r *rand.Rand, isFeminine bool, excludeRareNames bool) string {
	dataSetNames := []string{"male_name"}
	if !excludeRareNames {
		dataSetNames = append(dataSetNames, "rare_male_name")
	}

	name := getRandValue(r, dataSetNames)

	return getPatronymicFromName(name, isFeminine)
}

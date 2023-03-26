package russian_name_generator

import "math/rand"

// Patronymic generates a random patronymic name based on the specified parameters.
// If isFeminine is true, the generated name will be for a female name, otherwise for a male name.
// If excludeRareNames is true, rare names will be excluded when choosing the base name for the patronymic.
func Patronymic(isFeminine bool, excludeRareNames bool) string {
	return patronymic(globalFaker.Rand, isFeminine, excludeRareNames)
}

// Patronymic generates a random patronymic name based on the specified parameters.
// If isFeminine is true, the generated name will be for a female name, otherwise for a male name.
// If excludeRareNames is true, rare names will be excluded when choosing the base name for the patronymic.
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

package russian_name_generator

import "math/rand"

// Patronymic generates a patronymic name (middle name) based on the given gender and excludeRareNames flag.
// If gender is Any, it randomly selects male or female gender.
// Rare names are excluded if excludeRareNames is true.
func Patronymic(gender Gender, excludeRareNames bool) string {
	return patronymic(globalFaker.Rand, gender, excludeRareNames)
}

// Patronymic generates a patronymic name (middle name) based on the given gender and excludeRareNames flag.
// If gender is Any, it randomly selects male or female gender.
// Rare names are excluded if excludeRareNames is true.
func (f *Faker) Patronymic(gender Gender, excludeRareNames bool) string {
	return patronymic(f.Rand, gender, excludeRareNames)
}

func patronymic(r *rand.Rand, gender Gender, excludeRareNames bool) string {
	dataSetNames := make([]string, 0, 2)
	if gender == GenderAny {
		gender = GenderMale
		if r.Intn(2) == 1 {
			gender = GenderFemale
		}
	}

	switch gender {
	case GenderMale:
		dataSetNames = append(dataSetNames, "male_name")
	case GenderFemale:
		dataSetNames = append(dataSetNames, "female_name")
	}

	if !excludeRareNames {
		switch gender {
		case GenderMale:
			dataSetNames = append(dataSetNames, "rare_male_name")
		case GenderFemale:
			dataSetNames = append(dataSetNames, "rare_female_name")
		}
	}

	name := getRandValue(r, dataSetNames)

	return getPatronymicFromName(name, gender == GenderFemale)
}

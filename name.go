package russian_name_generator

import "math/rand"

// Name returns a random name based on the given gender and whether or not to exclude rare names.
func Name(gender Gender, excludeRareNames bool) string {
	return name(globalFaker.Rand, gender, excludeRareNames)
}

// Name returns a random name based on the given gender and whether or not to exclude rare names.
func (f *Faker) Name(gender Gender, excludeRareNames bool) string {
	return name(f.Rand, gender, excludeRareNames)
}

func name(r *rand.Rand, gender Gender, excludeRareNames bool) string {
	dataSetNames := make([]string, 0, 4)
	switch gender {
	case GenderMale:
		dataSetNames = append(dataSetNames, "male_name")
	case GenderFemale:
		dataSetNames = append(dataSetNames, "female_name")
	default:
		dataSetNames = append(dataSetNames, "male_name", "female_name")
	}

	if !excludeRareNames {
		switch gender {
		case GenderMale:
			dataSetNames = append(dataSetNames, "rare_male_name")
		case GenderFemale:
			dataSetNames = append(dataSetNames, "rare_female_name")
		default:
			dataSetNames = append(dataSetNames, "rare_male_name", "rare_female_name")
		}
	}

	return getRandValue(r, dataSetNames)
}

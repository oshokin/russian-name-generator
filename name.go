package russian_name_generator

import "math/rand"

// Name returns a random name based on the given gender and whether or not to exclude rare names.
func Name(g GenderType, excludeRareNames bool) string {
	return name(globalFaker.Rand, g, excludeRareNames)
}

// Name returns a random name based on the given gender and whether or not to exclude rare names.
func (f *Faker) Name(g GenderType, excludeRareNames bool) string {
	return name(f.Rand, g, excludeRareNames)
}

func name(r *rand.Rand, g GenderType, excludeRareNames bool) string {
	const maxNumberOfSets = 4

	if !g.IsDefined() {
		g = gender(r)
	}

	var dataSetNames = make([]string, 0, maxNumberOfSets)

	switch g {
	case GenderMale:
		dataSetNames = append(dataSetNames, "male_name")
	case GenderFemale:
		dataSetNames = append(dataSetNames, "female_name")
	case GenderAny:
		fallthrough
	default:
		dataSetNames = append(dataSetNames, "male_name", "female_name")
	}

	if !excludeRareNames {
		switch g {
		case GenderMale:
			dataSetNames = append(dataSetNames, "rare_male_name")
		case GenderFemale:
			dataSetNames = append(dataSetNames, "rare_female_name")
		case GenderAny:
			fallthrough
		default:
			dataSetNames = append(dataSetNames, "rare_male_name", "rare_female_name")
		}
	}

	return getRandValue(r, dataSetNames)
}

package russian_name_generator

import "math/rand"

// Surname generates a random surname according to the specified gender.
// If gender is Any, a surname can be generated for any gender.
func Surname(g GenderType) string {
	return surname(globalFaker.Rand, g)
}

// Surname generates a random surname according to the specified gender.
// If gender is Any, a surname can be generated for any gender.
func (f *Faker) Surname(g GenderType) string {
	return surname(f.Rand, g)
}

func surname(r *rand.Rand, g GenderType) string {
	const maxNumberOfSets = 2

	if !g.IsDefined() {
		g = gender(r)
	}

	var dataSetNames = make([]string, 0, maxNumberOfSets)

	switch g {
	case GenderMale:
		dataSetNames = append(dataSetNames, "male_surname")
	case GenderFemale:
		dataSetNames = append(dataSetNames, "female_surname")
	case GenderAny:
		fallthrough
	default:
		dataSetNames = append(dataSetNames, "male_surname", "female_surname")
	}

	return getRandValue(r, dataSetNames)
}

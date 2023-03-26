package russian_name_generator

import "math/rand"

// Surname generates a random surname according to the specified gender.
// If gender is Any, a surname can be generated for any gender.
func Surname(gender Gender) string {
	return surname(globalFaker.Rand, gender)
}

// Surname generates a random surname according to the specified gender.
// If gender is Any, a surname can be generated for any gender.
func (f *Faker) Surname(gender Gender) string {
	return surname(f.Rand, gender)
}

func surname(r *rand.Rand, gender Gender) string {
	dataSetNames := make([]string, 0, 2)
	switch gender {
	case GenderMale:
		dataSetNames = append(dataSetNames, "male_surname")
	case GenderFemale:
		dataSetNames = append(dataSetNames, "female_surname")
	default:
		dataSetNames = append(dataSetNames, "male_surname", "female_surname")
	}

	return getRandValue(r, dataSetNames)
}

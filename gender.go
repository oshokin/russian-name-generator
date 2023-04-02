package russian_name_generator

import "math/rand"

// GenderType represents a person's gender.
type GenderType uint8

const (
	// GenderAny is any gender.
	GenderAny GenderType = iota
	// GenderMale is a male gender.
	GenderMale
	// GenderFemale is a female gender.
	GenderFemale
)

// IsDefined returns a boolean indicating whether the given gender is defined.
func (g GenderType) IsDefined() bool {
	return g == GenderMale || g == GenderFemale
}

// IsFeminine returns a boolean indicating whether the given gender is feminine.
func (g GenderType) IsFeminine() bool {
	return g == GenderFemale
}

// String returns the string representation of the Gender type.
func (g GenderType) String() string {
	switch g {
	case GenderAny:
		return "Any"
	case GenderMale:
		return "Male"
	case GenderFemale:
		return "Female"
	default:
		return "Unknown"
	}
}

// Gender returns a random gender.
func Gender() GenderType {
	return gender(globalFaker.Rand)
}

// Gender returns a random gender.
func (f *Faker) Gender() GenderType {
	return gender(f.Rand)
}

func gender(r *rand.Rand) GenderType {
	if getRandomBool(r) {
		return GenderMale
	}

	return GenderFemale
}

func getRandomBool(r *rand.Rand) bool {
	return randIntRange(r, 0, 1) == 1
}

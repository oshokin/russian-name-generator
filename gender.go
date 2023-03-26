package russian_name_generator

// Gender type represents a person's gender.
type Gender uint8

const (
	// GenderAny is any gender.
	GenderAny Gender = iota
	// GenderMale is a male gender.
	GenderMale
	// GenderFemale is a female gender.
	GenderFemale
)

// String returns the string representation of the Gender type.
func (g Gender) String() string {
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

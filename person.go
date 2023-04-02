package russian_name_generator

import "math/rand"

type (
	// PersonFields define the fields for generating a person's info.
	PersonFields struct {
		Name             bool       // Include person's name.
		Surname          bool       // Include person's surname.
		Patronymic       bool       // Include person's patronymic.
		Gender           GenderType // Person's gender.
		ExcludeRareNames bool       // Whether to exclude rare names.
	}

	// PersonInfo represents a generated person's info.
	PersonInfo struct {
		Name       string     // Person's name.
		Surname    string     // Person's surname.
		Patronymic string     // Person's patronymic.
		Gender     GenderType // Person's gender.
	}
)

var defaultFields = &PersonFields{
	Name:             true,
	Surname:          true,
	Patronymic:       true,
	Gender:           GenderAny,
	ExcludeRareNames: false,
}

// Person returns a random person's info based on the specified fields.
// If the provided fields argument is nil, the default fields will be used,
// including the person's name, surname, patronymic,
// and gender set to GenderAny, with rare names included.
func Person(fields *PersonFields) *PersonInfo {
	return person(globalFaker.Rand, fields)
}

// Person returns a random person's info based on the specified fields.
// If the provided fields argument is nil, the default fields will be used,
// including the person's name, surname, patronymic,
// and gender set to GenderAny, with rare names included.
func (f *Faker) Person(fields *PersonFields) *PersonInfo {
	return person(f.Rand, fields)
}

func person(r *rand.Rand, fields *PersonFields) *PersonInfo {
	if fields == nil {
		fields = defaultFields
	}

	g := fields.Gender
	if !g.IsDefined() {
		g = gender(r)
	}

	return &PersonInfo{
		Name:       name(r, g, fields.ExcludeRareNames),
		Surname:    surname(r, g),
		Patronymic: patronymic(r, g.IsFeminine(), fields.ExcludeRareNames),
		Gender:     g,
	}
}

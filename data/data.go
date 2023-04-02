package data

// data is a map of string keys to slices of strings containing common names and surnames.
var data = map[string][]string{
	"male_name":        MaleNames,
	"rare_male_name":   RareMaleNames,
	"female_name":      FemaleNames,
	"rare_female_name": RareFemaleNames,
	"male_surname":     MaleSurnames,
	"female_surname":   FemaleSurnames,
}

// GetList returns a slice of slices of strings containing the requested lists of names and surnames.
// If no valid list names are specified, it returns nil.
func GetList(name []string) [][]string {
	if len(name) == 0 {
		return nil
	}

	result := make([][]string, 0, len(name))

	for _, v := range name {
		l, ok := data[v]
		if !ok {
			continue
		}

		result = append(result, l)
	}

	return result
}

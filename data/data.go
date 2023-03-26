package data

var data = map[string][]string{
	"male_name":        MaleNames,
	"rare_male_name":   RareMaleNames,
	"female_name":      FemaleNames,
	"rare_female_name": RareFemaleNames,
	"male_surname":     MaleSurnames,
	"female_surname":   FemaleSurnames,
}

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

package russian_name_generator

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strings"
	"unicode"

	"github.com/oshokin/russian-name-generator/data"
)

// Seed will set the global random value. Setting seed to 0 will use crypto/rand
func Seed(seed int64) {
	if seed == 0 {
		binary.Read(crand.Reader, binary.BigEndian, &seed)
		globalFaker.Rand.Seed(seed)
	} else {
		globalFaker.Rand.Seed(seed)
	}
}

// Transliterate transliterates string from Russian to Latin letters
func Transliterate(text string) string {
	if len(text) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.Grow(len(text))

	for _, rc := range text {
		if lc, ok := translitTable[rc]; ok {
			sb.WriteString(lc)
			continue
		}

		sb.WriteRune(rc)
	}

	return sb.String()
}

func getPatronymicFromName(name string, isFeminine bool) string {
	if name == "" {
		return ""
	}

	name = strings.ToLower(name)
	if name == "пётр" {
		name = strings.Replace(name, "ё", "е", 1)
	}

	switch name {
	case "павел":
		name = "павлов"
	case "лев":
		name = "львов"
	case "яков":
		name = "яковлев"
	}

	var (
		runeName  = []rune(name)
		runeCount = len(runeName)
	)

	var (
		lastRune               = runeName[runeCount-1]
		lastLetter             = string(lastRune)
		penultimateLetter      = string(runeName[runeCount-2 : runeCount-1])
		lastTwoLetters         = string(runeName[runeCount-2:])
		lastThreeLetters       = string(runeName[runeCount-3:])
		allButLastLetter       = string(runeName[:runeCount-1])
		allButLastTwoLetters   = string(runeName[:runeCount-2])
		allButLastThreeLetters = string(runeName[:runeCount-3])
		suffix                 = "ич"
	)

	if isFeminine {
		endsWithAorYa := lastLetter == "а" || lastLetter == "я"
		endsWithYeya := lastTwoLetters == "ея" || lastTwoLetters == "ия"
		suffix = "на"

		if endsWithAorYa && !endsWithYeya {
			suffix = "ична"
		}
	}

	var (
		isLastRuneExceptionConsonant = strings.ContainsRune("жцчшщЖЦЧШЩ", lastRune)
		baseOfName                   string
		whoseSuffix                  string
	)

	switch {
	case lastLetter == "й" ||
		lastTwoLetters == "ея" ||
		lastTwoLetters == "ия":
		whoseSuffix = "ев"
		baseOfName = allButLastLetter
		letterBeforeEnding := runeName[runeCount-3]
		isConsonantBeforeEnding := isRussianConsonant(letterBeforeEnding)
		isVowelBeforeConsonant := isRussianVowel(runeName[runeCount-4])
		isNTBeforeEnding := string(runeName[runeCount-4:runeCount-2]) == "нт"
		isExceptionConsonant := strings.ContainsRune("кхц", letterBeforeEnding)

		if lastTwoLetters == "ий" &&
			((isConsonantBeforeEnding && isVowelBeforeConsonant && !isExceptionConsonant) ||
				isNTBeforeEnding) {
			whoseSuffix = "ьев"
			baseOfName = allButLastTwoLetters
		}
	case isRussianVowel(lastRune):
		var (
			indexOfA      = strings.IndexRune(name, 'а')
			isSpecialName = name == "фока" || name == "мина"
		)

		baseOfName = allButLastLetter

		switch {
		case (lastLetter == "а" || lastLetter == "я") &&
			getSyllablesCount(name) <= 2 &&
			!isSpecialName &&
			(indexOfA < 0 || indexOfA > runeCount-2):
			if isFeminine {
				whoseSuffix = "ин"
			}
		case lastLetter == "а" &&
			strings.ContainsAny(penultimateLetter, "лмн"):
			whoseSuffix = "ов"
			if isFeminine {
				suffix = "на"
			}
		case strings.ContainsAny(lastLetter, "ео"):
			whoseSuffix = "в"
			baseOfName += lastLetter
		case lastLetter == "и":
			whoseSuffix = "ев"
			baseOfName += lastLetter
		}
	case isLastRuneExceptionConsonant || lastLetter == "ь":
		whoseSuffix = "ев"
		baseOfName = allButLastLetter
		if isLastRuneExceptionConsonant {
			baseOfName = name
		}
	default:
		if lastTwoLetters != "ов" && lastTwoLetters != "ев" {
			whoseSuffix = "ов"
		}
		switch {
		case lastThreeLetters == "иил":
			baseOfName = allButLastTwoLetters + lastLetter
		case lastThreeLetters == "аил" || lastThreeLetters == "уил":
			baseOfName = allButLastThreeLetters + "о"
			if lastThreeLetters == "аил" {
				baseOfName = allButLastTwoLetters
			}
			baseOfName += "йл"
		default:
			baseOfName = name
		}
	}

	return capitalizeFirstLetter(
		strings.Join(
			[]string{baseOfName,
				whoseSuffix,
				suffix}, ""))
}

func getRandValue(r *rand.Rand, dataSetName []string) string {
	dataSets := data.GetList(dataSetName)
	if len(dataSets) == 0 {
		return ""
	}

	dataSet := dataSets[r.Intn(len(dataSets))]
	return dataSet[r.Intn(len(dataSet))]
}

func getSyllablesCount(name string) (count uint8) {
	for _, letter := range name {
		if !isRussianVowel(letter) {
			continue
		}

		count++
	}

	return
}

func isRussianVowel(symbol rune) bool {
	return strings.ContainsRune("аеиоуыэюяАЕИОУЫЭЮЯ", symbol)
}

func isRussianConsonant(symbol rune) bool {
	return strings.ContainsRune("бвгджзйклмнпрстфхцчшщБВГДЖЗЙКЛМНПРСТФХЦЧШЩs", symbol)
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}

	v := []rune(s)
	if !unicode.IsLower(v[0]) {
		return s
	}

	v[0] = unicode.ToUpper(v[0])

	return string(v)
}

package russian_name_generator

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strings"
	"unicode"

	"github.com/oshokin/russian-name-generator/data"
)

// Seed will set the global random value.
// Setting seed to 0 will use crypto/rand.
func Seed(seed int64) {
	if seed == 0 {
		_ = binary.Read(crand.Reader, binary.BigEndian, &seed)
		globalFaker.Rand.Seed(seed)
	} else {
		globalFaker.Rand.Seed(seed)
	}
}

// Transliterate transliterates string from Russian to Latin letters.
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
		lastRune               = safeRuneAt(runeName, runeCount-1)
		lastLetter             = string(lastRune)
		penultimateLetter      = string(safeRuneSlice(runeName, runeCount-2, runeCount-1))
		lastTwoLetters         = string(safeRuneSlice(runeName, runeCount-2, runeCount))
		lastThreeLetters       = string(safeRuneSlice(runeName, runeCount-3, runeCount))
		allButLastLetter       = string(safeRuneSlice(runeName, 0, runeCount-1))
		allButLastTwoLetters   = string(safeRuneSlice(runeName, 0, runeCount-2))
		allButLastThreeLetters = string(safeRuneSlice(runeName, 0, runeCount-3))
		suffix                 = "ич"
	)

	const (
		aLetter    = "а"
		yaLetter   = "я"
		ailSuffix  = "аил"
		evSuffix   = "ев"
		yeyaSuffix = "ея"
		iyaSuffix  = "ия"
		naSuffix   = "на"
		ovSuffix   = "ов"
	)

	if isFeminine {
		endsWithAorYa := lastLetter == aLetter || lastLetter == yaLetter
		endsWithYeya := lastTwoLetters == yeyaSuffix || lastTwoLetters == iyaSuffix
		suffix = naSuffix

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
		lastTwoLetters == yeyaSuffix ||
		lastTwoLetters == iyaSuffix:
		whoseSuffix = evSuffix
		baseOfName = allButLastLetter
		letterBeforeEnding := safeRuneAt(runeName, runeCount-3)
		isConsonantBeforeEnding := isRussianConsonant(letterBeforeEnding)
		isVowelBeforeConsonant := isRussianVowel(safeRuneAt(runeName, runeCount-4))
		isNTBeforeEnding := string(safeRuneSlice(runeName, runeCount-4, runeCount-2)) == "нт"
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
		case (lastLetter == aLetter || lastLetter == yaLetter) &&
			getSyllablesCount(name) <= 2 &&
			!isSpecialName &&
			(indexOfA < 0 || indexOfA > runeCount-2):
			if isFeminine {
				whoseSuffix = "ин"
			}
		case lastLetter == aLetter &&
			strings.ContainsAny(penultimateLetter, "лмн"):
			whoseSuffix = ovSuffix

			if isFeminine {
				suffix = naSuffix
			}
		case strings.ContainsAny(lastLetter, "ео"):
			whoseSuffix = "в"
			baseOfName += lastLetter
		case lastLetter == "и":
			whoseSuffix = evSuffix
			baseOfName += lastLetter
		}
	case isLastRuneExceptionConsonant || lastLetter == "ь":
		whoseSuffix = evSuffix
		baseOfName = allButLastLetter

		if isLastRuneExceptionConsonant {
			baseOfName = name
		}
	default:
		if lastTwoLetters != ovSuffix && lastTwoLetters != evSuffix {
			whoseSuffix = ovSuffix
		}

		switch {
		case lastThreeLetters == "иил":
			baseOfName = allButLastTwoLetters + lastLetter
		case lastThreeLetters == ailSuffix || lastThreeLetters == "уил":
			baseOfName = allButLastThreeLetters + "о"
			if lastThreeLetters == ailSuffix {
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

func safeRuneSlice(runes []rune, left, right int) []rune {
	if len(runes) == 0 || right < 0 {
		return nil
	}

	if left < 0 {
		left = 0
	}

	if right > len(runes) {
		right = len(runes)
	}

	return runes[left:right]
}

func safeRuneAt(runes []rune, index int) rune {
	if len(runes) == 0 {
		return 0
	}

	if index > len(runes)-1 {
		index = len(runes) - 1
	}

	if index < 0 {
		return 0
	}

	return runes[index]
}

func getRandValue(r *rand.Rand, dataSetName []string) string {
	dataSets := data.GetList(dataSetName)
	if len(dataSets) == 0 {
		return ""
	}

	dataSet := dataSets[r.Intn(len(dataSets))]

	return dataSet[r.Intn(len(dataSet))]
}

func randIntRange(r *rand.Rand, min, max int) int {
	if min == max {
		return min
	}

	if min > max {
		ogmin := min
		min = max
		max = ogmin
	}

	if max-min+1 > 0 {
		return min + int(r.Int63n(int64(max-min+1)))
	}

	for {
		v := int(r.Uint64())
		if (v >= min) && (v <= max) {
			return v
		}
	}
}

func getSyllablesCount(name string) uint8 {
	var count uint8

	for _, letter := range name {
		if !isRussianVowel(letter) {
			continue
		}

		count++
	}

	return count
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

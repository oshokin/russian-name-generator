# Russian Name Generator

`russian-name-generator` is a Go package for generating random Russian names, including first names, patronymics, and surnames.

## Inspiration

This project was inspired by the Go library [gofakeit](https://github.com/brianvoe/gofakeit), which provides similar functionality for generating random data across multiple categories.

## Data Sets

The project uses data sets of Russian names from the [russian-names](https://github.com/cybermatt/russian-names) repository. The patronymic grammatical rules are derived from the [linguistics_problems](https://github.com/roddar92/linguistics_problems) repository.

I would like to thank the authors of these repositories for sharing their work and making this project possible.

## Installation

You can install `russian-name-generator` using the following command:

```sh
go get github.com/oshokin/russian-name-generator
```

## Usage

### First, import the package:

```go
import rus_name_gen "github.com/oshokin/russian-name-generator"
```

### Then, you can generate a random name:

```go
name := rus_name_gen.Name(rus_name_gen.GenderAny, true)
```

This will return a random name for any gender, excluding rare names.

Parameters:

- `gender` (`Gender`): The gender for the generated name. Valid options are `GenderAny`, `GenderMale`, and `GenderFemale`.

- `excludeRareNames` (`bool`): Whether or not to exclude rare names from the dataset.

### You can also generate a random surname:

```go
surname := rus_name_gen.Surname(rus_name_gen.GenderMale)
```

This will return a random surname for a male.

Parameters:

- `gender` (`Gender`): The gender for the generated surname. Valid options are `GenderAny`, `GenderMale`, and `GenderFemale`.

### Finally, you can generate a random patronymic:

```go
patronymic := rus_name_gen.Patronymic(true, false)
```

This will return a random patronymic for a female name, including rare base names.

Parameters:

- `isFeminine` (`bool`): Whether the patronymic should be feminine (i.e., for a female name). If false, the method generates a masculine patronymic.

- `excludeRareNames` (`bool`): Whether to exclude rare names when choosing the base name for the patronymic.

### Also, you can transliterate a string from Russian to Latin letters:

```go
text := "Брат, братан, братишка, когда меня отпустит?"
transliteratedText := russian_name_generator.Transliterate(text)
```

This will return a transliterated string in Latin characters.

Parameters:

- `text` (`string`): The Russian text to be transliterated to Latin characters.

## Seed

If you are using the default global usage and dont care about seeding no need to set anything.
`russian-name-generator` will seed it with a cryptographically secure number.

If you need a reproducible outcome you can set it via the Seed function call. Every example in
this repo sets it for testing purposes.

```go
import rus_name_gen "github.com/oshokin/russian-name-generator"

rus_name_gen.Seed(0) // If 0 will use crypto/rand to generate a number

// or

rus_name_gen.Seed(14000088) // Set it to whatever number you want
```

## Random Sources

`russian-name-generator` has a few rand sources, by default it uses math.Rand and uses mutex locking to allow for safe goroutines.

If you want to use a more performant source please use NewUnlocked. Be aware that it is not goroutine safe.

```go
import rus_name_gen "github.com/oshokin/russian-name-generator"

// Uses math/rand(Pseudo) with mutex locking
faker := rus_name_gen.New(0)

// Uses math/rand(Pseudo) with NO mutext locking
// More performant but not goroutine safe.
faker := rus_name_gen.NewUnlocked(0)

// Uses crypto/rand(cryptographically secure) with mutext locking
faker := rus_name_gen.NewCrypto()

// Pass in your own random source
faker := rus_name_gen.NewCustom()
```

## Global Rand Set

If you would like to use the simple function calls but need to use something like
crypto/rand you can override the default global with the random source that you want.

```go
import rus_name_gen "github.com/oshokin/russian-name-generator"

faker := rus_name_gen.NewCrypto()
rus_name_gen.SetGlobalFaker(faker)
```

## Examples

```go
package main

import (
	"fmt"
	"sync"

	rus_name_gen "github.com/oshokin/russian-name-generator"
)

func main() {
	var (
		// Generate 100 full names concurrently
		namesCount = 100
		wg         sync.WaitGroup
	)

	wg.Add(namesCount)
	for i := 0; i < namesCount; i++ {
		go func(i int) {
			defer wg.Done()

			gender := rus_name_gen.GenderMale
			if i%2 == 0 {
				gender = rus_name_gen.GenderFemale
			}

			name := rus_name_gen.Name(gender, false)
			patronymic := rus_name_gen.Patronymic(gender == rus_name_gen.GenderFemale, false)
			surname := rus_name_gen.Surname(gender)

			fmt.Printf("%s %s %s\n", name, patronymic, surname)
		}(i)
	}

	wg.Wait()

	// Transliterate a Russian text to Latin letters
	text := "Брат, братан, братишка, когда меня отпустит?"
	transliteratedText := rus_name_gen.Transliterate(text)
	fmt.Println(transliteratedText)
}
```

## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.
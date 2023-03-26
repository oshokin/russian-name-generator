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
import rus_name_gen "github.com/[username]/russian-name-generator"
```

### Then, you can generate a random name:

```go
name := rus_name_gen.Name(rus_name_gen.GenderAny, true)
```

This will return a random name for any gender, excluding rare names.
Parameters:

`gender` (`Gender`): The gender for the generated name. Valid options are `GenderAny`, `GenderMale`, and `GenderFemale`.
`excludeRareNames` (`bool`): Whether or not to exclude rare names from the dataset.

### You can also generate a random surname:

```go
surname := rus_name_gen.Surname(rus_name_gen.GenderMale)
```

This will return a random surname for a male.
Parameters:

`gender` (`Gender`): The gender for the generated surname. Valid options are `GenderAny`, `GenderMale`, and `GenderFemale`.

### Finally, you can generate a random patronymic:

```go
patronymic := rus_name_gen.Patronymic(rus_name_gen.Any, false)
```

This will return a random patronymic for any gender, including rare names.
Parameters:

`gender` (`Gender`): The gender for the generated patronymic. Valid options are `GenderAny`, `GenderMale`, and `GenderFemale`.
`excludeRareNames` (`bool`): Whether or not to exclude rare names from the dataset.

### Also, you can transliterate a string from Russian to Latin letters:

```go
text := "Александр Сергеевич Пушкин"
transliteratedText := russian_name_generator.Transliterate(text)
```

This will return a transliterated string in Latin characters.
Parameters:

`text` (`string`): The Russian text to be transliterated to Latin characters.

## Examples

```go
package main

import (
	"fmt"
	"sync"
	"time"

	rus_name_gen "github.com/[username]/russian-name-generator"
)

func main() {
	// Generate 100 full names concurrently
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			gender := rus_name_gen.GenderMale
			if i%2 == 0 {
				gender = rus_name_gen.GenderFeMale
			}
            
            name := rus_name_gen.Name(gender, false)
			patronymic := rus_name_gen.Patronymic(gender, false)
			surname := rus_name_gen.Surname(gender)

			fmt.Printf("%s %s %s\n", name, patronymic, surname)
		}()
	}

	wg.Wait()

	// Transliterate a Russian text to Latin letters
	text := "Александр Сергеевич Пушкин"
	transliteratedText := rus_name_gen.Transliterate(text)
	fmt.Println(transliteratedText)
}
```

## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.
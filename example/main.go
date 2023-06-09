package main

import (
	"fmt"
	"sync"

	rus_name_gen "github.com/oshokin/russian-name-generator"
)

func main() {
	var (
		// Generate 100 full names concurrently.
		namesCount = 100
		wg         sync.WaitGroup
	)

	wg.Add(namesCount)

	for i := 0; i < namesCount; i++ {
		go func() {
			defer wg.Done()

			person := rus_name_gen.Person(nil)
			fmt.Println(person.Name,
				person.Patronymic,
				person.Surname)
		}()
	}

	wg.Wait()

	// Transliterate a Russian text to Latin letters.
	text := "Брат, братан, братишка, когда меня отпустит?"
	transliteratedText := rus_name_gen.Transliterate(text)
	fmt.Println(transliteratedText)
}

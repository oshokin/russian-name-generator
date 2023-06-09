package russian_name_generator

import (
	"fmt"
	"sync"
	"testing"
)

func Example() {
	Seed(2)
	fmt.Println("Surname:", Surname(GenderMale))
	fmt.Println("Name:", Name(GenderMale, true))
	fmt.Println("Patronymic:", Patronymic(false, true))
	// Output:
	// Surname: Палатов
	// Name: Руслан
	// Patronymic: Андреевич
}

func ExampleNew() {
	// Create new pseudo random faker struct and set initial seed
	fake := New(2)

	// All global functions are also available in the structs methods
	fmt.Println("Surname:", fake.Surname(GenderMale))
	fmt.Println("Name:", fake.Name(GenderMale, true))
	fmt.Println("Patronymic:", fake.Patronymic(false, true))
	// Output:
	// Surname: Палатов
	// Name: Руслан
	// Patronymic: Андреевич
}

func ExampleNewUnlocked() {
	fake := NewUnlocked(11)

	// All global functions are also available in the structs methods
	fmt.Println("Surname:", fake.Surname(GenderMale))
	fmt.Println("Name:", fake.Name(GenderMale, true))
	fmt.Println("Patronymic:", fake.Patronymic(false, true))
	// Output:
	// Surname: Губырин
	// Name: Олег
	// Patronymic: Антонович
}

func TestNewUnlocked(t *testing.T) {
	fake := NewUnlocked(0)
	if fake.Name(GenderMale, true) == "" {
		t.Error("Name was empty")
	}
}

func ExampleNewCrypto() {
	// Create new crypto faker struct
	fake := NewCrypto()

	// All global functions are also available in the structs methods
	fmt.Println("Surname:", fake.Surname(GenderMale))
	fmt.Println("Name:", fake.Name(GenderMale, true))
	fmt.Println("Patronymic:", fake.Patronymic(false, true))

	// Cannot output example as crypto/rand cant be predicted
}

func TestNewCrypto(t *testing.T) {
	// Create new crypto faker struct
	fake := NewCrypto()

	// All global functions are also available in the structs methods
	surname := fake.Surname(GenderMale)
	name := fake.Name(GenderMale, true)
	patronymic := fake.Patronymic(false, true)

	if name == "" || surname == "" || patronymic == "" {
		t.Error("One of the values was empty")
	}
}

type customRand struct{}

func (c *customRand) Seed(_ int64) {}

func (c *customRand) Uint64() uint64 {
	return 8675309
}

func (c *customRand) Int63() int64 {
	return int64(c.Uint64() & ^uint64(1<<63))
}

func ExampleNewCustom() {
	// Setup stuct and methods required to meet interface needs
	// type customRand struct {}
	// func (c *customRand) Seed(seed int64) {}
	// func (c *customRand) Uint64() uint64 { return 8675309 }
	// func (c *customRand) Int63() int64 { return int64(c.Uint64() & ^uint64(1<<63)) }

	// Create new custom faker struct
	fake := NewCustom(&customRand{})

	// All global functions are also available in the structs methods
	fmt.Println("Surname:", fake.Surname(GenderMale))
	fmt.Println("Name:", fake.Name(GenderMale, true))
	fmt.Println("Patronymic:", fake.Patronymic(false, true))
	// Output:
	// Surname: Абабилов
	// Name: Александр
	// Patronymic: Александрович
}

func ExampleSetGlobalFaker() {
	cryptoFaker := NewCrypto()
	SetGlobalFaker(cryptoFaker)
}

func TestSetGlobalFaker(t *testing.T) {
	cryptoFaker := NewCrypto()
	SetGlobalFaker(cryptoFaker)

	if name := Name(GenderMale, true); name == "" {
		t.Error("Name was empty")
	}

	// Set global back to psuedo
	SetGlobalFaker(New(0))
}

func TestConcurrency(_ *testing.T) {
	var (
		setupComplete sync.WaitGroup
		wg            sync.WaitGroup
	)

	setupComplete.Add(1)

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func() {
			setupComplete.Wait()
			Name(GenderMale, true)
			wg.Done()
		}()
	}

	setupComplete.Done()
	wg.Wait()
}

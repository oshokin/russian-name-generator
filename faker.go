package russian_name_generator

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sync"
)

type (
	// Faker is the primary struct for using localized.
	Faker struct {
		Rand *rand.Rand
	}

	lockedSource struct {
		lk  sync.Mutex
		src rand.Source64
	}

	cryptoRand struct {
		sync.Mutex
		buf []byte
	}
)

// Create global variable to deal with global function call.
var globalFaker = New(0)

func (r *lockedSource) Int63() int64 {
	r.lk.Lock()
	defer r.lk.Unlock()

	return r.src.Int63()
}

func (r *lockedSource) Uint64() uint64 {
	r.lk.Lock()
	defer r.lk.Unlock()

	return r.src.Uint64()
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	defer r.lk.Unlock()

	r.src.Seed(seed)
}

func (c *cryptoRand) Seed(_ int64) {}

func (c *cryptoRand) Uint64() uint64 {
	c.Lock()
	defer c.Unlock()

	if _, err := crand.Read(c.buf); err != nil {
		return 0
	}

	return binary.BigEndian.Uint64(c.buf)
}

func (c *cryptoRand) Int63() int64 {
	const lastUintBit = 63

	return int64(c.Uint64() & ^uint64(1<<lastUintBit))
}

// New will utilize math/rand for concurrent random usage.
// Setting seed to 0 will use crypto/rand for the initial seed number.
func New(seed int64) *Faker {
	// If passing 0 create crypto safe int64 for initial seed number
	if seed == 0 {
		_ = binary.Read(crand.Reader, binary.BigEndian, &seed)
	}

	return &Faker{
		Rand: rand.New(&lockedSource{
			src: rand.NewSource(seed).(rand.Source64),
		}),
	}
}

// NewUnlocked will utilize math/rand for non concurrent safe random usage.
// Setting seed to 0 will use crypto/rand for the initial seed number.
// NewUnlocked is more performant but not safe to run concurrently.
func NewUnlocked(seed int64) *Faker {
	// If passing 0 create crypto safe int64 for initial seed number.
	if seed == 0 {
		_ = binary.Read(crand.Reader, binary.BigEndian, &seed)
	}

	return &Faker{
		Rand: rand.New(rand.NewSource(seed)),
	}
}

// NewCrypto will utilize crypto/rand for concurrent random usage.
func NewCrypto() *Faker {
	const bufSize = 8

	return &Faker{
		Rand: rand.New(&cryptoRand{
			buf: make([]byte, bufSize),
		})}
}

// NewCustom will utilize a custom rand.Source64 for concurrent random usage
// See https://golang.org/src/math/rand/rand.go for required interface methods.
func NewCustom(source rand.Source64) *Faker {
	return &Faker{
		Rand: rand.New(source),
	}
}

// SetGlobalFaker will allow you to set what type of faker is globally used.
// Defailt is math/rand.
func SetGlobalFaker(faker *Faker) {
	globalFaker = faker
}

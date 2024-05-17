package randomstring

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

type LetterPoolConfig struct {
	Int       bool
	UpperCase bool
	LowerCase bool
}

var defaultConfig = LetterPoolConfig{
	LowerCase: true,
	UpperCase: true,
	Int:       true,
}

type LetterPoolOption func(*LetterPoolConfig)

func WithoutInt() LetterPoolOption {
	return func(c *LetterPoolConfig) {
		c.Int = false
	}
}

func WithoutUpperCase() LetterPoolOption {
	return func(c *LetterPoolConfig) {
		c.UpperCase = false
	}
}

func WithoutLowerCase() LetterPoolOption {
	return func(c *LetterPoolConfig) {
		c.LowerCase = false
	}
}

func GetLetterPoolConfig(options ...LetterPoolOption) LetterPoolConfig {
	config := defaultConfig
	for _, fu := range options {
		fu(&config)
	}
	return config
}

func GetLetterPool(c LetterPoolConfig) string {
	letterPool := ""
	if c.LowerCase {
		letterPool += "abcdefghijklmnopqrstuvwxyz"
	}
	if c.UpperCase {
		letterPool += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if c.Int {
		letterPool += "0123456789"
	}
	return letterPool
}

// Generator generates a random string of n characters, the generated string will of form ^[A-Za-z0-9]{n}$ by default, charset is customizable using options
type Generator struct {
	letterPool string
	src        *rand.Rand
	lock       sync.Mutex
}

func NewGenerator(options ...LetterPoolOption) *Generator {
	config := GetLetterPoolConfig(options...)
	letterPool := GetLetterPool(config)
	return &Generator{
		letterPool: letterPool,
		src:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *Generator) Generate(n int) string {
	if len(r.letterPool) == 0 {
		return ""
	}
	b := make([]byte, n)
	r.lock.Lock()
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, r.src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(r.letterPool) {
			b[i] = r.letterPool[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	r.lock.Unlock()
	return string(b)
}

var defaultRandomNumberGenerator = NewGenerator()

// Generate uses default RandomNumberGenerator object to generate random string that matches the regex the generated ^[A-Za-z0-9]{n}$
func Generate(n int) string {
	return defaultRandomNumberGenerator.Generate(n)
}

/*
GenerateWithPrefix generates a id with the prefix and totalLength
*/
func GenerateWithPrefix(totalLength int, prefix string) string {
	n := totalLength - len(prefix)
	return fmt.Sprintf("%v%v", prefix, Generate(n))
}

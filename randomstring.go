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

// Config represents the configuration options for generating random strings.
type Config struct {
	Int       bool // Int indicates whether to include integers in the generated string.
	UpperCase bool // UpperCase indicates whether to include uppercase letters in the generated string.
	LowerCase bool // LowerCase indicates whether to include lowercase letters in the generated string.
}

// defaultConfig represents the default configuration for generating random strings.
var defaultConfig = Config{
	LowerCase: true,
	UpperCase: true,
	Int:       true,
}

// LetterPoolOption represents functional options for configuring the letter pool for generating random strings.
type LetterPoolOption func(*Config)

// WithoutInt excludes integers from the letter pool when generating random strings.
func WithoutInt() LetterPoolOption {
	return func(c *Config) {
		c.Int = false
	}
}

// WithoutUpperCase excludes uppercase letters from the letter pool when generating random strings.
func WithoutUpperCase() LetterPoolOption {
	return func(c *Config) {
		c.UpperCase = false
	}
}

// WithoutLowerCase excludes lowercase letters from the letter pool when generating random strings.
func WithoutLowerCase() LetterPoolOption {
	return func(c *Config) {
		c.LowerCase = false
	}
}

// GetLetterPoolConfig returns the configuration for the letter pool based on the provided options.
func GetLetterPoolConfig(options ...LetterPoolOption) Config {
	config := defaultConfig
	for _, fu := range options {
		fu(&config)
	}
	return config
}

// GetLetterPool returns the letter pool based on the provided configuration.
func GetLetterPool(c Config) string {
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

// Generator generates random strings based on the specified letter pool.
type Generator struct {
	letterPool string
	src        *rand.Rand
	lock       sync.Mutex
}

// NewGenerator creates a new Generator instance with the specified options for generating random strings.
func NewGenerator(options ...LetterPoolOption) *Generator {
	config := GetLetterPoolConfig(options...)
	letterPool := GetLetterPool(config)
	return &Generator{
		letterPool: letterPool,
		src:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Generate generates a random string of n characters using the specified letter pool.
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

// defaultRandomNumberGenerator represents the default Generator instance for generating random strings.
var defaultRandomNumberGenerator = NewGenerator()

// Generate generates a random string of n characters using the default Generator instance.
func Generate(n int) string {
	return defaultRandomNumberGenerator.Generate(n)
}

// GenerateWithPrefix generates a random string with the specified prefix and total length.
func GenerateWithPrefix(totalLength int, prefix string) string {
	n := totalLength - len(prefix)
	return fmt.Sprintf("%v%v", prefix, Generate(n))
}

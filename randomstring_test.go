package randomstring_test

import (
	"fmt"
	"regexp"
	"runtime"
	"sync"
	"testing"

	"github.com/sabariramc/randomstring"
	"gotest.tools/assert"
)

func ExampleGenerateWithPrefix() {
	x := randomstring.GenerateWithPrefix(20, "cust_")
	fmt.Println(len(x))
	match, _ := regexp.MatchString("^cust_[a-zA-Z0-9]{15}$", x)
	fmt.Println(match)
	//Output:
	//20
	//true
}

func ExampleGenerator() {
	gen := randomstring.NewGenerator()
	x := gen.Generate(10)
	match, _ := regexp.MatchString("^[a-zA-Z0-9]{10}$", x)
	fmt.Println(match)
	//Output:
	//true
}

func ExampleGenerator_onlynumerals() {
	gen := randomstring.NewGenerator(randomstring.WithoutLowerCase(), randomstring.WithoutUpperCase())
	x := gen.Generate(10)
	match, _ := regexp.MatchString("^[0-9]{10}$", x)
	fmt.Println(match)
	//Output:
	//true
}

func ExampleGenerator_onlyuppercase() {
	gen := randomstring.NewGenerator(randomstring.WithoutLowerCase(), randomstring.WithoutInt())
	x := gen.Generate(10)
	match, _ := regexp.MatchString("^[A-Z]{10}$", x)
	fmt.Println(match)
	//Output:
	//true
}

func ExampleGenerator_onlylowercase() {
	gen := randomstring.NewGenerator(randomstring.WithoutUpperCase(), randomstring.WithoutInt())
	x := gen.Generate(10)
	match, _ := regexp.MatchString("^[a-z]{10}$", x)
	fmt.Println(match)
	//Output:
	//true
}

func generatorDuplicateCheck(t *testing.T, cnt, n int) {
	ch := make(chan string, cnt)
	var wg sync.WaitGroup
	concurrencyFactor := 10000
	for i := 0; i < concurrencyFactor; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < cnt/concurrencyFactor; j++ {
				ch <- randomstring.Generate(n)
			}
			wg.Done()
		}()
	}
	wg.Add(1)
	duplicateCount := 0
	go func() {
		idSet := make(map[string]bool, cnt)
		total := 0
		for id := range ch {
			if _, ok := idSet[id]; ok {
				duplicateCount++
			}
			idSet[id] = true
			total++
			if total == cnt {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
	assert.Equal(t, duplicateCount, 0)
}

func TestGenerator(t *testing.T) {
	generatorDuplicateCheck(t, 10000, 5)
	generatorDuplicateCheck(t, 10000000, 10)
}

func BenchmarkGenerator(b *testing.B) {
	var goprocs = runtime.GOMAXPROCS(0)
	gen := randomstring.NewGenerator()
	for i := 1; i < 1000; i += 50 {
		b.Run(fmt.Sprintf("goroutines-%d", i*goprocs), func(b *testing.B) {
			b.SetParallelism(i)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					gen.Generate(20)
				}
			})
		})
	}
}

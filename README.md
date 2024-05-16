# Random String Generator

Random string generator, based on [stack overflow answer]. Depending on the length of the string can produce unique random string


[stack overflow answer]: https://stackoverflow.com/a/31832326/3769802


## Usage

```go
x := randomstring.Generate()
```

For advance usage refer [test file]


## Benchmark

Benchmark code can be found in [test file]

[test file]: randomstring_test.go

```
goos: linux
goarch: amd64
pkg: github.com/sabariramc/randomstring
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkGenerator/goroutines-8-8         	 8104476	       152.7 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-408-8       	 3814875	       318.4 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-808-8       	 3775564	       318.3 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-1208-8      	 4145314	       320.1 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-1608-8      	 3756034	       316.1 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-2008-8      	 3762130	       311.9 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-2408-8      	 4358134	       316.0 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-2808-8      	 4085209	       307.3 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-3208-8      	 3798603	       299.9 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-3608-8      	 4180372	       313.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-4008-8      	 4070395	       313.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-4408-8      	 4294604	       307.7 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-4808-8      	 4421068	       309.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-5208-8      	 4406760	       304.3 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-5608-8      	 4148553	       304.8 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-6008-8      	 3639664	       301.9 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-6408-8      	 4420321	       312.1 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-6808-8      	 3376665	       457.2 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-7208-8      	 3020142	       431.3 ns/op	      48 B/op	       2 allocs/op
BenchmarkGenerator/goroutines-7608-8      	 2765803	       433.3 ns/op	      48 B/op	       2 allocs/op
```
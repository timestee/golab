Version: 
>go version go1.12 darwin/amd64

Result:
```shell script
goos: darwin
goarch: amd64
Benchmark_NormalSwitch-12              	2000000000	         1.40 ns/op
Benchmark_InterfaceSwitch-12           	1000000000	         2.04 ns/op
Benchmark_interfaceSwitch-12           	200000000	         7.64 ns/op
Benchmark_TypeSwitch-12                	2000000000	         1.16 ns/op
Benchmark_TypeAssertionInterface-12    	200000000	         8.14 ns/op
Benchmark_TypeAssertionToPointer-12    	2000000000	         1.19 ns/op
PASS
```
- once go knows the type of something, it can generate optimal machine code for that type; it doesn't matter if the type was computed using type assertions or if it always had that type
- computing the type of something using a type assert involves at least one "if" statement, but "if" statements are not really slow (unless you're causing a lot of mispredictions)
- if it's an interface type, go can use virtual method calls, which are a little bit slower (a few extra instructions) but most of the time you shouldn't care
- if it's a concrete type, go just makes a regular function call, it can be inlined, etc
- if it isn't a pointer or interface, it might need to be copied, which could add overhead - but nothing unusual for value types

> - In real code, type assertion overhead will not show up (unless you're doing it in a loop 100,000 times per request or something)
> - `interface{} ` still a little evil.

Version:
>go version go1.12 darwin/amd64

Result:
```shell script
goos: darwin
goarch: amd64
Benchmark_Normal_Len_100-12                        	50000000	        32.2 ns/op
Benchmark_ByteStringStringsBuilder_Len_100-12      	2000000000	         0.53 ns/op
Benchmark_ByteStringKeepAlive_Len_100-12           	2000000000	         0.81 ns/op
Benchmark_ByteStringKeepAliveOff_Len_100-12        	2000000000	         0.54 ns/op
Benchmark_Normal_Len_10000-12                      	 2000000	       868 ns/op
Benchmark_ByteStringStringsBuilder_Len_10000-12    	2000000000	         0.53 ns/op
Benchmark_ByteStringKeepAlive_Len_10000-12         	2000000000	         0.81 ns/op
Benchmark_ByteStringKeepAliveOff_Len_10000-12      	2000000000	         0.54 ns/op
```

> [Feature: provide no-copy conversion from []byte to string #25484](https://github.com/golang/go/issues/25484)

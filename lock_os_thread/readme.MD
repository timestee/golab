Version: 
>go version go1.12 darwin/amd64

Result:
```shell script
goos: darwin
goarch: amd64
Benchmark_Normal-12        	10000000	       196 ns/op
Benchmark_LockThread-12    	  200000	      6348 ns/op
```

Golang supports LockOSThread() to make current goroutine exclusively tied to current OS thread, and it can also UnlockOSThread().

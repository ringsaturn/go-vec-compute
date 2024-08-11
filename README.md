```
goos: darwin
goarch: arm64
pkg: github.com/ringsaturn/go-vec-compute
BenchmarkApparentTemperature_Go/10_hours-16         	11637430	       100.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkApparentTemperature_Go/120_hours-16        	 1000000	      1178 ns/op	       0 B/op	       0 allocs/op
BenchmarkApparentTemperature_Go/384_hours-16        	  325849	      3762 ns/op	       0 B/op	       0 allocs/op
BenchmarkApparentTemperature_Go/480_hours-16        	  257200	      4689 ns/op	       0 B/op	       0 allocs/op
BenchmarkApparentTemperature_Arrow/10_hours-16      	 2426258	       493.1 ns/op	     976 B/op	       9 allocs/op
BenchmarkApparentTemperature_Arrow/120_hours-16     	  366506	      3223 ns/op	    2704 B/op	      11 allocs/op
BenchmarkApparentTemperature_Arrow/384_hours-16     	  122948	      9811 ns/op	    9872 B/op	      13 allocs/op
BenchmarkApparentTemperature_Arrow/480_hours-16     	   98649	     12058 ns/op	    9872 B/op	      13 allocs/op
PASS
coverage: 100.0% of statements
ok  	github.com/ringsaturn/go-vec-compute	10.934s
```

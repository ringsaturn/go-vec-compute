```
goos: darwin
goarch: arm64
pkg: github.com/ringsaturn/go-vec-compute
BenchmarkApparentTemperature_PureGo/10_hours-16         11278897                92.82 ns/op
BenchmarkApparentTemperature_PureGo/120_hours-16         1000000              1085 ns/op
BenchmarkApparentTemperature_PureGo/384_hours-16          342058              3478 ns/op
BenchmarkApparentTemperature_PureGo/480_hours-16          273854              4339 ns/op
BenchmarkApparentTemperature_GoIterArrowRecords/10_hours-16              2488632               480.3 ns/op
BenchmarkApparentTemperature_GoIterArrowRecords/120_hours-16              369253              3131 ns/op
BenchmarkApparentTemperature_GoIterArrowRecords/384_hours-16              126391              9425 ns/op
BenchmarkApparentTemperature_GoIterArrowRecords/480_hours-16              102184             11534 ns/op
PASS
ok      github.com/ringsaturn/go-vec-compute    11.015s
```

| Benchmark Hours       | 10 hours | 120 hours | 384 hours | 480 hours |
| --------------------- | -------- | --------- | --------- | --------- |
| Pure Go               | 92.82 ns | 1085 ns   | 3478 ns   | 4339 ns   |
| Go Iter Arrow Records | 480.3 ns | 3131 ns   | 9425 ns   | 11534 ns  |

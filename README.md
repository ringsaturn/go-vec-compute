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

```
----------------------------------------------------------------------------------------------- benchmark: 4 tests -----------------------------------------------------------------------------------------------
Name (time in us)                                    Min               Max              Mean            StdDev            Median               IQR            Outliers  OPS (Kops/s)            Rounds  Iterations
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
test_compute_apparent_temperature_by_hour_10      4.4600 (1.0)      5.7992 (1.0)      4.8217 (1.0)      0.2396 (1.02)     4.8000 (1.0)      0.2219 (1.75)         34;5      207.3974 (1.0)         100         100
test_compute_apparent_temperature_by_hour_120     4.8325 (1.08)     6.0542 (1.04)     5.1638 (1.07)     0.2339 (1.0)      5.2190 (1.09)     0.3200 (2.52)         32;3      193.6560 (0.93)        100         100
test_compute_apparent_temperature_by_hour_384     6.1613 (1.38)     8.8283 (1.52)     6.8640 (1.42)     0.5929 (2.53)     6.6523 (1.39)     0.2823 (2.23)        24;21      145.6882 (0.70)        100         100
test_compute_apparent_temperature_by_hour_480     6.2837 (1.41)     8.4413 (1.46)     6.8097 (1.41)     0.2681 (1.15)     6.7898 (1.41)     0.1269 (1.0)         21;21      146.8491 (0.71)        100         100
------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
```

| Benchmark Hours       | 10 hours  | 120 hours | 384 hours | 480 hours |
| --------------------- | --------- | --------- | --------- | --------- |
| Pure Go               | 92.82 ns  | 1085 ns   | 3478 ns   | 4339 ns   |
| Go Iter Arrow Records | 480.3 ns  | 3131 ns   | 9425 ns   | 11534 ns  |
| Python NumPy Vec      | 4.8217 us | 5.1638 us | 6.8640 us | 6.8097 us |

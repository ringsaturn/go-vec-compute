import numpy as np

# // https://w.wiki/AtGR
# func ComputeApparentTemperature(t float64, pres float64, ws float64, rh float64) float64 {
# 	e := rh / 100 * 6.105 * math.Exp(17.27*t/(237.7+t))
# 	at := 1.07*t + 0.2*e - 0.65*ws - 2.7
# 	return at
# }


def compute_apparent_temperature(t, ws, rh):
    e = rh / 100 * 6.105 * np.exp(17.27 * t / (237.7 + t))
    at = 1.07 * t + 0.2 * e - 0.65 * ws - 2.7
    return at


def rands_in_range(n, low, high):
    return np.random.uniform(low, high, n)


def _bench_compute_apparent_temperature(tm, ws, rh):
    _ = compute_apparent_temperature(tm, ws, rh)


def test_compute_apparent_temperature_by_hour_10(benchmark):
    hour = 10
    tmps = rands_in_range(hour, -10, 40)
    wss = rands_in_range(hour, 0, 20)
    rhs = rands_in_range(hour, 0, 100)
    benchmark.pedantic(
        _bench_compute_apparent_temperature,
        args=(tmps, wss, rhs),
        iterations=100,
        rounds=100,
    )


def test_compute_apparent_temperature_by_hour_120(benchmark):
    hour = 120
    tmps = rands_in_range(hour, -10, 40)
    wss = rands_in_range(hour, 0, 20)
    rhs = rands_in_range(hour, 0, 100)
    benchmark.pedantic(
        _bench_compute_apparent_temperature,
        args=(tmps, wss, rhs),
        iterations=100,
        rounds=100,
    )


def test_compute_apparent_temperature_by_hour_384(benchmark):
    hour = 384
    tmps = rands_in_range(hour, -10, 40)
    wss = rands_in_range(hour, 0, 20)
    rhs = rands_in_range(hour, 0, 100)
    benchmark.pedantic(
        _bench_compute_apparent_temperature,
        args=(tmps, wss, rhs),
        iterations=100,
        rounds=100,
    )


def test_compute_apparent_temperature_by_hour_480(benchmark):
    hour = 480
    tmps = rands_in_range(hour, -10, 40)
    wss = rands_in_range(hour, 0, 20)
    rhs = rands_in_range(hour, 0, 100)
    benchmark.pedantic(
        _bench_compute_apparent_temperature,
        args=(tmps, wss, rhs),
        iterations=100,
        rounds=100,
    )

package goveccompute

import (
	"math"
	"math/rand/v2"
)

// https://w.wiki/AtGR
func ComputeApparentTemperature(t float64, pres float64, ws float64, rh float64) float64 {
	e := rh / 100 * 6.105 * math.Exp(17.27*t/(237.7+t))
	at := 1.07*t + 0.2*e - 0.65*ws - 2.7
	return at
}

func RandsInRange(n int, min, max float64) []float64 {
	rands := make([]float64, n)
	for i := 0; i < n; i++ {
		rands[i] = min + (max-min)*rand.Float64()
	}
	return rands
}

package goveccompute_test

import (
	"fmt"
	"testing"

	goveccompute "github.com/ringsaturn/go-vec-compute"
)

func benchmarkApparentTemperature_Go(b *testing.B, l int) {
	temps := goveccompute.RandsInRange(l, 20, 40)
	humidities := goveccompute.RandsInRange(l, 50, 90)
	pressures := goveccompute.RandsInRange(l, 1000, 1020)
	windspeeds := goveccompute.RandsInRange(l, 1, 10)
	apparentTemps := make([]float64, l) // require to be filled

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		func() {
			for i := 0; i < len(temps); i++ {
				temp := temps[i]
				humidity := humidities[i]
				pressure := pressures[i]
				windspeed := windspeeds[i]

				apparentTemp := goveccompute.ComputeApparentTemperature(temp, pressure, windspeed, humidity)
				apparentTemps[i] = apparentTemp
			}
		}()
	}
}

func BenchmarkApparentTemperature_Go(b *testing.B) {
	for _, hours := range []int{10, 120, 384, 480} {
		b.Run(fmt.Sprintf("%v hours", hours), func(b *testing.B) {
			benchmarkApparentTemperature_Go(b, hours)
		})
	}
}

package etlutil

import "math"

func Round2(val float64) float64 {
	return math.Round(val*100) / 100
}

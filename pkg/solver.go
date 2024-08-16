package pkg

import (
	"math"
)

func EquationSolver(a, b, c float64) (float64, float64) {
	D := math.Pow(b, 2) - 4*(a*c)
	x1 := (-b + math.Sqrt(D)) / (2 * a)
	x2 := (-b - math.Sqrt(D)) / (2 * a)

	return x1, x2
}

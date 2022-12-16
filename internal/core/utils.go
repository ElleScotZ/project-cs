package core

import (
	"math"

	"github.com/ElleScotZ/project-cs/internal/algebra"
)

// DistributeInterval creates a numberOfPoints-sized slice (0,1)
// of floating point numbers - with EPSILON accuracy.
func DistributeInterval(numberOfPoints int) []float64 {
	var (
		parameters         = make([]float64, numberOfPoints)
		accuracyReciprocal = math.Round(1.0 / (0.1 * algebra.EPSILON))
	)

	parameters[0] = algebra.EPSILON

	delta := (1.0 - 2.0*algebra.EPSILON) / float64(numberOfPoints-1)

	for i := 1; i < numberOfPoints; i++ {
		parameters[i] = math.Round(float64(i)*delta*accuracyReciprocal) / accuracyReciprocal
	}

	parameters[numberOfPoints-1] -= algebra.EPSILON

	return parameters
}

package surface

import (
	"errors"

	"github.com/ElleScotZ/project-cs/internal/algebra"
)

// Knot represents a knot in the parametrised 2D space on a T-mesh.
type Knot struct {
	Position [2]float64 // [0]: horizontal coordinate, [1]: vertical coordinate
}

// ControlPoint represents a control point on a T-Spline surface.
type ControlPoint struct {
	Position   algebra.Vector3D
	Knotvector [5]Knot
}

// TSpline represents a T-Spline surface.
type TSpline struct {
	ControlPoints []ControlPoint
}

// getBasisMultiplier
func getBasisMultiplier(knotvector []float64, iterBasis, iterDegree int, parameter float64) float64 {
	if iterBasis+iterDegree >= len(knotvector) {
		return 1
	}

	numerator := (parameter - knotvector[iterBasis])
	denominator := knotvector[iterBasis+iterDegree] - knotvector[iterBasis]

	if denominator != 0 {
		return numerator / denominator
	}

	return 0
}

// calculateBasisFunction
func calculateBasisFunction(knotvector []float64, iterBasis, iterDegree int, parameter float64) (float64, error) {
	if iterBasis < 0 || iterDegree < 0 {
		return -1, errors.New("calculateBasisFunction with wrong iterBasis or iterDegree")
	}

	// B_(iterBasis)_0(parameter) is a step function.
	if iterDegree == 0 && parameter < knotvector[iterBasis+1] && parameter >= knotvector[iterBasis] {
		return 1, nil
	} else if iterDegree == 0 {
		return 0, nil
	}

	// Calculating multipliers for both members of basis function
	multiplierPO := getBasisMultiplier(knotvector, iterBasis, iterDegree, parameter)
	multiplierNC := getBasisMultiplier(knotvector, iterBasis+1, iterDegree, parameter)
	multiplierNC = 1 - multiplierNC

	// Calculating previous iterations
	basisPrevOrder, errPO := calculateBasisFunction(knotvector, iterBasis, iterDegree-1, parameter)
	basisNextControl, errNC := calculateBasisFunction(knotvector, iterBasis+1, iterDegree-1, parameter)

	// Value of basis function at parameter.
	basis := multiplierPO*basisPrevOrder + multiplierNC*basisNextControl

	if errPO != nil {
		return basis, errPO
	}

	return basis, errNC
}

// calculateBlendingFunction
func calculateBlendingFunction(knotvector [5]Knot, paramS, paramT float64) (float64, error) {
	var knotVectorS, knotVectorT = make([]float64, 5), make([]float64, 5)

	for i := 0; i < 5; i++ {
		knotVectorS[i] = knotvector[i].Position[0]
		knotVectorT[i] = knotvector[i].Position[1]
	}

	basis1, err := calculateBasisFunction(knotVectorS, 0, 3, paramS)
	if err != nil {
		return -1, err
	}

	basis2, err := calculateBasisFunction(knotVectorT, 0, 3, paramT)
	if err != nil {
		return -1, err
	}

	return basis1 * basis2, nil
}

// CalculateTSplineSurfacePoint generates a point on surface with paramS and paramT.
func (t *TSpline) CalculateTSplineSurfacePoint(paramS, paramT float64) (algebra.Vector3D, error) {
	var (
		position algebra.Vector3D
		divisor  float64
	)

	for _, cp := range t.ControlPoints {
		blending, err := calculateBlendingFunction(cp.Knotvector, paramS, paramT)
		if err != nil {
			return position, err
		}

		position = position.Add(cp.Position.Mul(blending))
		divisor += blending
	}

	position = position.Div(divisor)

	return position, nil
}

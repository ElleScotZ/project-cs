package surface

import (
	"errors"

	"github.com/ElleScotZ/project-cs/internal/algebra"
)

const (
	Degree = 3
)

// BSpline represents a B-Spline surface with degree = 3.
type BSpline struct {
	ControlPointMatrix [][]ControlPoint
}

//
func (b *BSpline) SetNormalisedUniformKnotvectorsForClampedSurface() error {
	if len(b.ControlPointMatrix) == 0 || len(b.ControlPointMatrix[0]) == 0 {
		return errors.New("b.ControlPointMatrix dimension(s) 0")
	}

	numberOfCP := len(b.ControlPointMatrix[0]) * len(b.ControlPointMatrix)

	delta := 1.0 / float64(numberOfCP-Degree)

	var knotvector = make([]float64, (Degree+1)+numberOfCP)

	for i := range knotvector {
		switch {
		case i < Degree+1:
			knotvector[i] = 0.0
		case i < numberOfCP-(Degree+1):
			knotvector[i] = delta * float64(i-Degree)
		default:
			knotvector[i] = 1.0
		}
	}

	for i := range b.ControlPointMatrix {
		for j := range b.ControlPointMatrix[i] {
			for k := 0; k < 5; k++ {
				b.ControlPointMatrix[i][j].Knotvector[k].Position[0] = knotvector[i*len(b.ControlPointMatrix)+j+k]
				b.ControlPointMatrix[i][j].Knotvector[k].Position[1] = knotvector[i*len(b.ControlPointMatrix)+j+k]
			}
		}
	}

	return nil
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

//
func (b *BSpline) CalculateRationalSurfacePoint(paramS, paramT float64) (algebra.Vector3D, error) {
	var position algebra.Vector3D

	if b.ControlPointMatrix == nil {
		return position, errors.New("generate ControlPointMatrix before calling CalculateRationalSurfacePoint")
	}

	for i := range b.ControlPointMatrix {
		var member algebra.Vector3D

		for j := 0; j < len(b.ControlPointMatrix[i]); j++ {
			var knotvectorS, knotvectorT = make([]float64, 5), make([]float64, 5)

			for k := 0; k < 5; k++ {
				knotvectorS[i] = b.ControlPointMatrix[i][j].Knotvector[k].Position[0]
				knotvectorT[i] = b.ControlPointMatrix[i][j].Knotvector[k].Position[1]
			}

			basisS, err := calculateBasisFunction(knotvectorS, i, Degree, paramS)
			if err != nil {
				return position, err
			}

			basisT, err := calculateBasisFunction(knotvectorT, j, Degree, paramT)
			if err != nil {
				return position, err
			}

			member = member.Add(b.ControlPointMatrix[i][j].Position.Mul(basisS * basisT))
		}

		position = position.Add(member)
	}

	return position, nil
}

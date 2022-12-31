package surface

import (
	"errors"

	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/internal/core"
	"github.com/ElleScotZ/project-cs/pkg/io"
)

const (
	Degree = 3
)

// BSpline represents a B-Spline surface with degree = 3.
type BSpline struct {
	Knotvector         []float64
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
		case i < numberOfCP:
			knotvector[i] = delta * float64(i-Degree)
		default:
			knotvector[i] = 1.0
		}
	}

	b.Knotvector = knotvector

	// This is just an FYI. Probably will not be used.
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
	basisPrevDegree, errPO := calculateBasisFunction(knotvector, iterBasis, iterDegree-1, parameter)
	basisNextControl, errNC := calculateBasisFunction(knotvector, iterBasis+1, iterDegree-1, parameter)

	// Value of basis function at parameter.
	basis := multiplierPO*basisPrevDegree + multiplierNC*basisNextControl

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
			basisS, err := calculateBasisFunction(b.Knotvector, i, Degree, paramS)
			if err != nil {
				return position, err
			}

			basisT, err := calculateBasisFunction(b.Knotvector, j, Degree, paramT)
			if err != nil {
				return position, err
			}

			member = member.Add(b.ControlPointMatrix[i][j].Position.Mul(basisS * basisT))
		}

		position = position.Add(member)
	}

	return position, nil
}

//
func (b *BSpline) GenerateSurface(resolution [2]int, fileName string) error {
	var mesh core.Mesh

	parameters1 := core.DistributeInterval(resolution[0])
	parameters2 := core.DistributeInterval(resolution[1])

	// Creating vertices
	for _, p1 := range parameters1 {
		for _, p2 := range parameters2 {
			point, err := b.CalculateRationalSurfacePoint(p1, p2)
			if err != nil {
				return err
			}

			mesh.Vertices = append(mesh.Vertices, core.Vertex{
				Position: point,
				Color:    algebra.Vector3D{Coordinates: [3]float64{0, 150, 0}},
			})
		}
	}

	// Creating faces
	for i := 0; i < len(parameters1)-1; i++ {
		for j := 0; j < len(parameters2)-1; j++ {
			var (
				face1, face2       core.Face
				id1, id2, id3, id4 int
			)

			id1 = len(parameters2)*i + j
			id2 = id1 + 1
			id3 = id1 + len(parameters2)
			id4 = id3 + 1

			// Checking length of edge
			distance14Vector := mesh.Vertices[id1].Position.Sub(mesh.Vertices[id4].Position)
			distance14 := distance14Vector.Dot(distance14Vector)

			distance23Vector := mesh.Vertices[id2].Position.Sub(mesh.Vertices[id3].Position)
			distance23 := distance23Vector.Dot(distance23Vector)

			if distance23 < distance14 {
				id1, id2, id3, id4 = id2, id4, id1, id3
			}

			face1.Vertices = [3]int{id2, id1, id4}
			face2.Vertices = [3]int{id4, id1, id3}

			mesh.Faces = append(mesh.Faces, face1)
			mesh.Faces = append(mesh.Faces, face2)
		}
	}

	// Adding control points
	for _, cpR := range b.ControlPointMatrix {
		for _, cp := range cpR {
			mesh.Vertices = append(mesh.Vertices, core.Vertex{
				Position: cp.Position,
				Color:    algebra.Vector3D{Coordinates: [3]float64{150, 0, 0}},
			})
		}
	}

	return io.ExportPLY(&mesh, fileName)
}

package surface

import (
	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/internal/core"
	"github.com/ElleScotZ/project-cs/pkg/io"
)

// TSpline represents a T-Spline surface.
type TSpline struct {
	ControlPoints []ControlPoint
}

// calculateBlendingFunction
func calculateBlendingFunction(knotvector [5]Knot, paramS, paramT float64) (float64, error) {
	var knotVectorS, knotVectorT = make([]float64, 5), make([]float64, 5)

	for i := 0; i < 5; i++ {
		knotVectorS[i] = knotvector[i].Position[0]
		knotVectorT[i] = knotvector[i].Position[1]
	}

	basis1, err := calculateBasisFunction(knotVectorS, 0, Degree, paramS)
	if err != nil {
		return -1, err
	}

	basis2, err := calculateBasisFunction(knotVectorT, 0, Degree, paramT)
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

//
func (t *TSpline) GenerateSurface(resolution [2]int, fileName string) error {
	var mesh core.Mesh

	parameters1 := core.DistributeInterval(resolution[0])
	parameters2 := core.DistributeInterval(resolution[1])

	// Creating vertices
	for _, p1 := range parameters1 {
		for _, p2 := range parameters2 {
			point, err := t.CalculateTSplineSurfacePoint(p1, p2)
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
	for _, cp := range t.ControlPoints {
		mesh.Vertices = append(mesh.Vertices, core.Vertex{
			Position: cp.Position,
			Color:    algebra.Vector3D{Coordinates: [3]float64{150, 0, 0}},
		})
	}

	return io.ExportPLY(&mesh, fileName)
}

package surface

import (
	"testing"

	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/internal/core"
	"github.com/ElleScotZ/project-cs/pkg/io"
)

func TestBSplineSurface(t *testing.T) {
	t.Parallel()

	var (
		bspline BSpline
		cp      [16]ControlPoint
	)

	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 0.0, 5.0}}
	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 5.0}}
	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 1.5, 0.0}}
	cp[7].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 1.50, 0.0}}
	cp[8].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 3.5, 0.0}}
	cp[11].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 3.50, 0.0}}
	cp[13].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 5.0, 5.0}}
	cp[14].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 5.0, 5.0}}

	cp[0].Position = cp[1].Position.Add(cp[4].Position).Div(2.0)
	cp[3].Position = cp[2].Position.Add(cp[7].Position).Div(2.0)
	cp[5].Position = cp[1].Position.Mul(0.67).Add(cp[13].Position.Mul(0.33))
	cp[5].Position.Coordinates[2] = 4.0
	cp[6].Position = cp[2].Position.Mul(0.67).Add(cp[14].Position.Mul(0.33))
	cp[6].Position.Coordinates[2] = 4.0
	cp[9].Position = cp[1].Position.Mul(0.33).Add(cp[13].Position.Mul(0.67))
	cp[9].Position.Coordinates[2] = 4.0
	cp[10].Position = cp[2].Position.Mul(0.33).Add(cp[14].Position.Mul(0.67))
	cp[10].Position.Coordinates[2] = 4.0
	cp[12].Position = cp[13].Position.Add(cp[8].Position).Div(2.0)
	cp[15].Position = cp[11].Position.Add(cp[14].Position).Div(2.0)

	bspline.ControlPointMatrix = make([][]ControlPoint, 4)

	for i := range bspline.ControlPointMatrix {
		bspline.ControlPointMatrix[3-i] = make([]ControlPoint, 4)

		for j := range bspline.ControlPointMatrix[3-i] {
			bspline.ControlPointMatrix[3-i][j].Position = cp[4*i+j].Position
		}
	}

	bspline.SetNormalisedUniformKnotvectorsForClampedSurface()

	err := bspline.GenerateSurface([2]int{50, 50}, "saddle_bspline")
	if err != nil {
		t.Error(err)
	}
}

func TestBSplineBasis(t *testing.T) {
	t.Parallel()

	var m core.Mesh

	parameters := core.DistributeInterval(100)

	for _, p := range parameters {
		basis, err := calculateBasisFunction([]float64{0.0, 0.2, 0.4, 0.6, 0.8, 1.0, 1.2, 1.4, 1.6, 1.8, 2.0}, 3, 3, p*2)
		if err != nil {
			t.Error(err)
		}

		m.Vertices = append(m.Vertices, core.Vertex{Position: algebra.Vector3D{Coordinates: [3]float64{p * 2, basis, 0.0}}})
	}

	err := io.ExportPLY(&m, "basis")
	if err != nil {
		t.Error(err)
	}
}

package surface

import (
	"testing"

	"github.com/ElleScotZ/project-cs/internal/algebra"
)

func TestSplineCurve(t *testing.T) {
	var (
		tspline TSpline
		cp      [8]ControlPoint
	)

	cp[0].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 0.0, 0.0}}
	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{0.50, 0.0, 0.0}}
	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 0.0}}
	cp[3].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.50, 0.0}}
	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 1.0, 0.0}}
	cp[5].Position = algebra.Vector3D{Coordinates: [3]float64{0.50, 1.0, 0.0}}
	cp[6].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 1.0, 0.0}}
	cp[7].Position = algebra.Vector3D{Coordinates: [3]float64{0.50, 0.50, 0.50}}

	cp[0].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.50}},
		{Position: [2]float64{0.0, 1.00}},
		{Position: [2]float64{0.20, 1.0}}}

	cp[1].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.50}},
		{Position: [2]float64{0.20, 1.00}},
		{Position: [2]float64{0.40, 1.0}}}

	cp[2].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.20, 0.50}},
		{Position: [2]float64{0.40, 1.00}},
		{Position: [2]float64{0.60, 1.0}}}

	cp[3].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.20, 0.0}},
		{Position: [2]float64{0.40, 0.50}},
		{Position: [2]float64{0.60, 1.00}},
		{Position: [2]float64{0.80, 1.0}}}

	cp[4].Knotvector = [5]Knot{
		{Position: [2]float64{0.20, 0.0}},
		{Position: [2]float64{0.40, 0.0}},
		{Position: [2]float64{0.60, 0.50}},
		{Position: [2]float64{0.80, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[5].Knotvector = [5]Knot{
		{Position: [2]float64{0.40, 0.0}},
		{Position: [2]float64{0.60, 0.0}},
		{Position: [2]float64{0.80, 0.50}},
		{Position: [2]float64{1.0, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[6].Knotvector = [5]Knot{
		{Position: [2]float64{0.60, 0.0}},
		{Position: [2]float64{0.80, 0.0}},
		{Position: [2]float64{1.0, 0.50}},
		{Position: [2]float64{1.0, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[7].Knotvector = [5]Knot{
		{Position: [2]float64{0.80, 0.0}},
		{Position: [2]float64{1.0, 0.0}},
		{Position: [2]float64{1.0, 0.50}},
		{Position: [2]float64{1.0, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "curve")
	if err != nil {
		t.Error(err)
	}
}

func TestSplineSurface(t *testing.T) {
	var (
		tspline TSpline
		cp      [5]ControlPoint
	)

	cp[0].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 0.0, 0.0}}
	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 0.0}}
	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 1.0, 0.0}}
	cp[3].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 1.0, 0.0}}
	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{0.50, 0.50, 0.50}}

	cp[0].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[1].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.50, 0.0}},
		{Position: [2]float64{1.0, 0.0}},
		{Position: [2]float64{1.0, 0.50}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[2].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.50}},
		{Position: [2]float64{0.0, 1.0}},
		{Position: [2]float64{0.50, 1.0}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[3].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[4].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.40, 0.40}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{0.60, 0.60}},
		{Position: [2]float64{1.0, 1.0}}}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "surface")
	if err != nil {
		t.Error(err)
	}
}

package surface

import (
	"testing"

	"github.com/ElleScotZ/project-cs/internal/algebra"
)

func TestSplineCurve(t *testing.T) {
	var (
		tspline TSpline
		cp      [5]ControlPoint
	)

	cp[0].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 0.0, 0.0}}
	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 0.0}}
	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 1.0, 0.0}}
	cp[3].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 1.0, 0.0}}
	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{0.50, 0.50, 0.50}}

	cp[0].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.50}},
		{Position: [2]float64{0.25, 1.00}},
		{Position: [2]float64{0.50, 1.0}}}

	cp[1].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.25, 0.50}},
		{Position: [2]float64{0.50, 1.00}},
		{Position: [2]float64{0.75, 1.0}}}

	cp[2].Knotvector = [5]Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.25, 0.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{0.75, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[3].Knotvector = [5]Knot{
		{Position: [2]float64{0.25, 0.0}},
		{Position: [2]float64{0.50, 0.0}},
		{Position: [2]float64{0.75, 0.50}},
		{Position: [2]float64{1.0, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	cp[4].Knotvector = [5]Knot{
		{Position: [2]float64{0.50, 0.0}},
		{Position: [2]float64{0.75, 0.0}},
		{Position: [2]float64{1.0, 0.50}},
		{Position: [2]float64{1.0, 1.00}},
		{Position: [2]float64{1.0, 1.0}}}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{1, 1}, "curve")
	if err != nil {
		t.Error(err)
	}
}

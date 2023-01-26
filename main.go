package main

import (
	"log"

	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/pkg/surface"
)

func saddle1a(cp [9]surface.ControlPoint) {
	var tspline surface.TSpline

	cp[0].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[1].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.0, 0.67}},
		{Position: [2]float64{0.33, 1.0}},
		{Position: [2]float64{0.67, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[2].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.0, 0.67}},
		{Position: [2]float64{0.33, 1.0}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[3].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[4].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
	}

	cp[5].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.0}},
		{Position: [2]float64{1.0, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
	}

	cp[6].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.0}},
		{Position: [2]float64{1.0, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[7].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[8].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle1a")
	if err != nil {
		log.Print(err)
	}
}

func saddle1b(cp [9]surface.ControlPoint) {
	var tspline surface.TSpline

	cp[0].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.0, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[1].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.0, 0.67}},
		{Position: [2]float64{0.33, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[2].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.67}},
		{Position: [2]float64{0.33, 1.0}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[3].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 1.0}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[4].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{1.0, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
	}

	cp[5].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.67, 0.0}},
		{Position: [2]float64{1.0, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
	}

	cp[6].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.0}},
		{Position: [2]float64{1.0, 0.33}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[7].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.0}},
		{Position: [2]float64{1.0, 0.67}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[8].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle1b")
	if err != nil {
		log.Print(err)
	}
}

func main() {
	// Creating control points for saddle
	var cp [9]surface.ControlPoint

	cp[0].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 0.0, 5.0}}
	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 5.0}}
	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 1.50, 0.0}}
	cp[3].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 3.50, 0.0}}
	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 5.0, 5.0}}
	cp[5].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 5.0, 5.0}}
	cp[6].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 3.50, 0.0}}
	cp[7].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 1.50, 0.0}}
	cp[8].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 2.50, 5.0}}

	saddle1a(cp)
	saddle1b(cp)
}

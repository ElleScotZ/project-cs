package main

import (
	"log"

	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/pkg/surface"
)

func main() {
	// Creating a saddle
	var (
		tspline surface.TSpline
		cp      [9]surface.ControlPoint
	)

	cp[0].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 0.0, 5.0}}
	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 5.0}}

	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 1.50, 0.0}}
	cp[3].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 3.50, 0.0}}

	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 1.5, 0.0}}
	cp[5].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 3.5, 0.0}}

	cp[6].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 5.0, 5.0}}
	cp[7].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 5.0, 5.0}}

	cp[8].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 2.50, 5.0}}

	cp[0].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 1.0 / 3.0}},
		{Position: [2]float64{1.0 / 3.0, 2.0 / 3.0}},
		{Position: [2]float64{2.0 / 3.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[1].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 1.0 / 3.0}},
		{Position: [2]float64{0.0, 2.0 / 3.0}},
		{Position: [2]float64{1.0 / 3.0, 1.0}},
		{Position: [2]float64{2.0 / 3.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[2].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 1.0 / 3.0}},
		{Position: [2]float64{0.0, 2.0 / 3.0}},
		{Position: [2]float64{1.0 / 3.0, 1.0}},
		{Position: [2]float64{2.0 / 3.0, 1.0}},
	}

	cp[3].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 1.0 / 3.0}},
		{Position: [2]float64{1.0 / 3.0, 2.0 / 3.0}},
		{Position: [2]float64{2.0 / 3.0, 1.0}},
	}

	cp[7].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{1.0 / 3.0, 0.0}},
		{Position: [2]float64{2.0 / 3.0, 1.0 / 3.0}},
		{Position: [2]float64{1.0, 2.0 / 3.0}},
	}

	cp[6].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{1.0 / 3.0, 0.0}},
		{Position: [2]float64{2.0 / 3.0, 0.0}},
		{Position: [2]float64{1.0, 1.0 / 3.0}},
		{Position: [2]float64{1.0, 2.0 / 3.0}},
	}

	cp[5].Knotvector = [5]surface.Knot{
		{Position: [2]float64{1.0 / 3.0, 0.0}},
		{Position: [2]float64{2.0 / 3.0, 0.0}},
		{Position: [2]float64{1.0, 1.0 / 3.0}},
		{Position: [2]float64{1.0, 2.0 / 3.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[4].Knotvector = [5]surface.Knot{
		{Position: [2]float64{1.0 / 3.0, 0.0}},
		{Position: [2]float64{2.0 / 3.0, 1.0 / 3.0}},
		{Position: [2]float64{1.0, 2.0 / 3.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[8].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{1.0 / 3.0, 1.0 / 3.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{2.0 / 3.0, 2.0 / 3.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle")
	if err != nil {
		log.Print(err)
	}
}
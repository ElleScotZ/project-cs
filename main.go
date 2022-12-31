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

	cp[8].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 2.50, 4.50}}

	// cp[9].Position = cp[0].Position.Add(cp[4].Position).Div(2.0)
	// //cp[9].Position.Coordinates[0] = cp[4].Position.Coordinates[0]
	// cp[10].Position = cp[1].Position.Add(cp[2].Position).Div(2.0)
	// //cp[10].Position.Coordinates[0] = cp[1].Position.Coordinates[0]
	// cp[11].Position = cp[3].Position.Add(cp[7].Position).Div(2.0)
	// //cp[11].Position.Coordinates[0] = cp[3].Position.Coordinates[0]
	// cp[12].Position = cp[6].Position.Add(cp[5].Position).Div(2.0)
	// //cp[12].Position.Coordinates[0] = cp[5].Position.Coordinates[0]

	// cp[13].Position = cp[9].Position.Add(cp[10].Position).Div(2.0)
	// cp[14].Position = cp[10].Position.Add(cp[11].Position).Div(2.0)
	// cp[15].Position = cp[11].Position.Add(cp[12].Position).Div(2.0)
	// cp[16].Position = cp[12].Position.Add(cp[9].Position).Div(2.0)

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

	// cp[9].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{0.0, 0.0}},
	// 	{Position: [2]float64{1.0, 1.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 4.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 4.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 4.0 / 3.0}},
	// }

	// cp[10].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{-1.0 / 3.0, 0.0}},
	// 	{Position: [2]float64{-1.0 / 3.0, 1.0}},
	// 	{Position: [2]float64{-1.0 / 3.0, 4.0 / 3.0}},
	// 	{Position: [2]float64{0.0, 4.0 / 3.0}},
	// 	{Position: [2]float64{1.0, 4.0 / 3.0}},
	// }

	// cp[11].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{-1.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{-1.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{-1.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{0.0, 0.0}},
	// 	{Position: [2]float64{1.0, 1.0}},
	// }

	// cp[12].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{0.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{1.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 0.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 1.0}},
	// }

	// cp[13].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{-1.0 / 3.0, 2.0 / 3.0}},
	// 	{Position: [2]float64{1.0 / 3.0, 1.0}},
	// 	{Position: [2]float64{0.50, 4.0 / 3.0}},
	// 	{Position: [2]float64{2.0 / 3.0, 4.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 4.0 / 3.0}},
	// }

	// cp[14].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{-1.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{-1.0 / 3.0, 1.0 / 3.0}},
	// 	{Position: [2]float64{-1.0 / 3.0, 0.50}},
	// 	{Position: [2]float64{0.0, 2.0 / 3.0}},
	// 	{Position: [2]float64{1.0 / 3.0, 4.0 / 3.0}},
	// }

	// cp[15].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{-1.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{1.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{0.50, -1.0 / 3.0}},
	// 	{Position: [2]float64{2.0 / 3.0, 0.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 1.0 / 3.0}},
	// }

	// cp[16].Knotvector = [5]surface.Knot{
	// 	{Position: [2]float64{2.0 / 3.0, -1.0 / 3.0}},
	// 	{Position: [2]float64{1.0, 1.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 0.50}},
	// 	{Position: [2]float64{4.0 / 3.0, 2.0 / 3.0}},
	// 	{Position: [2]float64{4.0 / 3.0, 4.0 / 3.0}},
	// }

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle")
	if err != nil {
		log.Print(err)
	}
}

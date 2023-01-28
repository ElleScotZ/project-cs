package main

import (
	"log"

	"github.com/ElleScotZ/project-cs/internal/algebra"
	"github.com/ElleScotZ/project-cs/internal/core"
	"github.com/ElleScotZ/project-cs/pkg/io"
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

func saddle2a(cp [9]surface.ControlPoint) {
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
		{Position: [2]float64{0.33, 0.33}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{0.67, 0.67}},
		{Position: [2]float64{1.0, 1.0}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle2a")
	if err != nil {
		log.Print(err)
	}
}

func saddle2b(cp [9]surface.ControlPoint) {
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

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle2b")
	if err != nil {
		log.Print(err)
	}
}

func saddle3a(cp [9]surface.ControlPoint) {
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

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle3a")
	if err != nil {
		log.Print(err)
	}
}

func saddle3b(cp [9]surface.ControlPoint) {
	var tspline surface.TSpline

	cp[0].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
		{Position: [2]float64{0.85, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[1].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.15, 0.67}},
		{Position: [2]float64{0.33, 1.0}},
		{Position: [2]float64{0.67, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[2].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.0, 0.67}},
		{Position: [2]float64{0.33, 0.85}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[3].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.15}},
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[4].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.15, 0.0}},
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
	}

	cp[5].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.0}},
		{Position: [2]float64{0.85, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
	}

	cp[6].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.15}},
		{Position: [2]float64{1.0, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
		{Position: [2]float64{1.0, 1.0}},
	}

	cp[7].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.67, 0.33}},
		{Position: [2]float64{1.0, 0.67}},
		{Position: [2]float64{1.0, 0.85}},
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

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle3b")
	if err != nil {
		log.Print(err)
	}
}

func saddle3c(cp [9]surface.ControlPoint) {
	var tspline surface.TSpline

	cp[0].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.15, 1.0}},
	}

	cp[1].Knotvector = [5]surface.Knot{
		{Position: [2]float64{-0.15, 0.33}},
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
		{Position: [2]float64{0.67, 1.15}},
	}

	cp[3].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, -0.15}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.33}},
		{Position: [2]float64{0.33, 0.67}},
		{Position: [2]float64{0.67, 1.0}},
	}

	cp[4].Knotvector = [5]surface.Knot{
		{Position: [2]float64{-0.15, 0.0}},
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
		{Position: [2]float64{1.15, 0.67}},
	}

	cp[6].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.33, -0.15}},
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
		{Position: [2]float64{1.0, 1.15}},
	}

	cp[8].Knotvector = [5]surface.Knot{
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.0, 0.0}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{1.0, 1.0}},
		{Position: [2]float64{1.0, 1.0}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle3c")
	if err != nil {
		log.Print(err)
	}
}

func saddle4a(cp [8]surface.ControlPoint) {
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

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle4a")
	if err != nil {
		log.Print(err)
	}
}

func saddle4b(cp [9]surface.ControlPoint) {
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
		{Position: [2]float64{0.0, 0.40}},
		{Position: [2]float64{0.20, 0.45}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{0.80, 0.55}},
		{Position: [2]float64{1.0, 0.60}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle4b")
	if err != nil {
		log.Print(err)
	}
}

func saddle4c(cp [9]surface.ControlPoint) {
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
		{Position: [2]float64{0.40, 0.0}},
		{Position: [2]float64{0.45, 0.20}},
		{Position: [2]float64{0.50, 0.50}},
		{Position: [2]float64{0.55, 0.80}},
		{Position: [2]float64{0.60, 1.0}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle4c")
	if err != nil {
		log.Print(err)
	}
}

func saddle5a(cp [8]surface.ControlPoint) {
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

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle5a")
	if err != nil {
		log.Print(err)
	}
}

func saddle5b(cp [9]surface.ControlPoint) {
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
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.33, 0.0}},
		{Position: [2]float64{0.50, 0.0}},
		{Position: [2]float64{0.67, 0.33}},
		{Position: [2]float64{0.67, 0.67}},
	}

	tspline.ControlPoints = append(tspline.ControlPoints, cp[:]...)

	err := tspline.GenerateSurface([2]int{50, 50}, "saddle5b")
	if err != nil {
		log.Print(err)
	}
}

func main() {
	// Creating control points for saddle
	var (
		mesh core.Mesh
		cp   [9]surface.ControlPoint
	)

	cp[0].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 0.0, 5.0}}
	cp[1].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 0.0, 5.0}}
	cp[2].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 1.50, 0.0}}
	cp[3].Position = algebra.Vector3D{Coordinates: [3]float64{2.50, 3.50, 0.0}}
	cp[4].Position = algebra.Vector3D{Coordinates: [3]float64{1.0, 5.0, 5.0}}
	cp[5].Position = algebra.Vector3D{Coordinates: [3]float64{-1.0, 5.0, 5.0}}
	cp[6].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 3.50, 0.0}}
	cp[7].Position = algebra.Vector3D{Coordinates: [3]float64{-2.50, 1.50, 0.0}}
	cp[8].Position = algebra.Vector3D{Coordinates: [3]float64{0.0, 2.50, 5.0}}

	for i := range cp {
		mesh.Vertices = append(mesh.Vertices, core.Vertex{
			Position: cp[i].Position,
			Color:    algebra.Vector3D{Coordinates: [3]float64{150, 0, 0}},
		})
	}

	io.ExportPLY(&mesh, "controlpoints")

	saddle1a(cp)
	saddle1b(cp)
	saddle2a(cp)
	saddle2b(cp)
	saddle3a(cp)
	saddle3b(cp)
	saddle3c(cp)
	saddle4a([8]surface.ControlPoint{cp[0], cp[1], cp[2], cp[3], cp[4], cp[5], cp[6], cp[7]})
	saddle4b(cp)
	saddle4c(cp)
	saddle5a([8]surface.ControlPoint{cp[0], cp[1], cp[2], cp[3], cp[4], cp[5], cp[6], cp[7]})
	saddle5b(cp)
}

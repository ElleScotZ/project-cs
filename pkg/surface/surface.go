package surface

import "github.com/ElleScotZ/project-cs/internal/algebra"

// Knot represents a knot in the parametrised 2D space on a T-mesh.
type Knot struct {
	Position [2]float64 // [0]: S/horizontal coordinate, [1]: T/vertical coordinate
}

// ControlPoint represents a control point on a T-Spline surface.
type ControlPoint struct {
	Position   algebra.Vector3D
	Knotvector [5]Knot
}

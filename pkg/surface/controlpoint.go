package surface

import "github.com/ElleScotZ/project-cs/internal/algebra"

// ControlPoint represents a control point on a T-Spline surface.
type ControlPoint struct {
	Position   algebra.Vector3D
	Knotvector [5]Knot
}

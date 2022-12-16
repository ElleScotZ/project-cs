package core

import (
	"github.com/ElleScotZ/project-cs/internal/algebra"
)

// Face struct represents a face in a mesh.
type Face struct {
	Vertices [3]int
}

// Mesh contains all parts of the mesh.
type Mesh struct {
	Vertices []*algebra.Vector3D
	Faces    []*Face
}

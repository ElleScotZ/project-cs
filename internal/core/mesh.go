package core

import (
	"github.com/ElleScotZ/project-cs/internal/algebra"
)

//
type Vertex struct {
	Position algebra.Vector3D
	Normal   algebra.Vector3D
	Color    algebra.Vector3D
}

// Face struct represents a face in a mesh.
type Face struct {
	Vertices [3]int
}

// Mesh contains all parts of the mesh.
type Mesh struct {
	Vertices []Vertex
	Faces    []Face
}

package core

import (
	"github.com/ElleScotZ/project-cs/internal/algebra"
)

// Vertex represents a vertex in a mesh.
type Vertex struct {
	Position algebra.Vector3D
	Normal   algebra.Vector3D
	Color    algebra.Vector3D
}

// Face represents a face in a mesh.
type Face struct {
	Vertices [3]int
}

// Mesh represents a 3D mesh.
type Mesh struct {
	Vertices []Vertex
	Faces    []Face
}

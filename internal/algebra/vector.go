package algebra

// Vector3D
type Vector3D struct {
	Coordinates [3]float64 // X, Y, Z coordinates
}

//
func (v Vector3D) Add(v2 Vector3D) Vector3D {
	return Vector3D{
		Coordinates: [3]float64{
			v.Coordinates[0] + v2.Coordinates[0],
			v.Coordinates[1] + v2.Coordinates[1],
			v.Coordinates[2] + v2.Coordinates[2],
		},
	}
}

//
func (v Vector3D) Sub(v2 Vector3D) Vector3D {
	return v.Add(v2.Mul(-1.0))
}

//
func (v Vector3D) Dot(v2 Vector3D) float64 {
	return v.Coordinates[0]*v2.Coordinates[0] +
		v.Coordinates[1]*v2.Coordinates[1] +
		v.Coordinates[2]*v2.Coordinates[2]
}

//
func (v Vector3D) Mul(scalar float64) Vector3D {
	return Vector3D{
		Coordinates: [3]float64{
			v.Coordinates[0] * scalar,
			v.Coordinates[1] * scalar,
			v.Coordinates[2] * scalar,
		},
	}
}

//
func (v Vector3D) Div(scalar float64) Vector3D {
	return v.Mul(1.0 / scalar)
}

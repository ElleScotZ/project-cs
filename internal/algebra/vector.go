package algebra

// Vector3D
type Vector3D struct {
	Position [3]float64 // X, Y, Z coordinates
}

//
func (v *Vector3D) Add(v2 Vector3D) Vector3D {
	return Vector3D{
		Position: [3]float64{
			v.Position[0] + v2.Position[0],
			v.Position[1] + v2.Position[1],
			v.Position[2] + v2.Position[2],
		},
	}
}

//
func (v *Vector3D) Mul(scalar float64) Vector3D {
	return Vector3D{
		Position: [3]float64{
			v.Position[0] * scalar,
			v.Position[1] * scalar,
			v.Position[2] * scalar,
		},
	}
}

//
func (v *Vector3D) Div(scalar float64) Vector3D {
	return v.Mul(1.0 / scalar)
}

package vec

import (
	"math"
	"math/rand"
)

// Vector3D type: basic 3D vector
type Vector3D struct {
	X, Y, Z float64
}

func (v Vector3D) MulScalar(s float64) Vector3D {
	return NewVector3D(v.X*s, v.Y*s, v.Z*s)
}

func (v Vector3D) FracScalar(s float64) Vector3D {
	return v.MulScalar(1 / s)
}

func (v1 Vector3D) Add(v2 Vector3D) Vector3D {
	return NewVector3D(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

func (v1 Vector3D) Sub(v2 Vector3D) Vector3D {
	return NewVector3D(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

func (v Vector3D) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3D) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v1 Vector3D) Dot(v2 Vector3D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector3D) Cross(v2 Vector3D) Vector3D {
	return Vector3D{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v Vector3D) Normalize() Vector3D {
	return v.FracScalar(v.Length())
}

// NewVector3D func: create new Vector3D
func NewVector3D(x, y, z float64) Vector3D {
	return Vector3D{x, y, z}
}

func NewZeroVector3D() Vector3D {
	return Vector3D{0, 0, 0}
}

func NewUnitVector3D() Vector3D {
	return Vector3D{1, 1, 1}
}

func NewRandVector3D() Vector3D {
	return Vector3D{rand.Float64(), rand.Float64(), rand.Float64()}
}

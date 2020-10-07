package vec

import (
	"math"
	"math/rand"
)

// Vec3 type: basic 3D vector
type Vec3 struct {
	x, y, z float64
}

func (v *Vec3) X() float64 {
	return v.x
}

func (v *Vec3) Y() float64 {
	return v.y
}

func (v *Vec3) Z() float64 {
	return v.z
}

func (v *Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) MulScalar(s float64) *Vec3 {
	v.x = v.x * s
	v.y = v.y * s
	v.z = v.z * s
	return v
}

func (v *Vec3) DivScalar(s float64) *Vec3 {
	v.MulScalar(1 / s)
	return v
}

func (v *Vec3) Add(v2 *Vec3) *Vec3 {
	v.x = v.x + v2.x
	v.y = v.y + v2.y
	v.z = v.z + v2.z
	return v
}

func (v *Vec3) Sub(v2 *Vec3) *Vec3 {
	v.x = v.x - v2.x
	v.y = v.y - v2.y
	v.z = v.z - v2.z
	return v
}

func (v *Vec3) Dot(v2 *Vec3) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vec3) Cross(v2 *Vec3) *Vec3 {
	v.x = v.y*v2.z - v.z*v2.y
	v.y = v.z*v2.x - v.x*v2.z
	v.z = v.x*v2.y - v.y*v2.x
	return v
}

func (v Vec3) GetNormal() Vec3 {
	v.DivScalar(v.Length())
	return v
}

// New func: create new Vec3
func New(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func NewZero() Vec3 {
	return Vec3{0, 0, 0}
}

func NewUnit() Vec3 {
	return Vec3{1, 1, 1}
}

func NewRand() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func (v *Vec3) Clip(min, max float64) *Vec3 {
	if v.x < min {
		v.x = min
	}
	if v.x > max {
		v.x = max
	}
	if v.y < min {
		v.y = min
	}
	if v.y > max {
		v.y = max
	}
	if v.z < min {
		v.z = min
	}
	if v.z > max {
		v.z = max
	}
	return v
}

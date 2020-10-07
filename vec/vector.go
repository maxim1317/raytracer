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
	newV := new(Vec3)
	newV.x = v.x * s
	newV.y = v.y * s
	newV.z = v.z * s
	return newV
}

func (v *Vec3) DivScalar(s float64) *Vec3 {
	return v.MulScalar(1 / s)
}

func (v *Vec3) Add(v2 *Vec3) *Vec3 {
	newV := new(Vec3)
	newV.x = v.x + v2.x
	newV.y = v.y + v2.y
	newV.z = v.z + v2.z
	return newV
}

func (v *Vec3) Sub(v2 *Vec3) *Vec3 {
	newV := new(Vec3)
	newV.x = v.x - v2.x
	newV.y = v.y - v2.y
	newV.z = v.z - v2.z
	return newV
}

func (v *Vec3) Dot(v2 *Vec3) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vec3) Cross(v2 *Vec3) *Vec3 {
	newV := new(Vec3)
	newV.x = v.y*v2.z - v.z*v2.y
	newV.y = v.z*v2.x - v.x*v2.z
	newV.z = v.x*v2.y - v.y*v2.x
	return newV
}

func (v Vec3) GetNormal() *Vec3 {
	return v.DivScalar(v.Length())
}

// New func: create new Vec3
func New(x, y, z float64) *Vec3 {
	vec := new(Vec3)
	vec.x = x
	vec.y = y
	vec.z = z
	return vec
}

func NewZero() *Vec3 {
	vec := new(Vec3)
	vec.x = 0
	vec.y = 0
	vec.z = 0
	return vec
}

func NewUnit() *Vec3 {
	vec := new(Vec3)
	vec.x = 1
	vec.y = 1
	vec.z = 1
	return vec
}

func NewRand() *Vec3 {
	vec := new(Vec3)
	vec.x = rand.Float64()
	vec.y = rand.Float64()
	vec.z = rand.Float64()
	return vec
}

func (v *Vec3) Clip(min, max float64) *Vec3 {
	newV := new(Vec3)
	if v.x < min {
		newV.x = min
	}
	if v.x > max {
		newV.x = max
	}
	if v.y < min {
		newV.y = min
	}
	if v.y > max {
		newV.y = max
	}
	if v.z < min {
		newV.z = min
	}
	if v.z > max {
		newV.z = max
	}
	return newV
}

package vec

import (
	"math"

	ut "github.com/maxim1317/raytracer/utils"
)

// Vec3 structure provides basic 3D vector
type Vec3 struct {
	x, y, z float64
}

// X getter
func (v *Vec3) X() float64 {
	return v.x
}

// Y getter
func (v *Vec3) Y() float64 {
	return v.y
}

// Z getter
func (v *Vec3) Z() float64 {
	return v.z
}

// LengthSquared returns vector squared length
func (v *Vec3) LengthSquared() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

// Length returns vector length
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// MulScalar multiplies vector by scalar
func (v *Vec3) MulScalar(s float64) *Vec3 {
	newV := new(Vec3)
	newV.x = v.x * s
	newV.y = v.y * s
	newV.z = v.z * s
	return newV
}

// DivScalar divides vector by scalar
func (v *Vec3) DivScalar(s float64) *Vec3 {
	return v.MulScalar(1 / s)
}

// Add adds another vector to the vector
func (v *Vec3) Add(v2 *Vec3) *Vec3 {
	newV := new(Vec3)
	newV.x = v.x + v2.x
	newV.y = v.y + v2.y
	newV.z = v.z + v2.z
	return newV
}

// Sub substracts another vector from the vector
func (v *Vec3) Sub(v2 *Vec3) *Vec3 {
	newV := new(Vec3)
	newV.x = v.x - v2.x
	newV.y = v.y - v2.y
	newV.z = v.z - v2.z
	return newV
}

// Dot returns dot product of two vectors
func (v *Vec3) Dot(v2 *Vec3) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

// Cross returns cross product of two vectors
func (v *Vec3) Cross(v2 *Vec3) *Vec3 {
	newV := new(Vec3)
	newV.x = v.y*v2.z - v.z*v2.y
	newV.y = v.z*v2.x - v.x*v2.z
	newV.z = v.x*v2.y - v.y*v2.x
	return newV
}

// GetNormal returns unit vector
func (v Vec3) GetNormal() *Vec3 {
	return v.DivScalar(v.Length())
}

// New creates new Vec3
func New(x, y, z float64) *Vec3 {
	vec := new(Vec3)
	vec.x = x
	vec.y = y
	vec.z = z
	return vec
}

// NewZero creates new zeroed Vec3
func NewZero() *Vec3 {
	vec := new(Vec3)
	vec.x = 0
	vec.y = 0
	vec.z = 0
	return vec
}

// NewUnit creates new unit Vec3
func NewUnit() *Vec3 {
	vec := new(Vec3)
	vec.x = 1
	vec.y = 1
	vec.z = 1
	return vec
}

// NewRand creates new random Vec3
func NewRand() *Vec3 {
	vec := new(Vec3)
	vec.x = ut.Rand()
	vec.y = ut.Rand()
	vec.z = ut.Rand()
	return vec
}

// NewRandInRange creates new random Vec3
func NewRandInRange(a, b float64) *Vec3 {
	vec := new(Vec3)
	vec.x = ut.RandRange(a, b)
	vec.y = ut.RandRange(a, b)
	vec.z = ut.RandRange(a, b)
	return vec
}

// NewRandInUnitSphere creates new random Vec3 in unit sphere
func NewRandInUnitSphere() *Vec3 {
	for {
		p := NewRandInRange(-1.0, 1.0)
		if p.LengthSquared() >= 1.0 {
			continue
		}
		return p
	}
}

// NewRandInUnitDisk creates new random Vec3 in unit disk
func NewRandInUnitDisk() *Vec3 {
	for {
		p := New(
			ut.RandRange(-1.0, 1.0),
			ut.RandRange(-1.0, 1.0),
			0.0,
		)
		if p.LengthSquared() >= 1.0 {
			continue
		}
		return p
	}
}

// NewRandUnit creates new random unit Vec3
func NewRandUnit() *Vec3 {
	v := new(Vec3)
	a := ut.RandRange(0, 2*math.Pi)
	z := ut.RandRange(-1.0, 1.0)
	r := math.Sqrt(1 - z*z)

	v.x = r * math.Cos(a)
	v.y = r * math.Sin(a)
	v.z = z
	return v
}

// Clip returns clipped vector
func (v *Vec3) Clip(min, max float64) *Vec3 {
	newV := new(Vec3)
	switch {
	case v.x < min:
		newV.x = min
	case v.x > max:
		newV.x = max
	default:
		newV.x = v.x
	}
	switch {
	case v.y < min:
		newV.y = min
	case v.y > max:
		newV.y = max
	default:
		newV.y = v.y
	}
	switch {
	case v.z < min:
		newV.z = min
	case v.z > max:
		newV.z = max
	default:
		newV.z = v.z
	}
	return newV
}

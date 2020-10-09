package vec

import (
	"math"

	ut "github.com/maxim1317/raytracer/utils"
)

// Vec3 structure provides basic 3D vector
type Vec3 struct {
	vec [3]float64
}

// X getter
func (v *Vec3) X() float64 {
	return v.vec[0]
}

// Y getter
func (v *Vec3) Y() float64 {
	return v.vec[1]
}

// Z getter
func (v *Vec3) Z() float64 {
	return v.vec[2]
}

// Ind getter
func (v *Vec3) Ind(i int) float64 {
	return v.vec[i]
}

// SetInd setter
func (v *Vec3) SetInd(i int, value float64) {
	v.vec[i] = value
}

// LengthSquared returns vector squared length
func (v *Vec3) LengthSquared() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

// Length returns vector length
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// MulScalar multiplies vector by scalar
func (v *Vec3) MulScalar(s float64) *Vec3 {
	var vec [3]float64
	vec[0] = v.X() * s
	vec[1] = v.Y() * s
	vec[2] = v.Z() * s

	return &Vec3{
		vec: vec,
	}
}

// DivScalar divides vector by scalar
func (v *Vec3) DivScalar(s float64) *Vec3 {
	return v.MulScalar(1 / s)
}

// Add adds another vector to the vector
func (v *Vec3) Add(v2 *Vec3) *Vec3 {
	var vec [3]float64
	vec[0] = v.X() + v2.X()
	vec[1] = v.Y() + v2.Y()
	vec[2] = v.Z() + v2.Z()

	return &Vec3{
		vec: vec,
	}
}

// Sub substracts another vector from the vector
func (v *Vec3) Sub(v2 *Vec3) *Vec3 {
	var vec [3]float64
	vec[0] = v.X() - v2.X()
	vec[1] = v.Y() - v2.Y()
	vec[2] = v.Z() - v2.Z()

	return &Vec3{
		vec: vec,
	}
}

// Dot returns dot product of two vectors
func (v *Vec3) Dot(v2 *Vec3) float64 {
	return v.X()*v2.X() + v.Y()*v2.Y() + v.Z()*v2.Z()
}

// Cross returns cross product of two vectors
func (v *Vec3) Cross(v2 *Vec3) *Vec3 {
	var vec [3]float64
	vec[0] = v.Y()*v2.Z() - v.Z()*v2.Y()
	vec[1] = v.Z()*v2.X() - v.X()*v2.Z()
	vec[2] = v.X()*v2.Y() - v.Y()*v2.X()

	return &Vec3{
		vec: vec,
	}
}

// GetNormal returns unit vector
func (v Vec3) GetNormal() *Vec3 {
	return v.DivScalar(v.Length())
}

// New creates new Vec3
func New(x, y, z float64) *Vec3 {
	var vec [3]float64
	vec[0] = x
	vec[1] = y
	vec[2] = z

	return &Vec3{
		vec: vec,
	}
}

// Copy creates Vec3
func (v *Vec3) Copy() *Vec3 {
	var vec [3]float64
	vec[0] = v.X()
	vec[1] = v.Y()
	vec[2] = v.Z()

	return &Vec3{
		vec: vec,
	}
}

// NewZero creates new zeroed Vec3
func NewZero() *Vec3 {
	var vec [3]float64
	vec[0] = 0
	vec[1] = 0
	vec[2] = 0

	return &Vec3{
		vec: vec,
	}
}

// NewUnit creates new unit Vec3
func NewUnit() *Vec3 {
	var vec [3]float64
	vec[0] = 1
	vec[1] = 1
	vec[2] = 1

	return &Vec3{
		vec: vec,
	}
}

// NewRand creates new random Vec3
func NewRand() *Vec3 {
	var vec [3]float64
	vec[0] = ut.Rand()
	vec[1] = ut.Rand()
	vec[2] = ut.Rand()

	return &Vec3{
		vec: vec,
	}
}

// NewRandInRange creates new random Vec3
func NewRandInRange(a, b float64) *Vec3 {
	var vec [3]float64
	vec[0] = ut.RandRange(a, b)
	vec[1] = ut.RandRange(a, b)
	vec[2] = ut.RandRange(a, b)

	return &Vec3{
		vec: vec,
	}
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
	var vec [3]float64

	a := ut.RandRange(0, 2*math.Pi)
	z := ut.RandRange(-1.0, 1.0)
	r := math.Sqrt(1 - z*z)

	vec[0] = r * math.Cos(a)
	vec[1] = r * math.Sin(a)
	vec[2] = z

	return &Vec3{
		vec: vec,
	}
}

// Clip returns clipped vector
func (v *Vec3) Clip(min, max float64) *Vec3 {
	return New(
		math.Min(math.Max(v.X(), min), max),
		math.Min(math.Max(v.Y(), min), max),
		math.Min(math.Max(v.Z(), min), max),
	)
}

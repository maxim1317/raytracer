package color

import (
	"math"
	"math/rand"

	"github.com/maxim1317/raytracer/vec"
)

const maxColor = 255.99

type Color struct {
	r, g, b float64
}

func (c *Color) R() float64 {
	return c.r
}

func (c *Color) G() float64 {
	return c.g
}

func (c *Color) B() float64 {
	return c.b
}

func (c *Color) FromVec3(v *vec.Vec3) *Color {
	c.r = v.X()
	c.g = v.Y()
	c.b = v.Z()
	return c
}

func (v *Color) LengthSquared() float64 {
	return v.r*v.r + v.g*v.g + v.b*v.b
}

func (v *Color) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Color) MulScalar(s float64) *Color {
	v.r = v.r * s
	v.g = v.g * s
	v.b = v.b * s
	return v
}

func (v *Color) DivScalar(s float64) *Color {
	v.MulScalar(1 / s)
	return v
}

func (v *Color) Add(v2 *Color) *Color {
	v.r = v.r + v2.r
	v.g = v.g + v2.g
	v.b = v.b + v2.b
	return v
}

func (v *Color) Sub(v2 *Color) *Color {
	v.r = v.r - v2.r
	v.g = v.g - v2.g
	v.b = v.b - v2.b
	return v
}

func (v *Color) Dot(v2 *Color) float64 {
	return v.r*v2.r + v.g*v2.g + v.b*v2.b
}

func (v *Color) Cross(v2 *Color) *Color {
	v.r = v.g*v2.b - v.b*v2.g
	v.g = v.b*v2.r - v.r*v2.b
	v.b = v.r*v2.g - v.g*v2.r
	return v
}

func (v Color) GetNormal() Color {
	v.DivScalar(v.Length())
	return v
}

// New func: create new Color
func New(x, y, z float64) Color {
	return Color{x, y, z}
}

func NewBlack() Color {
	return Color{0, 0, 0}
}

func NewWhite() Color {
	return Color{1, 1, 1}
}

func NewRand() Color {
	return Color{rand.Float64(), rand.Float64(), rand.Float64()}
}

func (v *Color) Clip(min, max float64) *Color {
	if v.r < min {
		v.r = min
	}
	if v.r > max {
		v.r = max
	}
	if v.g < min {
		v.g = min
	}
	if v.g > max {
		v.g = max
	}
	if v.b < min {
		v.b = min
	}
	if v.b > max {
		v.b = max
	}
	return v
}

package color

import (
	"math"

	ut "github.com/maxim1317/raytracer/utils"
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

func (c *Color) LengthSquared() float64 {
	return c.r*c.r + c.g*c.g + c.b*c.b
}

func (c *Color) Length() float64 {
	return math.Sqrt(c.LengthSquared())
}

func (c *Color) MulScalar(s float64) *Color {
	newC := new(Color)
	newC.r = c.r * s
	newC.g = c.g * s
	newC.b = c.b * s
	return newC
}

func (c *Color) DivScalar(s float64) *Color {
	return c.MulScalar(1 / s)
}

func (c *Color) Add(v2 *Color) *Color {
	newC := new(Color)
	newC.r = c.r + v2.r
	newC.g = c.g + v2.g
	newC.b = c.b + v2.b
	return newC
}

func (c *Color) Sub(v2 *Color) *Color {
	newC := new(Color)
	newC.r = c.r - v2.r
	newC.g = c.g - v2.g
	newC.b = c.b - v2.b
	return newC
}

func (c *Color) Dot(v2 *Color) float64 {
	return c.r*v2.r + c.g*v2.g + c.b*v2.b
}

func (c *Color) Cross(v2 *Color) *Color {
	newC := new(Color)
	newC.r = c.g*v2.b - c.b*v2.g
	newC.g = c.b*v2.r - c.r*v2.b
	newC.b = c.r*v2.g - c.g*v2.r
	return newC
}

func (c *Color) Gamma2() *Color {
	newC := new(Color)
	newC.r = math.Sqrt(c.r)
	newC.g = math.Sqrt(c.g)
	newC.b = math.Sqrt(c.b)
	return newC
}

// New func: create new Vec3
func New(r, g, b float64) *Color {
	c := new(Color)
	c.r = r
	c.g = g
	c.b = b
	return c
}

func Black() *Color {
	c := new(Color)
	c.r = 0
	c.g = 0
	c.b = 0
	return c
}

func White() *Color {
	c := new(Color)
	c.r = 1
	c.g = 1
	c.b = 1
	return c
}

func Rand() *Color {
	c := new(Color)
	c.r = ut.Rand()
	c.g = ut.Rand()
	c.b = ut.Rand()
	return c
}

func RandInRange(a, b float64) *Color {
	c := new(Color)
	c.r = ut.RandRange(a, b)
	c.g = ut.RandRange(a, b)
	c.b = ut.RandRange(a, b)
	return c
}

func (c *Color) Clip(min, max float64) *Color {
	newC := new(Color)
	switch {
	case c.r < min:
		newC.r = min
	case c.r > max:
		newC.r = max
	default:
		newC.r = c.r
	}
	switch {
	case c.g < min:
		newC.g = min
	case c.g > max:
		newC.g = max
	default:
		newC.g = c.g
	}
	switch {
	case c.b < min:
		newC.b = min
	case c.b > max:
		newC.b = max
	default:
		newC.b = c.b
	}
	return newC
}

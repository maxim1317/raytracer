package texture

import (
	"math"

	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type Checker struct {
	odd, even Texture
}

func NewCheckerTextured(t0, t1 *Texture) *Checker {
	return &Checker{
		odd:  *t0,
		even: *t1,
	}
}
func NewCheckerColored(c1, c2 *color.Color) *Checker {
	return &Checker{
		odd:  NewSolidColor(c1),
		even: NewSolidColor(c2),
	}
}

func (c *Checker) Value(u, v float64, p *vec.Vec3) *color.Color {
	sines := math.Sin(10*p.X()) * math.Sin(10*p.Y()) * math.Sin(10*p.Z())
	if sines < 0 {
		return c.odd.Value(u, v, p)
	} else {
		return c.even.Value(u, v, p)
	}
}

package texture

import (
	"math"

	"github.com/maxim1317/raytracer/color"
	n "github.com/maxim1317/raytracer/utils/noise"
	"github.com/maxim1317/raytracer/vec"
)

type NoiseTexture struct {
	noise *n.Perlin
	scale float64
}

func NewNoiseTexture(scale float64) *NoiseTexture {
	return &NoiseTexture{
		scale: scale,
		noise: n.NewPerlin(),
	}
}

func (c *NoiseTexture) Value(u, v float64, p *vec.Vec3) *color.Color {
	return color.New(1, 1, 1).MulScalar(0.5).MulScalar(1 + math.Sin(c.scale*p.Z()+10*c.noise.Turb(p, 7)))
}

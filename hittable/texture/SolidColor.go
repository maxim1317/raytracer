package texture

import (
	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type SolidColor struct {
	color *color.Color
}

func (s *SolidColor) Value(u, v float64, p *vec.Vec3) *color.Color {
	return s.color
}

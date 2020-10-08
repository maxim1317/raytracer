package texture

import (
	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

// Texture interface
type Texture interface {
	Value(u, v float64, p *vec.Vec3) *color.Color
}

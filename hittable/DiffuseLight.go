package hittable

import (
	c "github.com/maxim1317/raytracer/color"
	t "github.com/maxim1317/raytracer/hittable/texture"
	"github.com/maxim1317/raytracer/vec"
)

type DiffuseLight struct {
	emit t.Texture
}

func NewDiffuseLightTextured(emit t.Texture) DiffuseLight {
	l := new(DiffuseLight)
	l.emit = emit
	return *l
}

func NewDiffuseLightColored(emit *c.Color) DiffuseLight {
	l := new(DiffuseLight)
	l.emit = t.NewSolidColor(emit)
	return *l
}

func (l DiffuseLight) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color) {
	return false, scattered, attenuation
}

func (l DiffuseLight) Emitted(u, v float64, p *vec.Vec3) *c.Color {
	return l.emit.Value(u, v, p)
}

package hittable

import (
	"github.com/maxim1317/raytracer/color"
	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/hittable/texture"
	"github.com/maxim1317/raytracer/vec"
)

type Isotropic struct {
	albedo texture.Texture
}

func NewIsotropicColored(c *color.Color) *Isotropic {
	return &Isotropic{
		albedo: texture.NewSolidColor(c),
	}
}

func NewIsotropicTextured(t *texture.Texture) *Isotropic {
	return &Isotropic{
		albedo: *t,
	}
}

func (iso *Isotropic) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color) {
	scattered = vec.NewRay(rec.P, vec.NewRandInUnitSphere(), rIn.Time())
	attenuation = iso.albedo.Value(rec.U, rec.V, rec.P)
	return true, scattered, attenuation
}

func (iso *Isotropic) Emitted(u, v float64, p *vec.Vec3) *c.Color {
	return c.Black()
}

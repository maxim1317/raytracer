package hittable

import (
	c "github.com/maxim1317/raytracer/color"
	t "github.com/maxim1317/raytracer/hittable/texture"
	"github.com/maxim1317/raytracer/vec"
)

type Lambertian struct {
	albedo t.Texture
}

func NewLambertianTextured(albedo t.Texture) Lambertian {
	l := new(Lambertian)
	l.albedo = albedo
	return *l
}

func NewLambertianColored(albedo *c.Color) Lambertian {
	l := new(Lambertian)
	l.albedo = t.NewSolidColor(albedo)
	return *l
}

func (l Lambertian) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color) {
	scatterDirection := rec.Normal.Add(vec.NewRandUnit())
	scattered = vec.NewRay(rec.P, scatterDirection, rIn.Time())
	attenuation = l.albedo.Value(rec.U, rec.V, rec.P)
	return true, scattered, attenuation
}

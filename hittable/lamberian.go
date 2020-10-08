package hittable

import (
	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type Lambertian struct {
	albedo *c.Color
}

func NewLambertian(albedo *c.Color) Lambertian {
	l := new(Lambertian)
	l.albedo = albedo
	return *l
}

func (l Lambertian) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color) {
	scatterDirection := rec.Normal.Add(vec.NewRandUnit())
	scattered = vec.NewRay(rec.P, scatterDirection, rIn.Time())
	attenuation = l.albedo
	return true, scattered, attenuation
}

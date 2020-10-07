package hittable

import (
	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type Metal struct {
	Albedo *c.Color
}

func NewMetal(albedo *c.Color) Metal {
	m := new(Metal)
	m.Albedo = albedo
	return *m
}

func (m Metal) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) bool {
	reflected := Reflect(rIn.Dir, rec.Normal)
	scattered = vec.NewRay(rec.P, reflected)
	attenuation = m.Albedo
	return scattered.Dir.Dot(rec.Normal) > 0
}

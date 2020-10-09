package hittable

import (
	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type Metal struct {
	albedo *c.Color
	fuzz   float64
}

func NewMetal(albedo *c.Color, fuzz float64) Metal {
	m := new(Metal)
	m.albedo = albedo
	if fuzz < 1 {
		m.fuzz = fuzz
	} else {
		m.fuzz = 1
	}
	return *m
}

func (m Metal) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color) {
	reflected := Reflect(rIn.Direction(), rec.Normal)
	scattered = vec.NewRay(rec.P, reflected.Add(vec.NewRandInUnitSphere().MulScalar(m.fuzz)), rIn.Time())
	attenuation = m.albedo
	return scattered.Direction().Dot(rec.Normal) > 0, scattered, attenuation
}

func (m Metal) Emitted(u, v float64, p *vec.Vec3) *c.Color {
	return c.Black()
}

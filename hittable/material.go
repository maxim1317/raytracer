package hittable

import (
	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type Material interface {
	Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) bool
}

func Reflect(v, n *vec.Vec3) *vec.Vec3 {
	return v.Sub(n.MulScalar(2 * v.Dot(n)))
}

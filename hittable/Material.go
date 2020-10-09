package hittable

import (
	"math"

	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/vec"
)

type Material interface {
	Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color)
	Emitted(u, v float64, p *vec.Vec3) *c.Color
}

func Reflect(v, n *vec.Vec3) *vec.Vec3 {
	return v.Sub(n.MulScalar(2 * v.Dot(n)))
}

func Refract(uv, n *vec.Vec3, etaiOverEtat float64) *vec.Vec3 {
	cosTheta := vec.NewZero().Sub(uv).Dot(n)
	rOutPerp := uv.Add(n.MulScalar(cosTheta)).MulScalar(etaiOverEtat)
	rOutParallel := n.MulScalar(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

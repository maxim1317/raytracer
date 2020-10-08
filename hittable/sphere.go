package hittable

import (
	"math"

	"github.com/maxim1317/raytracer/vec"
)

type Sphere struct {
	center *vec.Vec3
	radius float64
	Material
}

func NewSphere(center *vec.Vec3, radius float64, mat Material) *Sphere {
	return &Sphere{
		center:   center,
		radius:   radius,
		Material: mat,
	}
}
func (s Sphere) Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool {

	// t^2*b*b + 2*t*b*(A−C) + (A−C)*(A−C) − r^2 = 0

	oc := r.Origin().Sub(s.center)

	var a, halfb, c, discriminant float64

	a = r.Direction().LengthSquared()
	halfb = oc.Dot(r.Direction())
	c = oc.LengthSquared() - s.radius*s.radius
	discriminant = halfb*halfb - a*c

	if discriminant > 0 {
		var root = math.Sqrt(discriminant)

		var temp = (-halfb - root) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.At(rec.T)
			outwardNormal := rec.P.Sub(s.center).DivScalar(s.radius)
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Material
			return true
		}

		temp = (-halfb + root) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.At(rec.T)
			outwardNormal := rec.P.Sub(s.center).DivScalar(s.radius)
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Material
			return true
		}
	}

	return false
}

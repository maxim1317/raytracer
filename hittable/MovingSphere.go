package hittable

import (
	"math"

	"github.com/maxim1317/raytracer/vec"
)

type MovingSphere struct {
	center0, center1 *vec.Vec3
	time0, time1     float64
	radius           float64
	mat              Material
}

func NewMovingSphere(
	center0, center1 *vec.Vec3,
	time0, time1 float64,
	radius float64,
	mat Material,
) *MovingSphere {
	return &MovingSphere{
		center0: center0,
		center1: center1,
		time0:   time0,
		time1:   time1,
		radius:  radius,
		mat:     mat,
	}
}

func (s *MovingSphere) Center(time float64) *vec.Vec3 {
	return s.center0.Add((s.center1.Sub(s.center0)).MulScalar((time - s.time0) / (s.time1 - s.time0)))
}

func (s *MovingSphere) Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool {

	// t^2*b*b + 2*t*b*(A−C) + (A−C)*(A−C) − r^2 = 0

	oc := r.Origin().Sub(s.Center(r.Time()))

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
			outwardNormal := rec.P.Sub(s.Center(r.Time())).DivScalar(s.radius)
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.mat
			return true
		}

		temp = (-halfb + root) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.At(rec.T)
			outwardNormal := rec.P.Sub(s.Center(r.Time())).DivScalar(s.radius)
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.mat
			return true
		}
	}

	return false
}

package hittable

import (
	"math"

	"github.com/maxim1317/raytracer/vec"
)

type Sphere struct {
	center *vec.Vec3
	radius float64
	mat    Material
}

func NewSphere(center *vec.Vec3, radius float64, mat Material) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
		mat:    mat,
	}
}

func (s *Sphere) Center() *vec.Vec3 {
	return s.center
}

func (s *Sphere) Radius() float64 {
	return s.radius
}

func (s *Sphere) Mat() Material {
	return s.mat
}

func (s *Sphere) Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool {

	// t^2*b*b + 2*t*b*(A−C) + (A−C)*(A−C) − r^2 = 0

	oc := r.Origin().Sub(s.Center())

	var a, halfb, c, discriminant float64

	a = r.Direction().LengthSquared()
	halfb = oc.Dot(r.Direction())
	c = oc.LengthSquared() - s.Radius()*s.Radius()
	discriminant = halfb*halfb - a*c

	if discriminant > 0 {
		var root = math.Sqrt(discriminant)

		var temp = (-halfb - root) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.At(rec.T)
			outwardNormal := rec.P.Sub(s.Center()).DivScalar(s.Radius())
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Mat()
			return true
		}

		temp = (-halfb + root) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.P = r.At(rec.T)
			outwardNormal := rec.P.Sub(s.Center()).DivScalar(s.Radius())
			rec.SetFaceNormal(r, outwardNormal)
			rec.Mat = s.Mat()
			return true
		}
	}

	return false
}

func (s *Sphere) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	outputBox = NewAABB(
		s.Center().Sub(vec.New(s.Radius(), s.Radius(), s.Radius())),
		s.Center().Add(vec.New(s.Radius(), s.Radius(), s.Radius())),
	)
	return true, outputBox
}

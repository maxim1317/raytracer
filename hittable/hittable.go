package hittable

import (
	"github.com/maxim1317/raytracer/vec"
)

type HitRecord struct {
	P, Normal *vec.Vec3
	T         float64
	Mat       Material
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r *vec.Ray, outwardNormal *vec.Vec3) {
	h.FrontFace = r.Dir.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		zero := vec.NewZero()
		h.Normal = zero.Sub(outwardNormal)
	}
}

type Hittable interface {
	Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool
}

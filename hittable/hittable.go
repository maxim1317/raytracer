package hittable

import (
	"github.com/maxim1317/raytracer/vec"
)

type HitRecord struct {
	P, Normal vec.Vector3D
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r *vec.Ray, outwardNormal vec.Vector3D) {
	h.FrontFace = r.Dir.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = vec.NewZeroVector3D().Sub(outwardNormal)
	}
}

type Hittable interface {
	Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool
}

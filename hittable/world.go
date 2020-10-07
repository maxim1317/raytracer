package hittable

import (
	"github.com/maxim1317/raytracer/vec"
)

type World struct {
	Elements []Hittable
}

func (w *World) Add(h Hittable) {
	w.Elements = append(w.Elements, h)
}

func (w *World) Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool {
	hitAnything := false
	closest := tMax

	for _, element := range w.Elements {
		hit := element.Hit(r, tMin, closest, rec)

		if hit {
			hitAnything = true
			closest = rec.T
		}
	}
	return hitAnything
}

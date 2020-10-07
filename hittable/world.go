package hittable

import (
	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type World struct {
	Elements []*Hittable
}

func (w *World) Add(h Hittable) {
	w.Elements = append(w.Elements, &h)
}

func (w *World) Hit(r *vec.Ray, tMin, tMax float64, rec *HitRecord) bool {
	hitAnything := false
	closest := tMax

	for _, element := range w.Elements {
		hit := (*element).Hit(r, tMin, closest, rec)

		if hit {
			hitAnything = true
			closest = rec.T
		}
	}
	return hitAnything
}

func RandomWorld() *World {
	world := new(World)

	groundMaterial := NewLambertian(color.New(0.5, 0.5, 0.5))
	world.Add(NewSphere(vec.New(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := utils.Rand()
			center := vec.New(float64(a)+0.9*utils.Rand(), 0.2, float64(b)+0.9*utils.Rand())

			if (center.Sub(vec.New(4, 0.2, 0))).Length() > 0.9 {
				var sphereMaterial Material
				switch {
				case chooseMat < 0.8:
					// diffuse
					albedo := color.Rand().Mul(color.Rand())
					sphereMaterial = NewLambertian(albedo)
					world.Add(NewSphere(center, 0.2, sphereMaterial))
				case chooseMat < 0.95:
					// metal
					albedo := color.RandInRange(0.5, 1)
					fuzz := utils.RandRange(0, 0.5)
					sphereMaterial = NewMetal(albedo, fuzz)
					world.Add(NewSphere(center, 0.2, sphereMaterial))
				default:
					// glass
					sphereMaterial = NewDielectric(1.5)
					world.Add(NewSphere(center, 0.2, sphereMaterial))

				}
			}
		}
	}

	mat1 := NewDielectric(1.5)
	world.Add(NewSphere(vec.New(0, 1, 0), 1.0, mat1))

	mat2 := NewLambertian(color.New(0.4, 0.2, 0.1))
	world.Add(NewSphere(vec.New(-4, 1, 0), 1.0, mat2))

	mat3 := NewMetal(color.New(0.7, 0.6, 0.5), 0.0)
	world.Add(NewSphere(vec.New(4, 1, 0), 1.0, mat3))

	return world
}

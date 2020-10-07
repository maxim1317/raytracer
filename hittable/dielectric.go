package hittable

import (
	"math"

	c "github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type Dielectric struct {
	ir float64
}

func NewDielectric(ir float64) *Dielectric {
	d := Dielectric{
		ir: ir,
	}
	return &d
}

func reflectance(cosine, refIdx float64) float64 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

func (d Dielectric) Scatter(rIn *vec.Ray, rec *HitRecord, attenuation *c.Color, scattered *vec.Ray) (bool, *vec.Ray, *c.Color) {

	attenuation = c.New(1.0, 1.0, 1.0)
	var refractionRatio float64
	if rec.FrontFace {
		refractionRatio = (1.0 / d.ir)
	} else {
		refractionRatio = d.ir
	}

	unitDirection := rIn.Dir.GetNormal()
	cosTheta := math.Min(vec.NewZero().Sub(unitDirection).Dot(rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0
	var direction *vec.Vec3
	if cannotRefract || reflectance(cosTheta, refractionRatio) > utils.Rand() {
		direction = Reflect(unitDirection, rec.Normal)
	} else {
		direction = Refract(unitDirection, rec.Normal, refractionRatio)
	}

	scattered = vec.NewRay(rec.P, direction)
	return true, scattered, attenuation
}

package hittable

import (
	"fmt"
	"math"

	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/hittable/texture"
	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type ConstantMedium struct {
	boundary      Hittable
	phaseFunction Material
	negInvDensity float64
}

func NewConstantMediumTextured(b Hittable, d float64, a *texture.Texture) *ConstantMedium {
	return &ConstantMedium{
		boundary:      b,
		negInvDensity: -1.0 / d,
		phaseFunction: NewIsotropicTextured(a),
	}
}

func NewConstantMediumColored(b Hittable, d float64, c *color.Color) *ConstantMedium {
	return &ConstantMedium{
		boundary:      b,
		negInvDensity: -1.0 / d,
		phaseFunction: NewIsotropicColored(c),
	}
}

func (c *ConstantMedium) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	// Print occasional samples when debugging. To enable, set enableDebug true.
	enableDebug := false
	debugging := enableDebug && utils.Rand() < 0.00001

	rec1 := new(HitRecord)
	rec2 := new(HitRecord)

	infinity := math.MaxFloat64

	hit1, rec1 := c.boundary.Hit(r, -infinity, infinity, rec1)

	if !hit1 {
		return false, rec
	}

	hit2, rec2 := c.boundary.Hit(r, rec1.T+0.0001, infinity, rec2)
	if !hit2 {
		return false, rec
	}

	if debugging {
		fmt.Printf("\nt0 = %v, t1 = %v\n", rec1.T, rec2.T)
	}

	if rec1.T < t0 {
		rec1.T = t0
	}
	if rec2.T > t1 {
		rec2.T = t1
	}

	if rec1.T >= rec2.T {
		return false, rec
	}

	if rec1.T < 0 {
		rec1.T = 0
	}

	rayLength := r.Direction().Length()
	distanceInsideBoundary := (rec2.T - rec1.T) * rayLength
	hitDistance := c.negInvDensity * math.Log(utils.Rand())

	if hitDistance > distanceInsideBoundary {
		return false, rec
	}

	rec.T = rec1.T + hitDistance/rayLength
	rec.P = r.At(rec.T)

	if debugging {
		fmt.Printf("hitDistance = %v\nrec.T = %v\nrec.P = %v\n", hitDistance, rec.T, rec.P)
	}

	rec.Normal = vec.New(1, 0, 0) // arbitrary
	rec.FrontFace = true          // also arbitrary
	rec.Mat = c.phaseFunction

	return true, rec
}

func (c *ConstantMedium) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	return c.boundary.BoundingBox(t0, t1, outputBox)
}

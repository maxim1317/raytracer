package hittable

import (
	"math"

	"github.com/maxim1317/raytracer/vec"
)

type AxisAlignedBoundingBox struct {
	max, min *vec.Vec3
}

func (b *AxisAlignedBoundingBox) Max() *vec.Vec3 {
	return b.Max()
}

func (b *AxisAlignedBoundingBox) Min() *vec.Vec3 {
	return b.Min()
}

func NewAABB(min, max *vec.Vec3) *AxisAlignedBoundingBox {
	return &AxisAlignedBoundingBox{
		min: min,
		max: max,
	}
}

func (b *AxisAlignedBoundingBox) Hit(r *vec.Ray, tMin, tMax float64) bool {
	var t0, t1, invD float64
	invD = 1.0 / r.Direction().X()
	t0 = (b.Min().X() - r.Origin().X()) * invD
	t1 = (b.Max().X() - r.Origin().X()) * invD
	if invD < 0.0 {
		t0, t1 = t1, t0
	}
	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)
	if tMax <= tMin {
		return false
	}

	invD = 1.0 / r.Direction().Y()
	t0 = (b.Min().Y() - r.Origin().Y()) * invD
	t1 = (b.Max().Y() - r.Origin().Y()) * invD
	if invD < 0.0 {
		t0, t1 = t1, t0
	}
	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)
	if tMax <= tMin {
		return false
	}

	invD = 1.0 / r.Direction().Z()
	t0 = (b.Min().Z() - r.Origin().Z()) * invD
	t1 = (b.Max().Z() - r.Origin().Z()) * invD
	if invD < 0.0 {
		t0, t1 = t1, t0
	}
	tMin = math.Max(t0, tMin)
	tMax = math.Min(t1, tMax)
	if tMax <= tMin {
		return false
	}
	return true
}

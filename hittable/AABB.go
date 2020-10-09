package hittable

import (
	"math"

	"github.com/maxim1317/raytracer/vec"
)

type AABB struct {
	max, min *vec.Vec3
}

func (b *AABB) Max() *vec.Vec3 {
	return b.max
}

func (b *AABB) Min() *vec.Vec3 {
	return b.min
}

func NewAABB(min, max *vec.Vec3) *AABB {
	return &AABB{
		min: min,
		max: max,
	}
}

func (b AABB) Hit(r *vec.Ray, t0, t1 float64) bool {
	for i := 0; i < 3; i++ {
		invD := 1.0 / r.Direction().X()
		tMin := (b.Min().Ind(i) - r.Origin().Ind(i)) * invD
		tMax := (b.Max().Ind(i) - r.Origin().Ind(i)) * invD
		if invD < 0.0 {
			t0, t1 = t1, t0
		}
		t0 = math.Max(tMin, t0)
		t1 = math.Min(tMax, t1)
		if t1 <= t0 {
			return false
		}

	}

	return true
}

func SurroundingBox(box0, box1 *AABB) *AABB {
	var small, big *vec.Vec3
	small = vec.New(
		math.Min(box0.Min().X(), box1.Min().X()),
		math.Min(box0.Min().Y(), box1.Min().Y()),
		math.Min(box0.Min().Z(), box1.Min().Z()),
	)

	big = vec.New(
		math.Max(box0.Max().X(), box1.Max().X()),
		math.Max(box0.Max().Y(), box1.Max().Y()),
		math.Max(box0.Max().Z(), box1.Max().Z()),
	)

	return NewAABB(small, big)
}

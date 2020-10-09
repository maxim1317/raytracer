package hittable

import (
	"github.com/maxim1317/raytracer/vec"
)

type Box struct {
	boxMin, boxMax *vec.Vec3
	sides          HittableList
}

func NewBox(p0, p1 *vec.Vec3, mat Material) *Box {
	b := new(Box)

	b.boxMin = p0
	b.boxMax = p1

	side1 := NewXYRect(p0.X(), p1.X(), p0.Y(), p1.Y(), p1.Z(), mat)
	b.sides.Add(side1)
	side2 := NewXYRect(p0.X(), p1.X(), p0.Y(), p1.Y(), p0.Z(), mat)
	b.sides.Add(side2)

	side3 := NewXZRect(p0.X(), p1.X(), p0.Z(), p1.Z(), p1.Y(), mat)
	b.sides.Add(side3)
	side4 := NewXZRect(p0.X(), p1.X(), p0.Z(), p1.Z(), p0.Y(), mat)
	b.sides.Add(side4)

	side5 := NewYZRect(p0.Y(), p1.Y(), p0.Z(), p1.Z(), p1.X(), mat)
	b.sides.Add(side5)
	side6 := NewYZRect(p0.Y(), p1.Y(), p0.Z(), p1.Z(), p0.X(), mat)
	b.sides.Add(side6)

	return b
}

func (b *Box) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	return b.sides.Hit(r, t0, t1, rec)
}

func (b *Box) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	// The bounding box must have non-zero width in each dimension, so pad the Z
	// dimension a small amount.
	outputBox = NewAABB(b.boxMin, b.boxMax)
	return true, outputBox
}

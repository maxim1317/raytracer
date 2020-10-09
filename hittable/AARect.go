package hittable

import (
	"github.com/maxim1317/raytracer/vec"
)

type XYRect struct {
	x0, x1, y0, y1, k float64
	mat               Material
}

func NewXYRect(x0, x1, y0, y1, k float64, mat Material) *XYRect {
	return &XYRect{
		x0: x0, x1: x1,
		y0: y0, y1: y1,
		k:   k,
		mat: mat,
	}
}

func (rect *XYRect) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	t := (rect.k - r.Origin().Z()) / r.Direction().Z()
	if t < t0 || t > t1 {
		return false, rec
	}
	x := r.Origin().X() + t*r.Direction().X()
	y := r.Origin().Y() + t*r.Direction().Y()
	if x < rect.x0 || x > rect.x1 || y < rect.y0 || y > rect.y1 {
		return false, rec
	}
	rec.U = (x - rect.x0) / (rect.x1 - rect.x0)
	rec.V = (y - rect.y0) / (rect.y1 - rect.y0)
	rec.T = t
	outwardNormal := vec.New(0, 0, 1)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = rect.mat
	rec.P = r.At(t)
	return true, rec
}

func (rect *XYRect) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	// The bounding box must have non-zero width in each dimension, so pad the Z
	// dimension a small amount.
	outputBox = NewAABB(vec.New(rect.x0, rect.y0, rect.k-0.0001), vec.New(rect.x1, rect.y1, rect.k+0.0001))
	return true, outputBox
}

type YZRect struct {
	y0, y1, z0, z1, k float64
	mat               Material
}

func NewYZRect(y0, y1, z0, z1, k float64, mat Material) *YZRect {
	return &YZRect{
		y0: y0, y1: y1,
		z0: z0, z1: z1,
		k:   k,
		mat: mat,
	}
}

func (rect *YZRect) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	t := (rect.k - r.Origin().X()) / r.Direction().X()
	if t < t0 || t > t1 {
		return false, rec
	}
	y := r.Origin().Y() + t*r.Direction().Y()
	z := r.Origin().Z() + t*r.Direction().Z()
	if y < rect.y0 || y > rect.y1 || z < rect.z0 || z > rect.z1 {
		return false, rec
	}
	rec.U = (y - rect.y0) / (rect.y1 - rect.y0)
	rec.V = (z - rect.z0) / (rect.z1 - rect.z0)
	rec.T = t
	outwardNormal := vec.New(1, 0, 0)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = rect.mat
	rec.P = r.At(t)
	return true, rec
}

func (rect *YZRect) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	// The bounding boy must have non-zero width in each dimension, so pad the Z
	// dimension a small amount.
	outputBox = NewAABB(vec.New(rect.k-0.0001, rect.y0, rect.z0), vec.New(rect.k+0.0001, rect.y1, rect.z1))
	return true, outputBox
}

type XZRect struct {
	x0, x1, z0, z1, k float64
	mat               Material
}

func NewXZRect(x0, x1, z0, z1, k float64, mat Material) *XZRect {
	return &XZRect{
		x0: x0, x1: x1,
		z0: z0, z1: z1,
		k:   k,
		mat: mat,
	}
}

func (rect *XZRect) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	t := (rect.k - r.Origin().Y()) / r.Direction().Y()
	if t < t0 || t > t1 {
		return false, rec
	}
	x := r.Origin().X() + t*r.Direction().X()
	z := r.Origin().Z() + t*r.Direction().Z()
	if x < rect.x0 || x > rect.x1 || z < rect.z0 || z > rect.z1 {
		return false, rec
	}
	rec.U = (x - rect.x0) / (rect.x1 - rect.x0)
	rec.V = (z - rect.z0) / (rect.z1 - rect.z0)
	rec.T = t
	outwardNormal := vec.New(0, 1, 0)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Mat = rect.mat
	rec.P = r.At(t)
	return true, rec
}

func (rect *XZRect) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	// The bounding box must have non-zero width in each dimension, so pad the Z
	// dimension a small amount.
	outputBox = NewAABB(vec.New(rect.x0, rect.k-0.0001, rect.z0), vec.New(rect.x1, rect.k+0.0001, rect.z1))
	return true, outputBox
}

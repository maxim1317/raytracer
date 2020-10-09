package hittable

import (
	"math"

	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type HitRecord struct {
	P, Normal *vec.Vec3
	T, U, V   float64
	Mat       Material
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r *vec.Ray, outwardNormal *vec.Vec3) {
	h.FrontFace = r.Direction().Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		zero := vec.NewZero()
		h.Normal = zero.Sub(outwardNormal)
	}
}

type Hittable interface {
	Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord)
	BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB)
}

type RotateY struct {
	ptr                Hittable
	sinTheta, cosTheta float64
	hasbox             bool
	bbox               *AABB
}

func NewRotateY(ptr Hittable, angle float64) *RotateY {
	rot := new(RotateY)
	rot.ptr = ptr

	radians := utils.Degrees2Rad(angle)
	rot.sinTheta = math.Sin(radians)
	rot.cosTheta = math.Cos(radians)
	rot.hasbox, rot.bbox = ptr.BoundingBox(0, 1, rot.bbox)

	infinity := math.MaxFloat64

	min := vec.New(infinity, infinity, infinity)
	max := vec.New(-infinity, -infinity, -infinity)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				x := float64(i)*rot.bbox.Max().X() + float64(1-i)*rot.bbox.Min().X()
				y := float64(j)*rot.bbox.Max().Y() + float64(1-j)*rot.bbox.Min().Y()
				z := float64(k)*rot.bbox.Max().Z() + float64(1-k)*rot.bbox.Min().Z()

				newx := rot.cosTheta*x + rot.sinTheta*z
				newz := -rot.sinTheta*x + rot.cosTheta*z

				tester := vec.New(newx, y, newz)

				for c := 0; c < 3; c++ {
					min.SetInd(c, math.Min(min.Ind(c), tester.Ind(c)))
					max.SetInd(c, math.Max(max.Ind(c), tester.Ind(c)))
				}
			}
		}
	}

	rot.bbox = NewAABB(min, max)
	return rot
}

func (r *RotateY) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	return r.hasbox, r.bbox
}

func (rot *RotateY) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	origin := r.Origin()
	direction := r.Direction()

	origin.SetInd(0, rot.cosTheta*r.Origin().Ind(0)-rot.sinTheta*r.Origin().Ind(2))
	origin.SetInd(2, rot.sinTheta*r.Origin().Ind(0)+rot.cosTheta*r.Origin().Ind(2))

	direction.SetInd(0, rot.cosTheta*r.Direction().Ind(0)-rot.sinTheta*r.Direction().Ind(2))
	direction.SetInd(2, rot.sinTheta*r.Direction().Ind(0)+rot.cosTheta*r.Direction().Ind(2))

	rotatedR := vec.NewRay(origin, direction, r.Time())

	hit, rec := rot.ptr.Hit(rotatedR, t0, t1, rec)

	if !hit {
		return false, rec
	}

	p := rec.P.Copy()
	normal := rec.Normal.Copy()

	p.SetInd(0, rot.cosTheta*rec.P.Ind(0)+rot.sinTheta*rec.P.Ind(2))
	p.SetInd(2, -rot.sinTheta*rec.P.Ind(0)+rot.cosTheta*rec.P.Ind(2))

	normal.SetInd(0, rot.cosTheta*rec.Normal.Ind(0)+rot.sinTheta*rec.Normal.Ind(2))
	normal.SetInd(2, -rot.sinTheta*rec.Normal.Ind(0)+rot.cosTheta*rec.Normal.Ind(2))

	rec.P = p.Copy()
	rec.SetFaceNormal(rotatedR, normal)

	return true, rec
}

type Translate struct {
	ptr    Hittable
	offset *vec.Vec3
}

func NewTranslate(ptr Hittable, offset *vec.Vec3) *Translate {
	return &Translate{
		ptr:    ptr,
		offset: offset,
	}
}

func (b *Translate) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	movedR := vec.NewRay(r.Origin().Sub(b.offset), r.Direction(), r.Time())

	hit, rec := b.ptr.Hit(movedR, t0, t1, rec)
	if !hit {
		return false, rec
	}

	rec.P = rec.P.Add(b.offset)
	rec.SetFaceNormal(movedR, rec.Normal)

	return true, rec
}

func (b *Translate) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	bb, outputBox := b.ptr.BoundingBox(t0, t1, outputBox)
	if !bb {
		return false, outputBox
	}

	outputBox = NewAABB(
		outputBox.Min().Add(b.offset),
		outputBox.Max().Add(b.offset),
	)

	return true, outputBox
}

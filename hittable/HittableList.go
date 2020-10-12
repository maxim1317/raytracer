package hittable

import "github.com/maxim1317/raytracer/vec"

type HittableList struct {
	elements []*Hittable
}

func (w *HittableList) Add(h Hittable) {
	w.elements = append(w.elements, &h)
}

func (w *HittableList) Ind(i uint) *Hittable {
	return w.elements[i]
}

func (w *HittableList) Slice(from, to uint) []*Hittable {
	return w.elements[from:to]
}

func (w *HittableList) SetInd(i uint, v *Hittable) {
	w.elements[i] = v
}

func (w *HittableList) Count() int {
	return len(w.elements)
}

func (w *HittableList) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	hitAnything := false
	closest := t1

	for _, element := range w.elements {
		hit, rec := (*element).Hit(r, t0, closest, rec)

		if hit {
			hitAnything = true
			closest = rec.T
		}
	}
	return hitAnything, rec
}

func (w *HittableList) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	if w.Count() == 0 {
		return false, outputBox
	}

	tempBox := new(AABB)
	firstBox := true
	var bFlag bool

	for _, element := range w.elements {
		bFlag, tempBox = (*element).BoundingBox(t0, t1, tempBox)
		if !bFlag {
			return false, outputBox
		}
		if firstBox {
			outputBox = tempBox
		} else {
			outputBox = SurroundingBox(outputBox, tempBox)
		}
	}

	return true, outputBox
}

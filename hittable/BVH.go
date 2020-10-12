package hittable

import (
	"fmt"
	"sort"

	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type BVHNode struct {
	left, right Hittable
	box         AABB
}

func (n *BVHNode) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	if !n.box.Hit(r, t0, t1) {
		return false, rec
	}
	var t float64
	hitLeft, rec := n.left.Hit(r, t0, t1, rec)
	if hitLeft {
		t = rec.T
	} else {
		t = t1
	}
	hitRight, rec := n.right.Hit(r, t0, t, rec)

	return hitLeft || hitRight, rec
}

func (n *BVHNode) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	outputBox = &n.box
	return true, outputBox
}

func NewBVHNodeFull(
	objects HittableList,
	start, end uint,
	time0, time1 float64,
) *BVHNode {
	newNode := new(BVHNode)

	axis := utils.RandInt(0, 2)
	var comparator func(*Hittable, *Hittable) bool
	switch {
	case (axis == 0):
		comparator = boxXCompare
	case (axis == 1):
		comparator = boxYCompare
	default:
		comparator = boxZCompare
	}

	objectSpan := end - start

	switch {
	case objectSpan == 1:
		newNode.left = *objects.Ind(start)
		newNode.right = *objects.Ind(start)
	case objectSpan == 2:
		if comparator(objects.Ind(start), objects.Ind(start+1)) {
			newNode.left = *objects.Ind(start)
			newNode.right = *objects.Ind(start + 1)
		} else {
			newNode.left = *objects.Ind(start + 1)
			newNode.right = *objects.Ind(start)
		}
	default:
		sort.Slice(
			objects.Slice(start, end),
			func(a, b int) bool {
				return boxCompare(objects.Ind(uint(a)), objects.Ind(uint(b)), axis)
			},
		)

		mid := start + objectSpan/2
		newNode.left = NewBVHNodeFull(objects, start, mid, time0, time1)
		newNode.right = NewBVHNodeFull(objects, mid, end, time0, time1)
	}

	boxLeft := new(AABB)
	boxRight := new(AABB)

	lR, boxLeft := newNode.left.BoundingBox(time0, time1, boxLeft)
	rR, boxRight := newNode.right.BoundingBox(time0, time1, boxRight)
	if !lR || !rR {
		fmt.Printf("No bounding box in BVHNode constructor.\n")
	}

	box := SurroundingBox(boxLeft, boxRight)
	newNode.box = *box
	return newNode
}

func NewBVHNode(objects HittableList, time0, time1 float64) *BVHNode {
	return NewBVHNodeFull(objects, 0, uint(objects.Count()), time0, time1)
}

func boxCompare(a, b *Hittable, axis int) bool {
	boxA := new(AABB)
	boxB := new(AABB)

	rA, boxA := (*a).BoundingBox(0, 0, boxA)
	rB, boxB := (*b).BoundingBox(0, 0, boxB)

	if !rA || !rB {
		fmt.Printf("No bounding box in bvh_node constructor.\n")
	}

	return boxA.Min().Ind(axis) < boxB.Min().Ind(axis)
}

func boxXCompare(a, b *Hittable) bool {
	return boxCompare(a, b, 0)
}

func boxYCompare(a, b *Hittable) bool {
	return boxCompare(a, b, 1)
}

func boxZCompare(a, b *Hittable) bool {
	return boxCompare(a, b, 2)
}

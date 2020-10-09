package hittable

import (
	"fmt"
	"sort"

	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type BVHNode struct {
	left, right Hittable
	box         *AABB
}

func (n *BVHNode) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	if !n.box.Hit(r, t0, t1) {
		return false, rec
	}

	hitLeft, rec := n.left.Hit(r, t0, t1, rec)
	if hitLeft {
		t1 = rec.T
	}
	hitRight, rec := n.right.Hit(r, t0, t1, rec)

	return hitLeft || hitRight, rec
}

func (n *BVHNode) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	outputBox = n.box
	return true, outputBox
}

func NewBVHNode(
	objects []*Hittable,
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
		newNode.left = *objects[start]
		newNode.right = *objects[start]
	case objectSpan == 2:
		if comparator(objects[start], objects[start+1]) {
			newNode.left = *objects[start]
			newNode.right = *objects[start+1]
		} else {
			newNode.left = *objects[start+1]
			newNode.right = *objects[start]
		}
	default:
		sort.Slice(
			objects[start:end],
			func(a, b int) bool {
				return boxCompare(objects[a], objects[b], axis)
			},
		)

		mid := start + objectSpan/2
		newNode.left = NewBVHNode(objects, start, mid, time0, time1)
		newNode.right = NewBVHNode(objects, mid, end, time0, time1)
	}

	boxLeft := new(AABB)
	boxRight := new(AABB)

	lR, boxLeft := newNode.left.BoundingBox(time0, time1, boxLeft)
	rR, boxRight := newNode.right.BoundingBox(time0, time1, boxRight)
	if !lR || !rR {
		fmt.Printf("No bounding box in BVHNode constructor.\n")
	}

	newNode.box = SurroundingBox(boxLeft, boxRight)
	return newNode
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

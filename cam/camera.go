package cam

import (
	"math"

	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type Camera struct {
	lowerLeft, horizontal, vertical, origin *vec.Vec3
}

func NewCamera(lookFrom, lookAt, vUp *vec.Vec3, vFOV, aspectRatio float64) *Camera {
	c := new(Camera)

	theta := utils.Degrees2Rad(vFOV)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	const focalLength = 1.0

	w := lookFrom.Sub(lookAt).GetNormal()
	u := vUp.Cross(w).GetNormal()
	v := w.Cross(u)

	c.origin = lookFrom
	c.horizontal = u.MulScalar(viewportWidth)
	c.vertical = v.MulScalar(viewportHeight)
	c.lowerLeft = c.origin.Sub(c.horizontal.DivScalar(2)).Sub(c.vertical.DivScalar(2)).Sub(w)

	return c
}

func (c *Camera) RayAt(s, t float64) *vec.Ray {
	return vec.NewRay(c.origin, c.lowerLeft.Add(c.horizontal.MulScalar(s)).Add(c.vertical.MulScalar(t)).Sub(c.origin))
}

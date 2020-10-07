package cam

import (
	"math"

	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type Camera struct {
	lowerLeft, horizontal, vertical, origin *vec.Vec3
	u, v, w                                 *vec.Vec3
	lensRadius                              float64
}

func NewCamera(
	lookFrom, lookAt, vUp *vec.Vec3,
	vFOV, aspectRatio, aperture, focusDist float64,
) *Camera {
	c := new(Camera)

	theta := utils.Degrees2Rad(vFOV)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	const focalLength = 1.0

	c.w = lookFrom.Sub(lookAt).GetNormal()
	c.u = vUp.Cross(c.w).GetNormal()
	c.v = c.w.Cross(c.u)

	c.origin = lookFrom
	c.horizontal = c.u.MulScalar(viewportWidth * focusDist)
	c.vertical = c.v.MulScalar(viewportHeight * focusDist)
	c.lowerLeft = c.origin.Sub(c.horizontal.DivScalar(2)).Sub(c.vertical.DivScalar(2)).Sub(c.w.MulScalar(focusDist))

	c.lensRadius = aperture / 2.0

	return c
}

func (c *Camera) RayAt(s, t float64) *vec.Ray {
	var rd, offset *vec.Vec3

	rd = vec.NewRandInUnitDisk().MulScalar(c.lensRadius)
	offset = c.u.MulScalar(rd.X()).Add(c.v.MulScalar(rd.Y()))

	return vec.NewRay(c.origin.Add(offset), c.lowerLeft.Add(c.horizontal.MulScalar(s)).Add(c.vertical.MulScalar(t)).Sub(c.origin))
}

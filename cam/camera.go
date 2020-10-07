package cam

import (
	"github.com/maxim1317/raytracer/vec"
)

type Camera struct {
	lowerLeft, horizontal, vertical, origin vec.Vec3
}

func NewCamera() *Camera {
	c := new(Camera)

	var aspectRatio, viewportHeight, viewportWidth, focalLength float64

	aspectRatio = 16.0 / 9.0
	viewportHeight = 2.0
	viewportWidth = aspectRatio * viewportHeight
	focalLength = 1.0

	c.origin = vec.NewZero()
	c.horizontal = vec.New(viewportWidth, 0.0, 0.0)
	c.vertical = vec.New(0.0, viewportHeight, 0.0)

	fc := vec.New(0, 0, focalLength)
	c.lowerLeft = *c.origin.Sub(c.horizontal.DivScalar(2)).Sub(c.vertical.DivScalar(2)).Sub(&fc)

	return c
}

func (c *Camera) position(u float64, v float64) *vec.Vec3 {
	horizontal := c.horizontal.MulScalar(u)
	vertical := c.vertical.MulScalar(v)

	return horizontal.Add(vertical)
}

func (c *Camera) direction(position *vec.Vec3) vec.Vec3 {
	return *c.lowerLeft.Add(position)
}

func (c *Camera) RayAt(u float64, v float64) vec.Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return vec.NewRay(c.origin, direction)
}

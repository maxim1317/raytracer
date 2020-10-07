package cam

import (
	"github.com/maxim1317/raytracer/vec"
)

type Camera struct {
	lowerLeft, horizontal, vertical, origin vec.Vector3D
}

func NewCamera() *Camera {
	c := new(Camera)

	var aspectRatio, viewportHeight, viewportWidth, focalLength float64

	aspectRatio = 16.0 / 9.0
	viewportHeight = 2.0
	viewportWidth = aspectRatio * viewportHeight
	focalLength = 1.0

	c.origin = vec.NewZeroVector3D()
	c.horizontal = vec.NewVector3D(viewportWidth, 0.0, 0.0)
	c.vertical = vec.NewVector3D(0.0, viewportHeight, 0.0)
	c.lowerLeft = c.origin.Sub(c.horizontal.FracScalar(2)).Sub(c.vertical.FracScalar(2)).Sub(vec.NewVector3D(0, 0, focalLength))

	return c
}

func (c *Camera) position(u float64, v float64) vec.Vector3D {
	horizontal := c.horizontal.MulScalar(u)
	vertical := c.vertical.MulScalar(v)

	return horizontal.Add(vertical)
}

func (c *Camera) direction(position vec.Vector3D) vec.Vector3D {
	return c.lowerLeft.Add(position)
}

func (c *Camera) RayAt(u float64, v float64) vec.Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return vec.NewRay(c.origin, direction)
}

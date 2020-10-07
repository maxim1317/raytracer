package main

import (
	"fmt"
	"math"
	"os"

	h "github.com/maxim1317/raytracer/hittable"
	"github.com/maxim1317/raytracer/vec"
)

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		panic(e)
	}
}

const width int = 256
const height int = 256
const color = 255.99

func writePixel(file *os.File, pixel vec.Vector3D) {
	file.WriteString(
		fmt.Sprintf(
			"%v %v %v\n",
			int(color*pixel.X),
			int(color*pixel.Y),
			int(color*pixel.Z),
		),
	)
}

func hitSphere(center vec.Vector3D, radius float64, ray vec.Ray) float64 {
	// t^2*b*b + 2*t*b*(A−C) + (A−C)*(A−C) − r^2 = 0

	var oc vec.Vector3D = ray.Orig.Sub(center)

	var a, bHalf, c, discriminant float64

	a = ray.Dir.LengthSquared()
	bHalf = oc.Dot(ray.Dir)
	c = oc.LengthSquared() - radius*radius
	discriminant = bHalf*bHalf - a*c

	if discriminant < 0 {
		return -1.0
	}
	return (-bHalf - math.Sqrt(discriminant)) / a
}

func rayColor(r *vec.Ray, world *h.World) vec.Vector3D {
	var rec *h.HitRecord = &h.HitRecord{}
	if (*world).Hit(r, 0, math.MaxFloat64, rec) {
		return rec.Normal.Add(vec.NewUnitVector3D()).MulScalar(0.5)
	}
	var unitDir vec.Vector3D = r.Dir.Normalize()
	var t float64 = 0.5 * (unitDir.Y + 1.0)
	return vec.NewUnitVector3D().MulScalar(1.0 - t).Add(vec.NewVector3D(0.5, 0.7, 1.0).MulScalar(t))
}

func main() {
	var err error
	var file *os.File

	// Image

	var aspectRatio float64 = 16.0 / 9.0
	var imageWidth int = 400
	var imageHeight int = int(float64(imageWidth) / aspectRatio)

	// World

	sphere := h.Sphere{
		Center: vec.NewVector3D(0, 0, -1),
		Radius: 0.5,
	}
	floor := h.Sphere{
		Center: vec.NewVector3D(0, -100.5, -1),
		Radius: 100,
	}

	var world h.World = h.World{}

	world.Add(sphere)
	world.Add(floor)

	// Camera

	var viewportHeight float64 = 2.0
	var viewportWidth float64 = viewportHeight * aspectRatio
	var focalLength float64 = 1.0

	var origin vec.Vector3D = vec.NewZeroVector3D()
	var horizontal, vertical vec.Vector3D

	horizontal = vec.NewVector3D(viewportWidth, 0, 0)
	vertical = vec.NewVector3D(0, viewportHeight, 0)

	var lowerLeftCorner vec.Vector3D = origin.Sub(horizontal.FracScalar(2)).Sub(vertical.FracScalar(2)).Sub(vec.NewVector3D(0, 0, focalLength))

	// Render

	const filename = "out.ppm"

	file, err = os.Create(filename)
	check(err, "Couldn't open the file")

	defer file.Close()

	file.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight))

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			var u, v float64

			u = float64(i) / float64(imageWidth-1)
			v = float64(j) / float64(imageHeight-1)

			var pixel vec.Vector3D
			var ray vec.Ray = vec.Ray{
				Orig: origin,
				Dir:  lowerLeftCorner.Add(horizontal.MulScalar(u)).Add(vertical.MulScalar(v)).Sub(origin),
			}

			pixel = rayColor(&ray, &world)
			writePixel(file, pixel)
		}
	}

	// for i := 0; i < count; i++ {

	// }

}

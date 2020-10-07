package main

import (
	"fmt"
	"math"
	"os"

	"github.com/maxim1317/raytracer/cam"
	c "github.com/maxim1317/raytracer/color"
	h "github.com/maxim1317/raytracer/hittable"
	ut "github.com/maxim1317/raytracer/utils"
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

func writePixel(file *os.File, pixel *c.Color, samplesPerPixel int) {
	pixel.DivScalar(float64(samplesPerPixel))
	pixel.Clip(0.0, 0.999)
	file.WriteString(
		fmt.Sprintf(
			"%v %v %v\n",
			int(color*pixel.R()),
			int(color*pixel.G()),
			int(color*pixel.B()),
		),
	)
}

func rayColor(r *vec.Ray, world *h.World) *c.Color {
	color := new(c.Color)
	var rec *h.HitRecord = &h.HitRecord{}
	if (*world).Hit(r, 0, math.MaxFloat64, rec) {
		unit := vec.NewUnit()
		return color.FromVec3(rec.Normal.Add(unit).MulScalar(0.5))
	}
	unitDir := r.Dir.GetNormal()
	t := 0.5 * (unitDir.Y() + 1.0)

	unit := vec.NewUnit()
	blue := vec.New(0.5, 0.7, 1.0)
	return color.FromVec3(unit.MulScalar(1.0 - t).Add(blue.MulScalar(t)))
}

func main() {
	var err error
	var file *os.File

	// Image

	const aspectRatio float64 = 16.0 / 9.0
	const imageWidth int = 400
	const imageHeight int = int(float64(imageWidth) / aspectRatio)
	const samplesPerPixel = 100

	// World

	sphere := h.NewSphere(vec.New(0, 0, -1), 0.5)
	floor := h.NewSphere(vec.New(0, -100.5, -1), 100)

	var world h.World = h.World{}

	world.Add(sphere)
	world.Add(floor)

	// Camera

	camera := cam.NewCamera()

	// Render

	const filename = "out.ppm"

	file, err = os.Create(filename)
	check(err, "Couldn't open the file")

	defer file.Close()

	file.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight))

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			var u, v float64

			pixel := c.NewBlack()

			for s := 0; s < samplesPerPixel; s++ {
				u = (float64(i) + ut.Rand()) / float64(imageWidth-1)
				v = (float64(j) + ut.Rand()) / float64(imageHeight-1)

				ray := camera.RayAt(u, v)
				pixel = *pixel.Add(rayColor(ray, &world))
			}

			writePixel(file, &pixel, samplesPerPixel)
		}
	}

	// for i := 0; i < count; i++ {

	// }

}

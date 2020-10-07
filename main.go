package main

import (
	"fmt"
	"math"
	"os"
	"time"

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
	pixel = pixel.DivScalar(float64(samplesPerPixel)).Gamma2()
	pixel = pixel.Clip(0.0, 0.999)
	file.WriteString(
		fmt.Sprintf(
			"%v %v %v\n",
			int(color*pixel.R()),
			int(color*pixel.G()),
			int(color*pixel.B()),
		),
	)
}

func rayColor(r *vec.Ray, world *h.World, depth int) *c.Color {
	if r.Dir == nil {
		fmt.Printf("%v", depth)
	}
	if depth <= 0 {
		return c.Black()
	}
	color := new(c.Color)
	var rec *h.HitRecord = &h.HitRecord{}
	if (*world).Hit(r, 0.001, math.MaxFloat64, rec) {
		scattered := new(vec.Ray)
		attenuation := c.Black()
		isScut, scattered, attenuation := rec.Mat.Scatter(r, rec, attenuation, scattered)
		if isScut {
			return attenuation.Mul(rayColor(scattered, world, depth-1))
		}
		return c.Black()
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

	aspectRatio := 3.0 / 2.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)
	samplesPerPixel := 100
	maxDepth := 30

	// World

	world := h.RandomWorld()

	// Camera

	lookFrom := vec.New(13, 3, 2)
	lookAt := vec.New(0, 0, 0)
	vUp := vec.New(0, 1, 0)
	distToFocus := 10.0
	aperture := 0.1

	camera := cam.NewCamera(lookFrom, lookAt, vUp, 20, aspectRatio, aperture, distToFocus)

	// Render

	const filename = "out.ppm"

	file, err = os.Create(filename)
	check(err, "Couldn't open the file")

	defer file.Close()

	file.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight))

	startTime := time.Now()
	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Printf("Scanline %v out of %v\n", imageHeight-j, imageHeight)
		for i := 0; i < imageWidth; i++ {
			var u, v float64

			pixel := c.Black()

			for s := 0; s < samplesPerPixel; s++ {
				u = (float64(i) + ut.Rand()) / float64(imageWidth-1)
				v = (float64(j) + ut.Rand()) / float64(imageHeight-1)

				ray := camera.RayAt(u, v)
				pixel = pixel.Add(rayColor(ray, world, maxDepth))
			}

			writePixel(file, pixel, samplesPerPixel)
		}
	}
	endTime := time.Now()
	passed := endTime.Sub(startTime)
	fmt.Printf("Seconds passed: %v\n", passed)

	// for i := 0; i < count; i++ {

	// }

}

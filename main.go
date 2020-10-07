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

	const aspectRatio float64 = 16.0 / 9.0
	const imageWidth int = 400
	const imageHeight int = int(float64(imageWidth) / aspectRatio)
	const samplesPerPixel = 100
	const maxDepth = 30

	// World

	materialGround := h.NewLambertian(c.New(0.8, 0.8, 0.0))
	materialCenter := h.NewLambertian(c.New(0.7, 0.3, 0.3))
	materialLeft := h.NewDielectric(1.5)
	materialRight := h.NewMetal(c.New(0.8, 0.6, 0.2), 1.0)

	sphereGround := h.NewSphere(vec.New(0.0, -100.5, -1.0), 100.0, materialGround)
	sphereCenter := h.NewSphere(vec.New(0.0, 0.0, -1.0), 0.5, materialCenter)
	sphereLeft := h.NewSphere(vec.New(-1.0, 0.0, -1.0), 0.5, materialLeft)
	sphereLeftInner := h.NewSphere(vec.New(-1.0, 0.0, -1.0), -0.3, materialLeft)
	sphereRight := h.NewSphere(vec.New(1.0, 0.0, -1.0), 0.5, materialRight)

	var world h.World = h.World{}

	world.Add(sphereGround)
	world.Add(sphereCenter)
	world.Add(sphereLeft)
	world.Add(sphereLeftInner)
	world.Add(sphereRight)

	// Camera

	camera := cam.NewCamera()

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
				pixel = pixel.Add(rayColor(ray, &world, maxDepth))
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

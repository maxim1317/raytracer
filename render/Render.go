package render

import (
	"fmt"
	"image"
	"math"
	"sync"

	"github.com/maxim1317/raytracer/cam"
	c "github.com/maxim1317/raytracer/color"
	h "github.com/maxim1317/raytracer/hittable"
	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

const (
	maxDepth = 50
	tMin     = 0.001
)

func getPixelColor(r *vec.Ray, world *h.World, depth int) *c.Color {
	if r.Direction() == nil {
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
			return attenuation.Mul(getPixelColor(scattered, world, depth-1))
		}
		return c.Black()
	}

	unitDir := r.Direction().GetNormal()
	t := 0.5 * (unitDir.Y() + 1.0)
	unit := vec.NewUnit()
	blue := vec.New(0.5, 0.7, 1.0)
	return color.FromVec3(unit.MulScalar(1.0 - t).Add(blue.MulScalar(t)))
}

// sample samples rays for anti-aliasing
func sample(world *h.World, camera *cam.Camera, samples, width, height, i, j int) *c.Color {
	rgb := c.Black()

	for s := 0; s < samples; s++ {
		u := (float64(i) + utils.Rand()) / float64(width)
		v := (float64(j) + utils.Rand()) / float64(height)

		ray := camera.GetRay(u, v)
		rgb = rgb.Add(getPixelColor(ray, world, maxDepth))
	}

	// average
	return rgb.DivScalar(float64(samples))
}

// Do performs the render, sampling each pixel the provided number of times
func Do(world *h.World, camera *cam.Camera, cpus, samples, width, height int, ch chan<- int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup

	for i := 0; i < cpus; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for row := i; row < height; row += cpus {
				for col := 0; col < width; col++ {
					rgb := sample(world, camera, samples, width, height, col, row)
					img.Set(col, height-row-1, rgb)
				}
				ch <- 1
			}
		}(i)
	}

	wg.Wait()
	return img
}

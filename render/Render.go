package render

import (
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
	t0       = 0.001
)

func getPixelColor(r *vec.Ray, background *c.Color, world *h.World, depth int) *c.Color {
	var rec *h.HitRecord = new(h.HitRecord)

	if depth <= 0 {
		return c.Black()
	}

	hit, rec := (*world).Hit(r, 0.001, math.MaxFloat64, rec)

	if !hit {
		return background
	}

	scattered := new(vec.Ray)
	attenuation := c.Black()
	emitted := rec.Mat.Emitted(rec.U, rec.V, rec.P)

	isScut, scattered, attenuation := rec.Mat.Scatter(r, rec, attenuation, scattered)
	if !isScut {
		return emitted
	}
	return attenuation.Mul(getPixelColor(scattered, background, world, depth-1))
}

// sample samples rays for anti-aliasing
func sample(world *h.World, camera *cam.Camera, background *c.Color, samples, width, height, i, j int) *c.Color {
	rgb := c.Black()

	for s := 0; s < samples; s++ {
		u := (float64(i) + utils.Rand()) / float64(width)
		v := (float64(j) + utils.Rand()) / float64(height)

		ray := camera.GetRay(u, v)
		rgb = rgb.Add(getPixelColor(ray, background, world, maxDepth))
	}

	// average
	return rgb.DivScalar(float64(samples)).Gamma2().Clip(0.0, 1.0)
}

// Do performs the render, sampling each pixel the provided number of times
func Do(world *h.World, camera *cam.Camera, background *c.Color, cpus, samples, width, height int, ch chan<- int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup

	for i := 0; i < cpus; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			for row := i; row < height; row += cpus {
				for col := 0; col < width; col++ {
					rgb := sample(world, camera, background, samples, width, height, col, row)
					img.Set(col, height-row-1, rgb)
				}
				ch <- 1
			}
		}(i)
	}

	wg.Wait()
	return img
}

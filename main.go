package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/maxim1317/raytracer/cam"
	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/hittable"
	"github.com/maxim1317/raytracer/render"
	"github.com/maxim1317/raytracer/vec"
)

const (
	maxFov      = 120.0
	maxWidth    = 4096
	maxHeight   = 2160
	maxSamples  = 1000
	maxAperture = 0.9

	minFov      = 10.0
	minWidth    = 200
	minHeight   = 100
	minSamples  = 1
	minAperture = 0.001

	progressBarWidth = 80
)

type fileType int

const (
	pngType fileType = iota
	jpegType
)

var (
	cpus    int
	file    string
	x, y, z float64
	version bool

	imageTypes = map[string]interface{}{
		".png":  pngType,
		".jpg":  jpegType,
		".jpeg": jpegType,
	}

	lookfrom    *vec.Vec3 = vec.New(13, 2, 3)
	lookat      *vec.Vec3 = vec.New(0, 0, 0)
	vUp         *vec.Vec3 = vec.New(0, 1, 0)
	vfov        float64   = 40.0
	aperture    float64   = 0.01
	distToFocus float64   = 10.0
	samples     int       = 100

	aspectRatio = 4.0 / 3.0
	width       = 400
)

func main() {
	cpus = runtime.NumCPU()

	start := time.Now()

	// World

	var world *hittable.World
	background := color.Black()

	switch 4 {
	case 1:
		world = hittable.NewRandomWorld()
		background = color.New(0.70, 0.80, 1.00)
		lookfrom = vec.New(13, 2, 3)
		lookat = vec.New(0, 0, 0)
		vfov = 20.0
		aperture = 0.1
	case 2:
		world = hittable.NewTwoSphereWorld()
		background = color.New(0.70, 0.80, 1.00)
		lookfrom = vec.New(13, 2, 3)
		lookat = vec.New(0, 0, 0)
		vfov = 20.0
	case 3:
		world = hittable.NewSimpleLightWorld()
		background = color.Black()
		lookfrom = vec.New(26, 3, 6)
		lookat = vec.New(0, 2, 0)
		samples = 400
		vfov = 20.0
	case 4:
		world = hittable.NewCornellBox()
		width = 600
		aspectRatio = 1.0
		samples = 10
		background = color.Black()
		lookfrom = vec.New(278, 278, -800)
		lookat = vec.New(278, 278, 0)
		vfov = 40.0
	case 5:
		world = hittable.NewTwoBoxWorld()
		background = color.New(0.70, 0.80, 1.00)
		lookfrom = vec.New(13, 3, 0)
		lookat = vec.New(0, 2, 0)
		aspectRatio = 1.0
		samples = 10
		vfov = 20.0
	}

	// Camera

	camera := cam.NewCamera(
		lookfrom, lookat, vUp,
		vfov, aspectRatio,
		aperture, distToFocus,
		0.0, 1.0,
	)

	// Render

	height := int(float64(width) / aspectRatio)

	fmt.Printf("\nRendering %d x %d pixel scene with %d objects:", width, height, world.Count())
	fmt.Printf("\n[%d cpus, %d samples/pixel, %.2f° vfov, %.2f aperture]", cpus, samples, vfov, aperture)

	ch := make(chan int, height)
	defer close(ch)

	go outputProgress(ch, height)

	image := render.Do(world, camera, background, cpus, samples, width, height, ch)

	if err := writeFile("out.png", image); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nDone. Elapsed: %v", time.Since(start))
	fmt.Printf("\nOutput to: %s\n", file)
}

func outputProgress(ch <-chan int, rows int) {
	fmt.Println()
	for i := 1; i <= rows; i++ {
		<-ch
		pct := 100 * float64(i) / float64(rows)
		filled := (progressBarWidth * i) / rows
		bar := strings.Repeat("=", filled) + strings.Repeat("-", progressBarWidth-filled)
		fmt.Printf("\r[%s] %.2f%%", bar, pct)
	}
	fmt.Println()
}

func writeFile(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	ext := strings.ToLower(filepath.Ext(path))

	switch imageType := imageTypes[ext]; imageType {
	case jpegType:
		err = jpeg.Encode(file, img, nil)
	case pngType:
		err = png.Encode(file, img)
	default:
		err = fmt.Errorf("Invalid extension: %s", ext)
	}

	return err
}

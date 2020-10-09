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

	aspectRatio = 4.0 / 3.0
	width       = 400
	height      = int(width / aspectRatio)
	samples     = 100

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

	lookFrom    *vec.Vec3 = vec.New(13, 2, 3)
	lookAt      *vec.Vec3 = vec.New(0, 0, 0)
	vUp         *vec.Vec3 = vec.New(0, 1, 0)
	vFov        float64   = 40.0
	aperture    float64   = 0.0
	distToFocus float64   = 10.0
)

func main() {
	cpus = runtime.NumCPU()

	start := time.Now()

	// World

	var world *hittable.World
	background := color.Black()

	switch 1 {
	case 1:
		world = hittable.RandomWorld()
		background = color.New(0.70, 0.80, 1.00)
		lookFrom = vec.New(13, 2, 3)
		lookAt = vec.New(0, 0, 0)
		vFov = 20.0
		aperture = 0.1
	case 2:
		world = hittable.TwoSphereWorld()
		background = color.New(0.70, 0.80, 1.00)
		lookFrom = vec.New(13, 2, 3)
		lookAt = vec.New(0, 0, 0)
		vFov = 20.0
	case 3:
		background = color.New(0.0, 0.0, 0.0)
	}

	// Camera

	camera := cam.NewCamera(
		lookFrom, lookAt, vUp,
		vFov, float64(width)/float64(height),
		aperture, distToFocus,
		0.0, 1.0,
	)

	// Render

	fmt.Printf("\nRendering %d x %d pixel scene with %d objects:", width, height, world.Count())
	fmt.Printf("\n[%d cpus, %d samples/pixel, %.2f° vFov, %.2f aperture]", cpus, samples, vFov, aperture)

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

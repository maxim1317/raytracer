package texture

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

const bytesPerPixel = 3

type ImageTexture struct {
	data             [][]*color.Color
	width, height    int
	bytesPerScanline int
}

func NewImageTexture(image string) *ImageTexture {
	data, width, height := openImage(image)

	bytesPerScanline := bytesPerPixel * width

	return &ImageTexture{
		data:             data,
		width:            width,
		height:           height,
		bytesPerScanline: bytesPerScanline,
	}
}

func (im *ImageTexture) Value(u, v float64, p *vec.Vec3) *color.Color {
	// Clamp input texture coordinates to [0,1] x [1,0]
	u = utils.Clamp(u, 0.0, 1.0)
	v = 1.0 - utils.Clamp(v, 0.0, 1.0) // Flip V to image coordinates

	i := int(u * float64(im.width))
	j := int(v * float64(im.height))

	// Clamp integer mapping, since actual coordinates should be less than 1.0
	if i >= im.width {
		i = im.width - 1
	}
	if j >= im.height {
		j = im.height - 1
	}

	pixel := im.data[j][i]

	return pixel
}

func openImage(imagePath string) ([][]*color.Color, int, int) {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	file, err := os.Open(imagePath)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, w, h, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	return pixels, w, h
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]*color.Color, int, int, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, 0, 0, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]*color.Color
	for y := 0; y < height; y++ {
		var row []*color.Color
		for x := 0; x < width; x++ {
			row = append(row, rgbToColor(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, width, height, nil
}

// img.At(x, y).RGBA() returns four uint32 values we want a Pixel
func rgbToColor(r uint32, g uint32, b uint32, a uint32) *color.Color {
	return color.New(
		float64(float64(r)/(255.0*255.0)),
		float64(float64(g)/(255.0*255.0)),
		float64(float64(b)/(255.0*255.0)),
	)
}

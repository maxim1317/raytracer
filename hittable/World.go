package hittable

import (
	"github.com/maxim1317/raytracer/color"
	"github.com/maxim1317/raytracer/hittable/texture"
	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

type World struct {
	elements HittableList
}

func (w *World) Add(h Hittable) {
	w.elements.Add(h)
}

func (w *World) Count() int {
	return w.elements.Count()
}

func (w *World) Hit(r *vec.Ray, t0, t1 float64, rec *HitRecord) (bool, *HitRecord) {
	return w.elements.Hit(r, t0, t1, rec)
}

func (w *World) BoundingBox(t0, t1 float64, outputBox *AABB) (bool, *AABB) {
	return w.elements.BoundingBox(t0, t1, outputBox)
}

func NewRandomWorld() *World {
	world := new(World)

	// auto checker = make_shared<checker_texture>(color(0.2, 0.3, 0.1), color(0.9, 0.9, 0.9));
	// world.Add(make_shared<sphere>(vec.New(0,-1000,0), 1000, NewLambertianColored(checker)));

	groundMaterial := NewLambertianTextured(
		texture.NewCheckerColored(
			color.New(0.2, 0.3, 0.1),
			color.New(0.9, 0.9, 0.9),
		),
	)
	world.Add(NewSphere(vec.New(0, -1000, 0), 1000, groundMaterial))

	for a := -6; a < 6; a++ {
		for b := -6; b < 6; b++ {
			chooseMat := utils.Rand()
			center := vec.New(float64(a)+0.9*utils.Rand(), 0.2, float64(b)+0.9*utils.Rand())

			if (center.Sub(vec.New(4, 0.2, 0))).Length() > 0.9 {
				var sphereMaterial Material
				switch {
				case chooseMat < 0.6:
					// diffuse
					albedo := color.Rand().Mul(color.Rand())
					sphereMaterial = NewLambertianColored(albedo)
					center2 := center.Add(vec.New(0, utils.RandRange(0, 0.5), 0))
					world.Add(
						NewMovingSphere(
							center, center2,
							0.0, 1.0, 0.2,
							sphereMaterial,
						),
					)
				case chooseMat < 0.95:
					// metal
					albedo := color.RandInRange(0.5, 1)
					fuzz := utils.RandRange(0, 0.5)
					sphereMaterial = NewMetal(albedo, fuzz)
					world.Add(NewSphere(center, 0.2, sphereMaterial))
				default:
					// glass
					sphereMaterial = NewDielectric(1.5)
					world.Add(NewSphere(center, 0.2, sphereMaterial))

				}
			}
		}
	}

	mat1 := NewDielectric(1.5)
	world.Add(NewSphere(vec.New(0, 1, 0), 1.0, mat1))

	mat2 := NewLambertianColored(color.New(0.4, 0.2, 0.6))
	world.Add(NewSphere(vec.New(-4, 1, 0), 1.0, mat2))

	mat3 := NewMetal(color.New(0.7, 0.6, 0.5), 0.0)
	world.Add(NewSphere(vec.New(4, 1, 0), 1.0, mat3))

	return world
}

func NewTwoSphereWorld() *World {
	world := new(World)

	pertext := texture.NewNoiseTexture(4)
	imgtext := texture.NewImageTexture("world.jpg")

	t1 := NewLambertianTextured(pertext)
	t2 := NewLambertianTextured(imgtext)
	world.Add(NewSphere(vec.New(0, -1000, 0), 1000, t1))
	world.Add(NewSphere(vec.New(0, 2, 0), 2, t2))

	return world
}

func NewSimpleLightWorld() *World {
	world := new(World)

	pertext := texture.NewNoiseTexture(4)
	imgtext := texture.NewImageTexture("world.jpg")

	t1 := NewLambertianTextured(pertext)
	t2 := NewLambertianTextured(imgtext)
	world.Add(NewSphere(vec.New(0, -1000, 0), 1000, t1))
	world.Add(NewSphere(vec.New(0, 2, 0), 2, t2))

	difflight := NewDiffuseLightColored(color.New(0, 0, 4))
	difflight2 := NewDiffuseLightColored(color.New(4, 0, 0))
	world.Add(NewXYRect(3, 5, 1, 3, -2, difflight))
	world.Add(NewSphere(vec.New(0, 7, 0), 1, difflight2))

	return world
}

func NewTwoBoxWorld() *World {
	world := new(World)

	red := NewLambertianColored(color.New(0.65, 0.05, 0.05))
	// white := NewLambertianColored(color.New(0.73, 0.73, 0.73))
	green := NewLambertianColored(color.New(0.12, 0.45, 0.15))
	// blue := NewLambertianColored(color.New(0.5, 0.05, 0.65))
	// light := NewDiffuseLightColored(color.New(15, 15, 15))

	var box1 Hittable
	var box2 Hittable

	world.Add(NewXZRect(-100, 100, -100, 100, 0, red))

	box1 = NewBox(vec.New(-2, 0, -2), vec.New(0, 2, 0), green)
	box1 = NewRotateY(box1, -90)
	box1 = NewTranslate(box1, vec.New(0, 1, 0))

	// box2 = NewBox(vec.New(-2, 0, 0), vec.New(0, 2, 2), blue)
	// world.Add(box2)

	metal := NewMetal(color.New(0.5, 0.05, 0.65), 0)

	// imgtext := texture.NewImageTexture("world.jpg")
	// t2 := NewLambertianTextured(imgtext)
	box2 = NewSphere(vec.New(-1, 1, 1), 1, metal)
	// box2 = NewRotateY(box2, 0)
	// box2 = NewTranslate(box2, vec.New(0, 1, 0))
	world.Add(box1)
	world.Add(box2)

	return world
}

func NewCornellBox() *World {
	world := new(World)

	red := NewLambertianColored(color.New(0.65, 0.05, 0.05))
	white := NewLambertianColored(color.New(0.73, 0.73, 0.73))
	green := NewLambertianColored(color.New(0.12, 0.45, 0.15))
	light := NewDiffuseLightColored(color.New(15, 15, 15))

	world.Add(NewYZRect(0, 555, 0, 555, 555, green))
	world.Add(NewYZRect(0, 555, 0, 555, 0, red))
	world.Add(NewXZRect(213, 343, 227, 332, 554, light))
	world.Add(NewXZRect(0, 555, 0, 555, 0, white))
	world.Add(NewXZRect(0, 555, 0, 555, 555, white))
	world.Add(NewXYRect(0, 555, 0, 555, 555, white))

	// var box1 Hittable
	// var box2 Hittable

	box1 := NewBox(vec.NewZero(), vec.New(165, 330, 165), white)
	box1Rotated := NewRotateY(box1, 15)
	box1RotatedTranslated := NewTranslate(box1Rotated, vec.New(265, 0, 295))
	world.Add(box1RotatedTranslated)

	box2 := NewBox(vec.NewZero(), vec.New(165, 165, 165), white)
	box2Rotated := NewRotateY(box2, -18)
	box2Translated := NewTranslate(box2Rotated, vec.New(130, 0, 65))
	world.Add(box2Translated)

	return world
}

func NewCornellSmoke() *World {
	world := new(World)

	red := NewLambertianColored(color.New(0.65, 0.05, 0.05))
	white := NewLambertianColored(color.New(0.73, 0.73, 0.73))
	green := NewLambertianColored(color.New(0.12, 0.45, 0.15))
	light := NewDiffuseLightColored(color.New(15, 15, 15))

	world.Add(NewYZRect(0, 555, 0, 555, 555, green))
	world.Add(NewYZRect(0, 555, 0, 555, 0, red))
	world.Add(NewXZRect(213, 343, 227, 332, 554, light))
	world.Add(NewXZRect(0, 555, 0, 555, 0, white))
	world.Add(NewXZRect(0, 555, 0, 555, 555, white))
	world.Add(NewXYRect(0, 555, 0, 555, 555, white))

	var box1 Hittable
	var box2 Hittable

	box1 = NewBox(vec.NewZero(), vec.New(165, 330, 165), white)
	box1 = NewRotateY(box1, 15)
	box1 = NewTranslate(box1, vec.New(265, 0, 295))
	world.Add(NewConstantMediumColored(box1, 0.01, color.Black()))

	box2 = NewBox(vec.NewZero(), vec.New(165, 165, 165), white)
	box2 = NewRotateY(box2, -18)
	box2 = NewTranslate(box2, vec.New(130, 0, 65))
	world.Add(NewConstantMediumColored(box2, 0.01, color.White()))

	return world
}

func NewFinalScene() *World {
	world := new(World)
	boxes1 := new(HittableList)
	ground := NewLambertianColored(color.New(0.48, 0.83, 0.53))

	boxesPerSide := 20
	for i := 0; i < boxesPerSide; i++ {
		for j := 0; j < boxesPerSide; j++ {
			w := 100.0
			x0 := -1000.0 + float64(i)*w
			z0 := -1000.0 + float64(j)*w
			y0 := 0.0
			x1 := x0 + w
			y1 := utils.RandRange(1, 101)
			z1 := z0 + w

			boxes1.Add(NewBox(vec.New(x0, y0, z0), vec.New(x1, y1, z1), ground))
		}
	}

	world.Add(NewBVHNode(*boxes1, 0, 1))

	light := NewDiffuseLightColored(color.New(7, 7, 7))
	world.Add(NewXZRect(123, 423, 147, 412, 554, light))

	center1 := vec.New(400, 400, 200)
	center2 := center1.Add(vec.New(30, 0, 0))
	movingSphereMaterial := NewLambertianColored(color.New(0.7, 0.3, 0.1))
	world.Add(NewMovingSphere(center1, center2, 0, 1, 50, movingSphereMaterial))

	world.Add(NewSphere(vec.New(260, 150, 45), 50, NewDielectric(1.5)))
	world.Add(NewSphere(vec.New(0, 150, 145), 50, NewMetal(color.New(0.8, 0.8, 0.9), 1.0)))

	boundary := NewSphere(vec.New(360, 150, 145), 70, NewDielectric(1.5))
	world.Add(boundary)
	world.Add(NewConstantMediumColored(boundary, 0.2, color.New(0.2, 0.4, 0.9)))
	boundary = NewSphere(vec.New(0, 0, 0), 5000, NewDielectric(1.5))
	world.Add(NewConstantMediumColored(boundary, .0001, color.New(1, 1, 1)))

	emat := NewLambertianTextured(texture.NewImageTexture("world.jpg"))
	world.Add(NewSphere(vec.New(400, 200, 400), 100, emat))
	pertext := texture.NewNoiseTexture(0.1)
	world.Add(NewSphere(vec.New(220, 280, 300), 80, NewLambertianTextured(pertext)))

	boxes2 := new(HittableList)
	white := NewLambertianColored(color.New(0.73, 0.73, 0.73))
	ns := 1000
	for j := 0; j < ns; j++ {
		boxes2.Add(
			NewSphere(vec.NewRandInRange(0, 165), 10, white),
		)
	}

	world.Add(
		NewTranslate(
			NewRotateY(
				NewBVHNode(*boxes2, 0.0, 1.0), 15,
			),
			vec.New(-100, 270, 395),
		),
	)

	return world
}

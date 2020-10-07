package vec

type Ray struct {
	Orig, Dir *Vec3
}

func (r *Ray) At(t float64) *Vec3 {
	orig := r.Orig.Add(r.Dir.MulScalar(t))
	return orig
}

func NewRay(orig, dir *Vec3) *Ray {
	ray := new(Ray)
	ray.Orig = orig
	ray.Dir = dir
	return ray
}

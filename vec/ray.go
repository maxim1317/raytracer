package vec

type Ray struct {
	Orig, Dir Vector3D
}

func (r Ray) At(t float64) Vector3D {
	return r.Orig.Add(r.Dir.MulScalar(t))
}

func NewRay(orig, dir Vector3D) Ray {
	return Ray{
		Orig: orig,
		Dir:  dir,
	}
}

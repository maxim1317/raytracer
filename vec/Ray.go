package vec

// Ray structure provides simple storage for rays
type Ray struct {
	origin, direction *Vec3
	time              float64
}

// Origin getter
func (r *Ray) Origin() *Vec3 {
	return r.origin
}

// Direction getter
func (r *Ray) Direction() *Vec3 {
	return r.direction
}

// Time getter
func (r *Ray) Time() float64 {
	return r.time
}

// At returns subray from origin at point t
func (r *Ray) At(t float64) *Vec3 {
	rayAt := r.origin.Add(r.direction.MulScalar(t))
	return rayAt
}

// NewRay creates new ray
func NewRay(origin, direction *Vec3, time float64) *Ray {
	ray := new(Ray)
	ray.origin = origin
	ray.direction = direction
	ray.time = time
	return ray
}

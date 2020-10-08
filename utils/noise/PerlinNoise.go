package noise

import (
	"math"

	"github.com/maxim1317/raytracer/utils"
	"github.com/maxim1317/raytracer/vec"
)

const pointCount = 256

type Perlin struct {
	ranvec              [pointCount]*vec.Vec3
	permX, permY, permZ [pointCount]int
}

func NewPerlin() *Perlin {
	var ranvec [pointCount]*vec.Vec3
	for i := 0; i < pointCount; i++ {
		rand := vec.NewRandInRange(-1.0, 1.0)
		ranvec[i] = rand.GetNormal()
	}

	return &Perlin{
		ranvec: ranvec,
		permX:  perlinGeneratePerm(),
		permY:  perlinGeneratePerm(),
		permZ:  perlinGeneratePerm(),
	}
}

func (n *Perlin) Noise(p *vec.Vec3) float64 {
	u := p.X() - math.Floor(p.X())
	v := p.Y() - math.Floor(p.Y())
	w := p.Z() - math.Floor(p.Z())
	i := int(math.Floor(p.X()))
	j := int(math.Floor(p.Y()))
	k := int(math.Floor(p.Z()))
	var c [2][2][2]*vec.Vec3

	for di := 0; di < 2; di++ {
		for dj := 0; dj < 2; dj++ {
			for dk := 0; dk < 2; dk++ {
				ind := n.permX[(i+di)&255] ^ n.permY[(j+dj)&255] ^ n.permZ[(k+dk)&255]
				c[di][dj][dk] = n.ranvec[ind]
			}
		}
	}

	return perlinInterp(c, u, v, w)
}

func (n *Perlin) Turb(p *vec.Vec3, depth int) float64 {
	accum := 0.0
	tempP := p
	weight := 1.0

	for i := 0; i < depth; i++ {
		accum += weight * n.Noise(tempP)
		weight *= 0.5
		tempP = tempP.MulScalar(2)
	}

	return math.Abs(accum)
}

func perlinInterp(c [2][2][2]*vec.Vec3, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)
	accum := 0.0

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				weightV := vec.New(u-float64(i), v-float64(j), w-float64(k))
				tmp := float64(i)*uu + (1-float64(i))*(1-uu)
				tmp *= float64(j)*vv + (1-float64(j))*(1-vv)
				tmp *= float64(k)*ww + (1-float64(k))*(1-ww)
				accum += tmp * c[i][j][k].Dot(weightV)
			}
		}
	}

	return accum
}

func trilinearInterp(c [2][2][2]float64, u, v, w float64) float64 {
	accum := 0.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				tmp := float64(i)*u + (1-float64(i))*(1-u)
				tmp *= float64(j)*v + (1-float64(j))*(1-v)
				tmp *= float64(k)*w + (1-float64(k))*(1-w)
				accum += tmp * c[i][j][k]
			}
		}
	}
	return accum
}

func permute(p [pointCount]int) [pointCount]int {
	for i := pointCount - 1; i > 0; i-- {
		target := utils.RandInt(0, i)
		tmp := p[i]
		p[i] = p[target]
		p[target] = tmp
	}
	return p
}

func perlinGeneratePerm() [pointCount]int {
	var p [pointCount]int

	for i := 0; i < pointCount; i++ {
		p[i] = i
	}

	p = permute(p)

	return p
}

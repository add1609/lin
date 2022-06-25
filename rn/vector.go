package rn

import "math"

type Vec struct {
	N int
	X []float64
}

// VecN returns a Vec object that has length of n,
// and all elements set to val
func VecN(n int, val float64) Vec {
	vslice := make([]float64, n)
	for i := range vslice {
		vslice[i] = val
	}
	return Vec{N: n, X: vslice}
}

// VecEqual returns true if p.N == q.N and p.X == q.X
func VecEqual(p, q Vec) bool {
	if len(p.X) != len(q.X) || p.N != q.N {
		return false
	}
	for i, v := range p.X {
		if v != q.X[i] {
			return false
		}
	}
	return true
}

// At returns the value of p at index i
func (p *Vec) At(i int) float64 {
	if p.N <= i {
		panic(ErrIndexOutOfRange)
	}
	return p.X[i]
}

// Set Sets the value of p[i] = val
func (p *Vec) Set(i int, val float64) {
	if p.N <= i {
		panic(ErrIndexOutOfRange)
	}
	p.X[i] = val
}

// SetSlice Copies the float64 slice s onto p.X
func (p *Vec) SetSlice(s []float64) {
	if p.N < len(s) {
		panic(ErrSliceLengthMismatch)
	}
	copy(p.X, s)
}

// Abs returns the absolute value of p
func Abs(p Vec) Vec {
	if len(p.X) < 1 {
		panic(ErrZeroLengthVec)
	}
	u := Vec{N: len(p.X), X: make([]float64, len(p.X))}

	for i := 0; i < len(u.X); i++ {
		u.Set(i, math.Abs(p.X[i]))
	}
	return u
}

// Add returns the vector sum of p and q
func Add(p, q Vec) Vec {
	if len(p.X) != len(q.X) || p.N != q.N {
		panic(ErrShape)
	}
	u := Vec{N: len(p.X), X: make([]float64, len(p.X))}

	for i := 0; i < len(u.X); i++ {
		u.Set(i, p.X[i]+q.X[i])
	}
	return u
}

// Sub returns the vector sum of p and -q
func Sub(p, q Vec) Vec {
	if len(p.X) != len(q.X) || p.N != q.N {
		panic(ErrShape)
	}
	u := Vec{N: len(p.X), X: make([]float64, len(p.X))}

	for i := 0; i < len(u.X); i++ {
		u.Set(i, p.X[i]-q.X[i])
	}
	return u
}

// Scale returns the vector p scaled by r
func Scale(r float64, p Vec) Vec {
	if len(p.X) < 1 {
		panic(ErrZeroLengthVec)
	}
	u := Vec{N: len(p.X), X: make([]float64, len(p.X))}

	for i := 0; i < len(u.X); i++ {
		u.Set(i, r*p.X[i])
	}
	return u
}

// Dot returns the dot product of p and q
func Dot(p, q Vec) float64 {
	if len(p.X) != len(q.X) || p.N != q.N {
		panic(ErrShape)
	}
	d := 0.0
	for i := 0; i < len(p.X); i++ {
		d += p.X[i] * q.X[i]
	}
	return d
}

// Cross returns the cross product of p and q
func Cross(p, q Vec) Vec {
	if len(p.X) != len(q.X) || p.N != q.N {
		panic(ErrShape)
	}
	switch p.N {
	default:
		panic(ErrShape)
	case 3:
		u := Vec{N: 3, X: make([]float64, 3)}
		u.X[0] = (p.X[1] * q.X[2]) - (p.X[2] * q.X[1])
		u.X[1] = (p.X[2] * q.X[0]) - (p.X[0] * q.X[2])
		u.X[2] = (p.X[0] * q.X[1]) - (p.X[1] * q.X[0])
		return u
	}
}

// Norm returns the Euclidean norm of p
//
//	||p|| = sqrt(p · p)
func Norm(p Vec) float64 {
	return math.Sqrt(Dot(p, p))
}

// Dist returns the distance between p and q
//
//	||q - p||
func Dist(p, q Vec) float64 {
	return Norm(Sub(p, q))
}

// Cos returns the cosine of the opening angle between p and q
func Cos(p, q Vec) float64 {
	return Dot(p, q) / (Norm(p) * Norm(q))
}

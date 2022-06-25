package rn

import "fmt"

type Order uint

const (
	COL Order = iota // 0
	ROW              // 1
)

func (o Order) String() string {
	switch o {
	case COL:
		return "COL"
	case ROW:
		return "ROW"
	default:
		panic(ErrOrder)
	}
}

type Mat struct {
	// How many rows, columns this matrix has
	Rows, Cols int

	// The order of the matrix. Can be either
	// 0 => Column major
	// 1 => Row major
	Order

	// 1D array that represents the matrix in
	// either column major or row major order
	X []float64
}

func (m *Mat) String() string {
	return fmt.Sprintf(
		"{Rows: %v, Cols: %v, Order: %v\nMat: %v}",
		m.Rows, m.Cols, m.Order, m.X,
	)
}

func (m *Mat) Dims() (r, c int) {
	return m.Rows, m.Cols
}

func (m *Mat) T() Mat {
	tMat := *m
	// Swap dimensions of mat
	tMat.Cols, tMat.Rows = m.Dims()
	// Swap order of mat
	tMat.Order = (m.Order + 1) % 2

	return tMat
}

func (m *Mat) At(i, j int) float64 {
	if len(m.X) < i*j {
		panic(ErrIndexOutOfRange)
	}
	if i < 0 {
		panic(ErrNegativeDimension)
	}
	if j < 0 {
		panic(ErrNegativeDimension)
	}

	switch m.Order {
	case COL:
		if m.Cols <= i {
			panic(ErrColAccess)
		}
		if m.Rows <= j {
			panic(ErrRowAccess)
		}
		// [i][j]    [i + j*Cols]
		// [0][0] => [0 + 0*3] = [0]
		// [0][1] => [0 + 1*3] = [3]
		// [1][0] => [1 + 0*3] = [1]
		// [1][1] => [1 + 1*3] = [4]
		// [2][0] => [2 + 0*3] = [2]
		// [2][1] => [2 + 1*3] = [5]
		return m.X[i+j*m.Cols]
	case ROW:
		if m.Rows <= i {
			panic(ErrRowAccess)
		}
		if m.Cols <= j {
			panic(ErrColAccess)
		}
		// [i][j]    [i*Cols + j]
		// [0][0] => [0*3 + 0] = [0]
		// [0][1] => [0*3 + 1] = [1]
		// [0][2] => [0*3 + 2] = [2]
		// [1][0] => [1*3 + 0] = [3]
		// [1][1] => [1*3 + 1] = [4]
		// [1][2] => [1*3 + 2] = [5]
		return m.X[i*m.Cols+j]
	default:
		panic(ErrOrder)
	}
}

// NewMat returns a new mat object with
// Dims = (m, n), o order and all elements
// set to v
func NewMat(m, n int, o Order, v float64) Mat {
	if m < 1 || n < 1 {
		panic(ErrZeroLengthMat)
	}
	mSlice := make([]float64, m*n)
	for i := range mSlice {
		mSlice[i] = v
	}
	return Mat{Rows: m, Cols: n, X: mSlice, Order: o}
}

func MatEqual(a, b Mat) bool {
	if len(a.X) != len(b.X) || a.Rows != b.Rows || a.Cols != b.Cols || a.Order != b.Order {
		return false
	}
	for i, v := range a.X {
		if v != b.X[i] {
			return false
		}
	}
	return true
}

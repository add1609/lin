package rn

// Error represents matrix/vector handling errors
type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	ErrPivot               = Error{"lin: malformed pivot list"}
	ErrOrder               = Error{"lin: invalid order for matrix"}
	ErrShape               = Error{"lin: dimension mismatch"}
	ErrSquare              = Error{"lin: expect square matrix"}
	ErrSingular            = Error{"lin: matrix is singular"}
	ErrColLength           = Error{"lin: col length mismatch"}
	ErrNormOrder           = Error{"lin: invalid norm order for matrix"}
	ErrColAccess           = Error{"lin: column index out of range"}
	ErrRowAccess           = Error{"lin: row index out of range"}
	ErrRowLength           = Error{"lin: row length mismatch"}
	ErrVectorAccess        = Error{"lin: vector index out of range"}
	ErrZeroLengthMat       = Error{"lin: zero length in matrix dimension"}
	ErrZeroLengthVec       = Error{"lin: zero length in vector dimension"}
	ErrIndexOutOfRange     = Error{"lin: index out of range"}
	ErrNegativeDimension   = Error{"lin: negative dimension"}
	ErrSliceLengthMismatch = Error{"lin: input slice length mismatch"}
)

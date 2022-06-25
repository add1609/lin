package rn

// Error represents matrix/vector handling errors
type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	ErrPivot               = Error{"mat: malformed pivot list"}
	ErrOrder               = Error{"mat: invalid order for matrix"}
	ErrShape               = Error{"mat: dimension mismatch"}
	ErrSquare              = Error{"mat: expect square matrix"}
	ErrSingular            = Error{"mat: matrix is singular"}
	ErrColLength           = Error{"mat: col length mismatch"}
	ErrNormOrder           = Error{"mat: invalid norm order for matrix"}
	ErrColAccess           = Error{"mat: column index out of range"}
	ErrRowAccess           = Error{"mat: row index out of range"}
	ErrRowLength           = Error{"mat: row length mismatch"}
	ErrVectorAccess        = Error{"mat: vector index out of range"}
	ErrZeroLengthMat       = Error{"mat: zero length in matrix dimension"}
	ErrZeroLengthVec       = Error{"mat: zero length in vector dimension"}
	ErrIndexOutOfRange     = Error{"mat: index out of range"}
	ErrNegativeDimension   = Error{"mat: negative dimension"}
	ErrSliceLengthMismatch = Error{"mat: input slice length mismatch"}
)

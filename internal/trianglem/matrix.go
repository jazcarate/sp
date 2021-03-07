// Package trianglem represent a triangle matrix where one side is the inverse of the other
package trianglem

import (
	"errors"
)

// M represents a triangle matrix.
/*
Some optimizations hold where:
 1  2  3  4 		1=6=11=16= 🤷‍♂️
 5  6  7  8			2 = -5		7 = -10
 9 10 11 12			3 = -9		8 = -11
13 14 15 16			4 = -13		12 = -15

As we can represent the matrix with only half the values, the memory representation is:
data: 2 3 7 4 8 12
size: 4

This might look convoluted, but we are basically incrementig the matirx by adding a (size-1) column.
*/
type M struct {
	data []int
	size int
}

const def int = 0

// Get the value at position (x,y).
func (t *M) Get(x, y int) (int, error) {
	if t == nil {
		return -1, ErrOutOfBoundsMatrix
	}

	if x == y {
		return def, nil
	}

	if x > t.size || y > t.size {
		return -1, ErrOutOfBoundsMatrix
	}

	if x < y {
		val, err := t.Get(y, x)
		return -val, err
	}

	return t.data[toDiagCoordinates(x, y)], nil
}

// Iterate over a column.
func (t *M) Iterate(col int) []int {
	if t == nil {
		return nil
	}

	ret := make([]int, t.size)

	for y := 0; y < t.size; y++ {
		var err error
		ret[y], err = t.Get(col, y)

		if err != nil {
			panic("should be able to get the whole column")
		}
	}

	return ret
}

var (
	// ErrOutOfBoundsMatrix represents an error when trying to set outside bounds.
	ErrOutOfBoundsMatrix = errors.New("can't set on an empty matrix")
	// ErrCantSetDiagonal represents an error when trying to set the diagonal.
	ErrCantSetDiagonal = errors.New("can't do operations on the diagonal")
)

// Set the value at position (x,y).
func (t *M) Set(x, y, val int) error {
	if t == nil {
		return ErrOutOfBoundsMatrix
	}

	if x == y {
		return ErrCantSetDiagonal
	}

	if x < y {
		return t.Set(y, x, -val)
	}

	t.data[toDiagCoordinates(x, y)] = val

	return nil
}

// Modify the value at position (x,y) by a given function.
func (t *M) Modify(x, y int, mod func(int) int) error {
	val, err := t.Get(x, y)
	if err != nil {
		return err
	}

	return t.Set(x, y, mod(val))
}

// Incr the underlying storage by 1.
func (t *M) Incr() *M {
	return t.IncrD(1)
}

// IncrD the underlying storage by any size.
func (t *M) IncrD(delta int) *M {
	newSize := delta

	if t != nil {
		newSize += t.size
	}

	target := make([]int, realSize(newSize))

	if t != nil {
		copy(target, t.data)
	}

	return &M{data: target, size: newSize}
}

func realSize(size int) int {
	// Half of ( Area of a square - Diagonal )
	// (x² - x) / 2
	return (size*size - size) >> 1
}

func toDiagCoordinates(x, y int) int {
	return t(x-1) + y
}

func t(n int) int {
	// https://oeis.org/A000217
	return ((n + 1) * n) >> 1
}

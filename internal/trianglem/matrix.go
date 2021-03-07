// Package M represent a triangle matrix where one side is the inverse of the other
// Useful for tracking balances.
package trianglem

import (
	"errors"
)

/*
This represents a triangle matrix where:
 1  2  3  4 		1=6=11=16= 🤷‍♂️
 5  6  7  8			2 = -5		7 = -10
 9 10 11 12			3 = -9		8 = -11
13 14 15 16			4 = -13		12 = -15

As we can represent the matrix with only half the values, the memory representation is:
data: 4 3 8 2 7 12
size: 4

This might look convoluted, bt we are going from top left, in a ↘ diagonal.
*/
type M struct {
	data []int
	size int
}

const def int = 0

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

	return t.data[toDiagCoordinates(x, y, t.size)], nil
}

var (
	// ErrOutOfBoundsMatrix represents an error when trying to set outside bounds.
	ErrOutOfBoundsMatrix = errors.New("can't set on an empty matrix")
	// ErrCantSetDiagonal represents an error when trying to set the diagonal.
	ErrCantSetDiagonal = errors.New("can't do operations on the diagonal")
)

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

	t.data[toDiagCoordinates(x, y, t.size)] = val

	return nil
}

func (t *M) Incr() *M {
	return t.IncrD(1)
}

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
	return (size*size - size) / 2
}

func toDiagCoordinates(x, y, size int) int {
	invX := size - x - 1

	return t(y+1) + t(invX) - t(y-1) - 1
}

func t(n int) int {
	// https://oeis.org/A000217
	return ((n + 1) * n) >> 1
}

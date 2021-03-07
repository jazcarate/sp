package trianglem_test

import (
	"testing"

	"github.com/jazcarate/sp/internal/trianglem"
	"github.com/stretchr/testify/assert"
)

func TestMatrix_NilValueGet(t *testing.T) {
	var m *trianglem.M

	_, err := m.Get(3, 2)

	if assert.Error(t, err) {
		assert.Equal(t, trianglem.ErrOutOfBoundsMatrix, err)
	}
}

func TestMatrix_NilValueSet(t *testing.T) {
	var m *trianglem.M

	err := m.Set(3, 2, 10)

	if assert.Error(t, err) {
		assert.Equal(t, trianglem.ErrOutOfBoundsMatrix, err)
	}
}

func TestMatrix_OutOfBounds(t *testing.T) {
	var m *trianglem.M
	m = m.Incr()

	_, err := m.Get(3, 2)

	if assert.Error(t, err) {
		assert.Equal(t, trianglem.ErrOutOfBoundsMatrix, err)
	}
}

func TestMatrix_1x1(t *testing.T) {
	var m *trianglem.M
	m = m.Incr()

	err := m.Set(1, 1, 10)

	if assert.Error(t, err) {
		assert.Equal(t, trianglem.ErrCantSetDiagonal, err)
	}
}

func TestMatrix_IncreaseSize(t *testing.T) {
	var m *trianglem.M
	m = m.Incr()

	v, err := m.Get(0, 0)

	assert.Empty(t, err)
	assert.Empty(t, v)
}

func TestMatrix_StoreInBigMatrix(t *testing.T) {
	var m *trianglem.M
	m = m.IncrD(3)

	err1 := m.Set(0, 1, 10)
	err2 := m.Set(2, 0, -11)
	err3 := m.Set(1, 2, 12)

	assert.Empty(t, err1)
	assert.Empty(t, err2)
	assert.Empty(t, err3)

	assetGet(t, m, 10, 0, 1)
	assetGet(t, m, -10, 1, 0)

	assetGet(t, m, 11, 0, 2)
	assetGet(t, m, -11, 2, 0)

	assetGet(t, m, 12, 1, 2)
	assetGet(t, m, -12, 2, 1)
}

func TestMatrix_IncreaseSizeKeepsValues(t *testing.T) {
	var m *trianglem.M
	m = m.IncrD(3)

	err1 := m.Set(0, 1, 10)
	err2 := m.Set(2, 0, -11)
	err3 := m.Set(1, 2, 12)

	assert.Empty(t, err1)
	assert.Empty(t, err2)
	assert.Empty(t, err3)

	m = m.IncrD(1)

	assetGet(t, m, 10, 0, 1)
	assetGet(t, m, -10, 1, 0)

	assetGet(t, m, 11, 0, 2)
	assetGet(t, m, -11, 2, 0)

	assetGet(t, m, 12, 1, 2)
	assetGet(t, m, -12, 2, 1)

	assetGet(t, m, 0, 0, 3)
	assetGet(t, m, 0, 3, 0)
}

func assetGet(t *testing.T, m *trianglem.M, val, x, y int) {
	v, err := m.Get(x, y)

	assert.Empty(t, err)
	assert.Equal(t, val, v)
}

package trianglem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jazcarate/sp/internal/trianglem"
)

func TestMatrix_NilValueGet(t *testing.T) {
	var m *trianglem.M

	assert.Panics(t, func() { m.Get(3, 2) }, "The code did not panic")
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

	assert.Panics(t, func() { m.Get(3, 2) }, "The code did not panic")
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

	v := m.Get(0, 0)

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

func TestMatrix_Modify(t *testing.T) {
	var m *trianglem.M
	m = m.IncrD(2)

	_ = m.Set(0, 1, 10)
	err := m.Modify(0, 1, func(i int) int { return i * 2 })

	assert.Empty(t, err)
	assetGet(t, m, 10*2, 0, 1)
}

func TestMatrix_IterateEmpty(t *testing.T) {
	var m *trianglem.M

	val := m.Iterate(3)

	assert.Empty(t, val)
}

func TestMatrix_IterateWithValues(t *testing.T) {
	var m *trianglem.M
	m = m.IncrD(2)

	_ = m.Set(1, 0, 10)

	val := m.Iterate(1)

	assert.Equal(t, []int{10, 0}, val)
}

func TestMatrix_PreatyPrint(t *testing.T) {
	var m *trianglem.M
	m = m.IncrD(3)

	_ = m.Set(0, 1, 10)
	_ = m.Set(2, 0, -11)
	_ = m.Set(1, 2, 12)

	assert.Equal(t, "||   0|  10|  11||\n|| -10|   0|  12||\n|| -11| -12|   0||\n", m.String())
}

func assetGet(t *testing.T, m *trianglem.M, val, x, y int) {
	v := m.Get(x, y)

	assert.Equal(t, val, v)
}

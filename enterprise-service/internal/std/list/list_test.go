package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Integer int

func (i Integer) Compare(with Integer) int {
	if i == with {
		return 0
	}

	if i > with {
		return 1
	}

	return -1
}

// assert that created new ll is empty
func TestCreateList(t *testing.T) {
	//arrange
	sut := New[Integer]()

	assert.Nil(t, sut.head)
	assert.Nil(t, sut.head)
	assert.Equal(t, sut.Count(), 0)
}

func TestAddInEmptyList(t *testing.T) {
	sut := New[Integer]()

	sut.Add(Integer(10))

	assert.Equal(t, sut.Count(), 1)
	assert.Equal(t, sut.head, sut.tail)
}

// Adding many different notes
func TestAddManyList(t *testing.T) {
	tests := []node[Integer]{}
	var i Integer
	for i = 0; i < 10; i++ {
		tests = append(tests, node[Integer]{nil, i})
	}

	sut := New[Integer]()

	for _, edge := range tests {
		t.Run(fmt.Sprintf("%d", edge.value), func(t *testing.T) {
			sut.Add(Integer(edge.value))
			assert.Equal(t, sut.tail, &edge)
		})
	}
}

// checks find method with list
// with distinct values
func TestFindInSetList(t *testing.T) {
	sut := New[Integer]()
	var i Integer
	for i = 0; i < 10; i++ {
		sut.Add(Integer(i))
	}

	for i = 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(i)
			assert.Nil(t, err)
			assert.Equal(t, i, res.value)
		})
	}
}

// Checks find method in empty list
func TestFindInEmptyList(t *testing.T) {
	sut := New[Integer]()
	_, err := sut.Find(10)
	assert.NotNil(t, err)
}

// Checks find method in list with duplicates
func TestFindInList(t *testing.T) {
	duplicate := Integer(10)
	sut := New[Integer]()
	var i Integer
	for i = 0; i < 10; i++ {
		sut.Add(Integer(duplicate))
	}

	for i = 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res, err := sut.Find(duplicate)
			assert.Nil(t, err)
			assert.Equal(t, duplicate, res.value)
		})
	}
}

func TestCountInEmptyList(t *testing.T) {
	sut := New[Integer]()

	assert.Equal(t, 0, sut.Count())
}

func TestPopFrontInEmpty(t *testing.T) {
	sut := New[Integer]()

	_, ok := sut.PopFront()
	assert.False(t, ok)

}

func TestPopFrontInFilled(t *testing.T) {
	sut := New[Integer]()
	var i Integer
	for i = 0; i < 10; i++ {
		sut.Add(Integer(i))
	}

	for i = 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res, ok := sut.PopFront()
			assert.True(t, ok)
			assert.Equal(t, Integer(i), res)
		})
	}
}

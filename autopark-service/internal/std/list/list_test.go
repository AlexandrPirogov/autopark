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

	sut.PushBack(Integer(10))

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
			sut.PushBack(Integer(edge.value))
			assert.Equal(t, sut.tail, &edge)
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
		sut.PushBack(Integer(i))
	}

	for i = 0; i < 10; i++ {
		t.Run("", func(t *testing.T) {
			res, ok := sut.PopFront()
			assert.True(t, ok)
			assert.Equal(t, Integer(i), res)
		})
	}
}

func TestIteratorCreateInEmptyList(t *testing.T) {
	l := New[Integer]()
	sut := l.Iterator()

	assert.False(t, sut.Next())
}

func TestIteratorMovementInNonEmptyList(t *testing.T) {
	l := New[Integer]()
	var i Integer
	for i = 0; i < 10; i++ {
		l.PushBack(Integer(i))
	}

	sut := l.Iterator()
	for expected := 0; expected < 10; expected++ {
		actual := sut.Curr()
		sut.Next()
		assert.Equal(t, Integer(expected), actual)
	}

	assert.False(t, sut.Next())
}

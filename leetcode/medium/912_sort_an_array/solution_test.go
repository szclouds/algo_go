package _912_sort_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortArray(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 3, 5}, sortArray([]int{5, 2, 3, 1}))
	ast.Equal([]int{0, 0, 1, 1, 2, 5}, sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func TestSortArray2(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 3, 5}, sortArray2([]int{5, 2, 3, 1}))
	ast.Equal([]int{0, 0, 1, 1, 2, 5}, sortArray2([]int{5, 1, 1, 2, 0, 0}))
}

func TestSortArray3(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 3, 5}, sortArray3([]int{5, 2, 3, 1}))
	ast.Equal([]int{0, 0, 1, 1, 2, 5}, sortArray3([]int{5, 1, 1, 2, 0, 0}))
}
func TestSortArray4(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 3, 5}, sortArray4([]int{5, 2, 3, 1}))
	ast.Equal([]int{0, 0, 1, 1, 2, 5}, sortArray4([]int{5, 1, 1, 2, 0, 0}))
}
func TestSortArray5(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{1, 2, 3, 5}, sortArray5([]int{5, 2, 3, 1}))
	ast.Equal([]int{0, 0, 1, 1, 2, 5}, sortArray5([]int{5, 1, 1, 2, 0, 0}))
}

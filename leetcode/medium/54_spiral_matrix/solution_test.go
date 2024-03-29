package _54_rotate_image

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	ast := assert.New(t)
	nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//ast.Equal(spiralOrder(nums), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	//
	//ast.Equal(spiralOrder2(nums), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	nums = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	ast.Equal(spiralOrder2(nums), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}

// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 0 {
		return []int{1, 1}
	}

	result := make([]int, rowIndex+1)
	result[0] = 1
	result[1] = 1

	for i := 2; i <= rowIndex; i++ {
		backup := 1
		for j := 1; j < i; j++ {
			stash := result[j]
			result[j] += backup
			backup = stash
		}
		result[i] = 1
	}
	return result
}

func TestGetRow(t *testing.T) {
	assert.Equal(t, []int{1, 5, 10, 10, 5, 1}, getRow(5))
	assert.Equal(t, []int{1, 3, 3, 1}, getRow(3))
	assert.Equal(t, []int{1}, getRow(0))
	assert.Equal(t, []int{1, 1}, getRow(1))
}

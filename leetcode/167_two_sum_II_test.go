// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and
// numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers,index1 and index2, added by one as an integer array [index1,index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSumII(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1
	for lp < rp {
		if numbers[lp]+numbers[rp] == target {
			return []int{lp + 1, rp + 1}
		} else if numbers[lp]+numbers[rp] > target {
			rp -= 1
		} else {
			lp += 1
		}
	}
	return nil
}

func TestIsValidTwoSumII(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSumII([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 3}, twoSumII([]int{2, 3, 4}, 6))
	assert.Equal(t, []int{1, 2}, twoSumII([]int{-1, 0}, -1))
	assert.Nil(t, twoSumII([]int{-1, 0}, -2))
	assert.Nil(t, twoSumII([]int{}, 0))
}

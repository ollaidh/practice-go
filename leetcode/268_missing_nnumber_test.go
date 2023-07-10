// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func missingNumber(nums []int) int {
	maximum := len(nums)
	sumPerfect := maximum * (maximum + 1) / 2
	sumReal := 0
	for i := 0; i < maximum; i++ {
		sumReal += nums[i]
	}
	result := sumPerfect - sumReal
	return result
}

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, 2, missingNumber([]int{0, 1}))
	assert.Equal(t, 2, missingNumber([]int{3, 0, 1}))
	assert.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

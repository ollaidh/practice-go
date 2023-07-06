// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// Majority element always exists in the array.
// Solve linear time and in O(1) space

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func majorityElement(nums []int) int {
	counter := 0
	result := 0
	for _, num := range nums {
		if counter == 0 {
			result = num
		}
		if num == result {
			counter += 1
		} else {
			counter -= 1
		}
	}
	return result
}

func TestMajorElement(t *testing.T) {
	assert.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	assert.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	assert.Equal(t, 3, majorityElement([]int{3, 3, 3, 3}))
	assert.Equal(t, 5, majorityElement([]int{5}))
}

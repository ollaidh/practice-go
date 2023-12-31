// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	diff := make(map[int]int)
	for i, num := range nums {
		if _, ok := diff[num]; ok {
			return []int{diff[num], i}
		} else {
			diff[target-num] = i
		}
	}
	return nil
}

func TestIsValidTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
	assert.Nil(t, twoSum([]int{10, 11, 12}, 5))
	assert.Nil(t, twoSum([]int{}, 10))

}

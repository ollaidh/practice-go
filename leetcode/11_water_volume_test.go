// Given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith
// line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form a container,
// such that the container contains the most water. Return the maximum amount of water a container can store.

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	lp, rp := 0, len(height)-1
	maximumArea := findMin(height[rp], height[lp]) * (rp - lp)
	var currArea int

	for lp < rp {
		currArea = findMin(height[rp], height[lp]) * (rp - lp)
		maximumArea = findMax(currArea, maximumArea)
		if height[lp] < height[rp] {
			lp += 1
		} else {
			rp -= 1
		}
	}
	return maximumArea
}

func TestIsValidWater(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
}

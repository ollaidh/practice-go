// Function to find the longest common prefix string amongst an array of strings.
//If there is no common prefix, return an empty string "".

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	var result strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if len(str) <= i {
				return result.String()
			}
			if str[i] != strs[0][i] {
				return result.String()
			}
		}
		result.WriteByte(strs[0][i])
	}
	return result.String()
}

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "flow", "racecar"}))
	assert.Equal(t, "racecar", longestCommonPrefix([]string{"racecar", "racecar", "racecar"}))
	assert.Equal(t, "a", longestCommonPrefix([]string{"ab", "a", "ac"}))
}

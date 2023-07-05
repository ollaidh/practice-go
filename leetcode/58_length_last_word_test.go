// Given a string s consisting of words and spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func lengthOfLastWord(s string) int {
	result := 0
	isWord := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			result += 1
			isWord = true
		}
		if isWord && s[i] == ' ' {
			return result
		}
	}
	return result
}

func TestLengthLastWord(t *testing.T) {
	assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	assert.Equal(t, 4, lengthOfLastWord("   fly me   to   the moon  "))
	assert.Equal(t, 6, lengthOfLastWord("luffy is still joyboy"))
	assert.Equal(t, 4, lengthOfLastWord("aaaa"))
}

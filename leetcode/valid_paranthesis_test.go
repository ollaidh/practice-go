package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func keyInMap(key rune, inputMap map[rune]rune) bool {
	_, isInMap := inputMap[key]
	return isInMap
}

func isValid(s string) bool {
	parenthesis := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if keyInMap(c, parenthesis) {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			if parenthesis[stack[len(stack)-1]] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("(])"))

}

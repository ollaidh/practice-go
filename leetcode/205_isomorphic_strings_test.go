package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isIsomorphic(lhs string, rhs string) bool {
	relationsLhs := make(map[uint8]uint8)
	relationsRhs := make(map[uint8]uint8)
	for i := 0; i < len(lhs); i++ {

		if _, ok := relationsLhs[lhs[i]]; ok == false {
			relationsLhs[lhs[i]] = rhs[i]
		}
		if _, ok := relationsRhs[rhs[i]]; ok == false {
			relationsRhs[rhs[i]] = lhs[i]
		}
		if relationsRhs[rhs[i]] != lhs[i] || relationsLhs[lhs[i]] != rhs[i] {
			return false
		}
	}
	return true
}

func TestIsIsomorphic(t *testing.T) {
	assert.True(t, isIsomorphic("foo", "zaa"))
	assert.True(t, isIsomorphic("qwerty", "asdfgh"))
	assert.True(t, isIsomorphic("fozo", "zara"))
	assert.False(t, isIsomorphic("ooops", "spooo"))
	assert.False(t, isIsomorphic("ywerty", "asdfgh"))
	assert.False(t, isIsomorphic("badc", "baba"))
	assert.True(t, isIsomorphic("paper", "title"))
	assert.True(t, isIsomorphic("q", "q"))
	assert.True(t, isIsomorphic("q", "c"))
}

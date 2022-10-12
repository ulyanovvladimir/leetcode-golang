package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome(121))
	assert.False(t, isPalindrome(12))
	assert.True(t, isPalindrome(1))
	assert.True(t, isPalindrome(0))
	assert.True(t, isPalindrome(123321))
	assert.False(t, isPalindrome(12123))
}

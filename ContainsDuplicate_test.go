package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	assert.True(t, containsDuplicate([]int{1, 2, 3, 1}))
	assert.False(t, containsDuplicate([]int{1, 2, 3, 4}))
	assert.False(t, containsDuplicate([]int{}))
	assert.True(t, containsDuplicate([]int{1, 2, 2, 3}))
}

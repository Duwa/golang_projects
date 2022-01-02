package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T){
	assert.True(t, IsPalindrome("hannah"))
	assert.True(t, IsPalindrome("hanah"))
	assert.True(t, IsPalindrome("hello"))
}

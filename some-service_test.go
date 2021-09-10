package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSomeService(t *testing.T) {
	s := NewSomeService("")
	assert.Equal(t, &SomeService{}, s)
}

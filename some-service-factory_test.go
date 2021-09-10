package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateService(t *testing.T) {
	msg := "this is a test"
	s := CreateService(msg)
	assert.Equal(t, s, &SomeService{msg})
}

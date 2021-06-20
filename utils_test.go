package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIsColumnHide(t *testing.T) {
	expected := true
	args := []string{"color", "test", "size"}
	result := checkIsColumnHide(args, "color")
	assert.Equal(t, expected, result)
}

func TestGetColorTrue(t *testing.T) {
	expected := boldbrightgreen
	result := getColor(true, "dir")
	assert.Equal(t, expected, result)
}

func TestGetColorTFalse(t *testing.T) {
	expected := nc
	result := getColor(true, "")
	assert.Equal(t, expected, result)
}

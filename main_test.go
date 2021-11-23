package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBiggestNumber(t *testing.T) {
	listNumber := []int{1, 2, 3, 8, 9, 3, 2, 1}
	biggestNumber := getBiggestNumber(listNumber)
	assert.Equal(t, 3, biggestNumber)

}

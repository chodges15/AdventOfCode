package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPart1(t *testing.T) {
    assert.Equal(t, 7, Part1("test.txt"))
}

func TestPart2(t *testing.T) {
    assert.Equal(t, 5, Part2("test.txt"))
}
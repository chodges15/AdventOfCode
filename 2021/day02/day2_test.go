package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestDown(t *testing.T) {
    c := []command{
        {
            Direction: down,
            Distance:  1,
        },
        {
            Direction: forward,
            Distance:  1,
        },
        {
            Direction: down,
            Distance:  1,
        },
        {
            Direction: up,
            Distance:  1,
        },
    }
    x, y := move1(c)
    assert.Equal(t, y, 1)
    assert.Equal(t, x, 1)
}
func TestPart1(t *testing.T) {
    assert.Equal(t, 150, Execute(false, "test.txt"))
}

func TestPart2(t *testing.T) {
    assert.Equal(t, 900, Execute(true, "test.txt"))
}

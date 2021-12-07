package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var markedSpace = bingoSpace{
    isMarked: true,
    number:   0,
}
var unmarkedSpace = bingoSpace{
    isMarked: false,
    number:   0,
}

func TestRowComplete(t *testing.T) {
    card := [][]bingoSpace{
        {markedSpace, markedSpace},
        {unmarkedSpace, unmarkedSpace},
    }

    assert.True(t, cardHasRowCompleted(card))
    assert.False(t, cardHasColumnCompleted(card))
}


func TestColumnComplete(t *testing.T) {
    card := [][]bingoSpace{
        {markedSpace, unmarkedSpace},
        {markedSpace, unmarkedSpace},
    }

    assert.True(t, cardHasColumnCompleted(card))
    assert.False(t, cardHasRowCompleted(card))
}

func TestGetInput(t *testing.T) {
    _, cards := getInput("test.txt")
    assert.Equal(t, 3, len(cards))
}

func TestPart1(t *testing.T) {
    assert.Equal(t, 4512, Part1("test.txt"))
}

func TestPart2(t *testing.T) {
    assert.Equal(t, 1924, Part2("test.txt"))
}

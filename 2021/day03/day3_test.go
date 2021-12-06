package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPart1(t *testing.T) {
  assert.Equal(t, 198, Part1("test.txt"))
}

func TestPart2(t *testing.T) {
    assert.Equal(t, 230, Part2("test.txt"))
}

func TestGamma(t *testing.T) {
   assert.Equal(t, 22, computeVariable(gammaVar, getInput("test.txt")))
}

func TestEpsilon(t *testing.T) {
   assert.Equal(t, 9, computeVariable(epsilonVar, getInput("test.txt")))
}

func TestOxygenGenerator(t *testing.T) {
   assert.Equal(t, 23, computeRating(oxygenGenerator, getInput("test.txt")))
}

func TestCo2Scrubber(t *testing.T) {
   assert.Equal(t, 10, computeRating(co2Scrubber, getInput("test.txt")))
}

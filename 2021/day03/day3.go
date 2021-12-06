package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

func main() {
    fmt.Printf("Part 1: %d\n", Part1("input.txt"))
    fmt.Printf("Part 2: %d\n", Part2("input.txt"))
}

type variable int
const (
    gammaVar variable = iota
    epsilonVar
)

type rating int
const (
    oxygenGenerator rating = iota
    co2Scrubber
)

type comparison int
const (
    less comparison = iota
    more
    equal
)

func Part1(filename string) int {
    in := getInput(filename)

    return computeVariable(gammaVar, in) * computeVariable(epsilonVar, in)
}

func Part2(filename string) int {
    in := getInput(filename)

    return computeRating(oxygenGenerator, in) * computeRating(co2Scrubber, in)
}

func computeVariable(v variable, in [][]bool) int {
    variable := 0.0
    countOf1sComparison := getCountOf1sByColumn(in)
    for i := 0; i < len(countOf1sComparison); i++ {
        if (countOf1sComparison[i] == more && v == gammaVar) ||
            (countOf1sComparison[i] == less && v == epsilonVar) {
            variable += math.Pow(2, float64(len(in[0]) - (i+1)))
        }
    }
    return int(variable)
}

func computeRating(r rating, in [][]bool) int {
    width := len(in[0])

    for i := 0; i < width && len(in) > 1; i++ {
        comparison := getCountOf1sByColumn(in)
        purge1s := false
        purge0s := false
        if comparison[i] == more || (r == oxygenGenerator && comparison[i] == equal) {
            if r == oxygenGenerator {
                purge0s = true
            } else {
                purge1s = true
            }
        } else { // less
            if r == oxygenGenerator || (r == co2Scrubber && comparison[i] == equal){
                purge1s = true
            } else {
                purge0s = true
            }
        }

        newIn := make([][]bool, 0)
        for _, candidateRow := range in {
            addRow := (candidateRow[i] == true && !purge1s) || (candidateRow[i] == false && !purge0s)
            if addRow {
               newIn = append(newIn, candidateRow)
            }
        }
        in = make([][]bool, len(newIn))
        copy(in, newIn)
    }
    rating := 0.0
    for i := 0; i < len(in[0]); i++ {
        if in[0][i] {
            rating += math.Pow(2, float64(len(in[0]) - (i+1)))
        }
    }
    return int(rating)
}

func getCountOf1sByColumn(in [][]bool) []comparison {
    comparison := make([]comparison, 0)
    if len(in) == 0 {
        panic("Uh oh")
    }
    width := len(in[0])
    height := len(in)
    for i := 0; i < width; i++ {
        countOf1s := 0
        for j := 0; j < height; j++ {
            if in[j][i] == true {
                countOf1s += 1
            }
        }

        countOf0s := height - countOf1s

        if countOf1s < countOf0s {
            comparison = append(comparison, less)
        } else if countOf1s == countOf0s {
            comparison = append(comparison, equal)
        } else {
            comparison = append(comparison, more)
        }
    }

    return comparison
}



func getInput(filename string)  [][]bool {
    table := make([][]bool, 0)
    if f, err := os.Open(filename); err == nil {
        defer f.Close()
        scanner := bufio.NewScanner(f)

        for scanner.Scan() {
            raw := scanner.Text()
            row := make([]bool, 0)
            for _, b := range raw {
                value := true
                if string(b) == "0" {
                    value = false
                }
                row = append(row, value)
            }
            table = append(table, row)
        }
    } else {
        panic("File error")
    }
    return table
}

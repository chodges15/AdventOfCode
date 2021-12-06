package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    fmt.Printf("Part 1: %d\n", Part1("part1.txt"))
    fmt.Printf("Part 2: %d\n", Part2("part1.txt"))
}

func getInput(filename string) []int {
    if f, err := os.Open(filename); err == nil {
        defer f.Close()

        scanner := bufio.NewScanner(f)

        var inputArray []int
        for scanner.Scan() {
            if number, err := strconv.Atoi(scanner.Text()); err == nil {
                inputArray = append(inputArray, number)
            }
        }
        return inputArray
    } else {
        return []int{}
    }
}

func Part1(filename string) int {
    increases := 0
    inputArray := getInput(filename)
    lastNumber := 0
    for i, num := range inputArray {
        if i != 0 && num > lastNumber {
            increases += 1
        }
        lastNumber = num
    }

    return increases
}

func Part2(filename string) int {
    increases := 0
    inputArray := getInput(filename)
    lastWindow := 0
    for i, _ := range inputArray {
        if i > len(inputArray) - 3 {
            continue
        }
        sum := inputArray[i] + inputArray[i+1] + inputArray[i+2]
        if i != 0 && sum > lastWindow {
            increases += 1
        }
        lastWindow = sum
    }

    return increases
}

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type command struct {
    Direction string
    Distance  int
}

const (
    forward string = "forward"
    down = "down"
    up = "up"
)

func main() {
    fmt.Printf("Part 1: %d\n", Execute(false, "part1.txt"))
    fmt.Printf("Part 2: %d\n", Execute(true, "part1.txt"))
}

func Execute(isPart2 bool, filename string) int {
    in := getInput(filename)
    var x, y int
    if(isPart2) {
       x, y = move2(in)
    } else {
       x, y = move1(in)
    }
    return x*y
}

func move1(commands []command) (int, int) {
    var position, depth int
    for _, c := range commands {
        switch c.Direction {
        case forward:
            position += c.Distance
        case down:
            depth += c.Distance
        case up:
            depth -= c.Distance
        }
    }

    return position, depth
}

func move2(commands []command) (int, int) {
    var position, depth, aim int
    for _, c := range commands {
        switch c.Direction {
        case forward:
            position += c.Distance
            depth += aim * c.Distance
        case down:
            aim += c.Distance
        case up:
            aim -= +c.Distance
        }
    }

    return position, depth
}


func getInput(filename string) []command {
    if f, err := os.Open(filename); err == nil {
        defer f.Close()
        scanner := bufio.NewScanner(f)

        var inputArray []command
        for scanner.Scan() {
            commandRaw := strings.Split(scanner.Text(), " ")
            if number, err := strconv.Atoi(commandRaw[1]); err == nil {
                var c = command{
                    Direction: commandRaw[0],
                    Distance:  number,
                }
                inputArray = append(inputArray, c)
            }
        }
        return inputArray
    } else {
        return []command{}
    }
}

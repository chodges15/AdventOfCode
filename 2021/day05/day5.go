package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {

}

func getInput(filename string) [][][]int {
    listOfLines := make([][][]int, 0)

    if f, err := os.Open(filename); err == nil {
        defer f.Close()
        scanner := bufio.NewScanner(f)

        for scanner.Scan() {
            raw := scanner.Text()
            line := make([][]int, 0)
            coords := strings.Split(raw, "->")
            for _, coord := range coords {
                points := make([]int, 0)
                pointsRaw := strings.Split(coord, ",")
                for _, point := range pointsRaw {
                    num, err := strconv.Atoi(strings.TrimSpace(point))
                    if err != nil {
                        panic(err)
                    }
                    points = append(points, num)
                }
                line = append(line, points)
            }
            listOfLines = append(listOfLines, line)
        }
    } else {
        panic("File error")
    }

    return listOfLines
}
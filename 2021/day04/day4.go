package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type bingoSpace struct {
    isMarked bool
    number int
}
type bingoCard [][]bingoSpace

func main() {
    fmt.Printf("Part 1: %d\n", Part1("input.txt"))
    fmt.Printf("Part 2: %d\n", Part2("input.txt"))
}

func Part1(filename string) int {
    calls, cards := getInput(filename)
    finalCallIdx, winningCardIdx := playBingo(calls, cards, true)
    if finalCallIdx == -1 || winningCardIdx == -1 {
        panic("No solution found")
    }
    return sumOfUnmarked(cards[winningCardIdx]) * calls[finalCallIdx]
}

func Part2(filename string) int {
    calls, cards := getInput(filename)
    finalCallIdx, winningCardIdx := playBingo(calls, cards, false)
    if finalCallIdx == -1 || winningCardIdx == -1 {
        panic("No solution found")
    }
    return sumOfUnmarked(cards[winningCardIdx]) * calls[finalCallIdx]
}

func playBingo(calls []int, cards []bingoCard, stopAtWin bool) (int, int) {
    winners := make(map[int]interface{}, 0)
    for callIdx, call := range calls {
        for cardIdx, card := range cards {
            markCard(card, call)
            if isWinner(card)  {
                winners[cardIdx] = struct{}{}
                if stopAtWin || (!stopAtWin && len(winners) == len(cards)) {
                    return callIdx, cardIdx
                }
            }
        }
    }
    return -1, -1
}

func sumOfUnmarked(card bingoCard) int {
    sum := 0
    for _, row := range card {
        for _, space := range row {
            if !space.isMarked  {
                sum += space.number
            }
        }
    }
    return sum
}

func isWinner(card bingoCard) bool {
   return cardHasRowCompleted(card) || cardHasColumnCompleted(card)
}

func cardHasRowCompleted(card bingoCard) bool {
    for _, row := range card {
        length := len(row)
        markCount := 0
        for _, space := range row {
            if space.isMarked  {
                markCount++
            }
            if markCount == length {
                return true
            }
        }
    }
    return false
}

func cardHasColumnCompleted(card bingoCard) bool {
    for i := 0; i < len(card[0]); i++ {
        markCount := 0
        length := len(card)
        for j := 0; j < len(card); j++ {
            if card[j][i].isMarked {
                markCount++
            }
        }
        if markCount == length {
            return true
        }
    }
    return false
}


func markCard(card bingoCard, call int)  {
    for j, row := range card {
        for i, space := range row {
            if space.number == call {
                space.isMarked = true
                card[j][i] = space
            }
        }
    }
}

func getInput(filename string)  ([]int, []bingoCard) {
    cards := make([]bingoCard, 0)
    bingoCalls := make([]int, 0)
    if f, err := os.Open(filename); err == nil {
        defer f.Close()
        scanner := bufio.NewScanner(f)

        // Calls
        scanner.Scan()
        bingoCallsText := strings.Split(scanner.Text() ,",")
        for _, call := range bingoCallsText {
            num, _ := strconv.Atoi(call)
            bingoCalls = append(bingoCalls, num)
        }

        // Empty Line
        scanner.Scan()

        // Bingo cards
        card := make([][]bingoSpace, 0)
        for scanner.Scan() {
            raw := strings.Split(scanner.Text(), " ")
            if len(raw) < 5 {
                cards = append(cards, card)
                card = make([][]bingoSpace, 0)
                continue
            }
            row := make([]bingoSpace, 0)
            for _, b := range raw {
                num, err := strconv.Atoi(b)
                if err == nil {
                    row = append(row, bingoSpace{
                        isMarked: false,
                        number:   num,
                    })
                }
            }
            card = append(card, row)
        }
    } else {
        panic("File error")
    }
    return bingoCalls, cards
}

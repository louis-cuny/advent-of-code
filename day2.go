package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func checkPart1(min int, max int, letter string, pass string, count *int) {
    nb := strings.Count(pass, letter)
    if nb >= min && nb <= max {
        *count++
    }
}

func checkPart2(pos1 int, pos2 int, letter string, pass string, count *int) {
    nbSuccess := 0

    passs := []rune(pass)
    char := []rune(letter)

    if passs[pos1] == char[0] {
       nbSuccess++
    }
    if passs[pos2] == char[0] {
       nbSuccess++
    }

    if nbSuccess == 1 {
        *count++
    }
}

func main() {
    countPart1 := 0
    countPart2 := 0

    filename := "ressources/day2.txt"
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error")
    }

    values := strings.Split(string(content), "\n")
    for _, value := range values {
        policiesPasses := strings.Split(value, ": ")

        if len(policiesPasses) < 2 {
            continue
        }

        rangesLetters := strings.Split(policiesPasses[0], " ")
        minMax := strings.Split(rangesLetters[0], "-")

        // Passwords: policiesPasses[1]
        // Policy min: minMax[0]
        // Policy max: minMax[1]
        // Policy letter: rangesLetters[1]

        min, err := strconv.Atoi(minMax[0])
        if err != nil {
            continue
        }

        max, err := strconv.Atoi(minMax[1])
        if err != nil {
            continue
        }

        checkPart1(min, max, rangesLetters[1], policiesPasses[1], &countPart1)
        checkPart2(min-1, max-1, rangesLetters[1], policiesPasses[1], &countPart2)
    }

    fmt.Println(countPart1)
    fmt.Println(countPart2)
}

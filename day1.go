package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func getPart1(target int, values []string) int {
    for _, element := range values {
        intEl, err := strconv.Atoi(element)
        if err != nil {
            continue
        }
        for _, element2 := range values {
            intEl2, err := strconv.Atoi(element2)
            if err != nil {
                continue
            }
            if intEl+intEl2 == target {
                return intEl*intEl2
            }
        }
    }
    return 0
}

func getPart2(target int, values []string) int {
    for _, element := range values {
        intEl, err := strconv.Atoi(element)
        if err != nil {
             continue
        }
        for _, element2 := range values {
            intEl2, err := strconv.Atoi(element2)
            if err != nil {
                continue
            }
            for _, element3 := range values {
                intEl3, err := strconv.Atoi(element3)
                if err != nil {
                    continue
                }
                if intEl+intEl2+intEl3 == target {
                    return intEl*intEl2*intEl3
                }
            }
        }
    }
    return 0
}

func main() {
    target := 2020

    filename := "resources/day1.txt"
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error")
    }

    values := strings.Split(string(content), "\n")

    part1 := getPart1(target, values)
    part2 := getPart2(target, values)

    fmt.Println(part1)
    fmt.Println(part2)
}

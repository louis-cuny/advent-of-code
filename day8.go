package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var accValue = 0

func intInArray(ints []int, search int) bool {
	for _, value := range ints {
		if value == search {
			return true
		}
	}

	return false
}

func runPart1(lines []string) {
	var rd []int
	cursor := 0

	for i := 1; i > 0; i++ {
		if intInArray(rd, cursor) {
			return
		}
		rd = append(rd, cursor)

		cLine := lines[cursor]
		amount, _ := strconv.Atoi(cLine[5:])
		switch rune(cLine[0]) {
		case 'a':
			cursor++
			if rune(cLine[4]) == '+' {
				accValue += amount
			} else {
				accValue -= amount
			}
			break

		case 'j':
			if rune(cLine[4]) == '+' {
				cursor += amount
			} else {
				cursor -= amount
			}
			break

		case 'n':
			cursor++
		}
	}
}

func runPart2(lines []string) int {
	for f := 0; f < len(lines); f++ {
		if rune(lines[f][0]) == 'a' {
			continue
		}

		acc := 0
		cursor := 0
		var rd []int
		for i := 0; i >= 0; i++ {
			if cursor == len(lines) {
				return acc
			}

			if intInArray(rd, cursor) || cursor > len(lines) {
				break
			}

			rd = append(rd, cursor)

			cLine := lines[cursor]
			amount, _ := strconv.Atoi(cLine[5:])

			bleu := rune(cLine[0])
			if f == cursor {
				if bleu == 'n' {
					bleu = 'j'
				} else {
					bleu = 'n'
				}
			}

			switch bleu {
			case 'a':
				cursor++
				if rune(cLine[4]) == '+' {
					acc += amount
				} else {
					acc -= amount
				}
				break

			case 'j':
				if rune(cLine[4]) == '+' {
					cursor += amount
				} else {
					cursor -= amount
				}
				break

			case 'n':
				cursor++
			}
		}

	}
	return 0
}

func main() {
	filename := "resources/day8.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	lines := strings.Split(string(content), "\n")

	runPart1(lines)
	fmt.Println(accValue)


	fmt.Println(runPart2(lines))
}

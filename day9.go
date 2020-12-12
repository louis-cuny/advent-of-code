package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func isValidP1(lines []int, target int) bool {
	for _, line := range lines {
		for _, lineb := range lines {
			if line + lineb == target && line != lineb {
				return true
			}
		}
	}
	
	return false
}

func d9Part1(lines []int) int {
	for i := 25; i > -1; i++ {
		if !isValidP1(lines[i-25:i], lines[i]) {
			return lines[i]
		}
	}

	return 0
}

func handleRes(numbers []int) int {
	min := numbers[0]
	max := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	return min + max
}

func findP2(lines []int, nb int, target int) []int {
	for i, _ := range lines {
		if i + nb > len(lines) {
			break
		}

		sum := 0
		for n := nb-1; n >= 0; n-- {
			sum += lines[i+n]
		}

		if sum == target {
			var res []int
			for n := nb-1; n >= 0; n-- {
				res = append(res, lines[i+n])
			}

			return res
		}
	}

	return nil
}

func d9Part2(lines []int, target int) int {
	for nbDigits := 2; nbDigits > -1; nbDigits++ {
		res := findP2(lines, nbDigits, target)
		if res == nil {
			continue
		}

		return handleRes(res)
	}

	return 0
}

func main() {
	filename := "resources/day9.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	lines := strings.Split(string(content), "\n")
	var numbers []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	d9part1 := d9Part1(numbers)
	fmt.Println(d9part1)

	d9part2 := d9Part2(numbers, d9part1)
	fmt.Println(d9part2)
}

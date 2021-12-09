package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPart1D3(values []string) int {
	var sums [12]int
	for _, element := range values {
		for f := 0; f < 12; f++ {
			if rune(element[f]) == '1' {
				sums[f]++
			} else {
				sums[f]--
			}
		}
	}
	var tmpG string
	var tmpE string
	for _, item := range sums {
		if item > 0 {
			tmpG += "1"
			tmpE += "0"
		} else {
			tmpG += "0"
			tmpE += "1"
		}
	}
	gamma, err := strconv.ParseInt(tmpG, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	epsilon, err := strconv.ParseInt(tmpE, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	return int(gamma) * int(epsilon)
}

func o2(values []string, keys []int, col int) string {
	if len(keys) == 1 {
		return values[keys[0]]
	}

	var ones []int
	var zeros []int
	for _, key := range keys {
		if rune(values[key][col]) == '1' {
			ones = append(ones, key)
		} else {
			zeros = append(zeros, key)
		}
	}
	col++
	if len(ones) >= len(zeros) {
		return o2(values, ones, col)
	} else {
		return o2(values, zeros, col)
	}
}

func co2(values []string, keys []int, col int) string {
	if len(keys) == 1 {
		return values[keys[0]]
	}

	var ones []int
	var zeros []int
	for _, key := range keys {
		if rune(values[key][col]) == '1' {
			ones = append(ones, key)
		} else {
			zeros = append(zeros, key)
		}
	}
	col++
	if len(ones) < len(zeros) {
		return co2(values, ones, col)
	} else {
		return co2(values, zeros, col)
	}
}

func getPart2D3(values []string) int {
	keys := makeRange(0, len(values)-1)
	o2, err := strconv.ParseInt(o2(values, keys, 0), 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	co2, err := strconv.ParseInt(co2(values, keys, 0), 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	return int(o2) * int(co2)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {

	filename := "resources/day3.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
	}

	values := strings.Split(string(content), "\n")

	part1 := getPart1D3(values)
	part2 := getPart2D3(values)

	fmt.Println(part1)
	fmt.Println(part2)
}

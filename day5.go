package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func getPart1D5(directions [500][4]int) int {
	var res = make(map[int]map[int]int)
	for _, dir := range directions {
		if dir[0] == dir[2] { // Vertical
			a := dir[1]
			b := dir[3]
			if dir[1] > dir[3] {
				a = dir[3]
				b = dir[1]
			}
			for i := a; i <= b; i++ {
				if _, ok := res[dir[0]]; ok {
					if _, ok := res[dir[0]][i]; ok {
						res[dir[0]][i]++
					} else {
						res[dir[0]][i] = 1
					}
				} else {
					res[dir[0]] = map[int]int{}
					res[dir[0]][i] = 1
				}
			}
		} else if dir[1] == dir[3] { // Horizontal
			a := dir[0]
			b := dir[2]
			if dir[0] > dir[2] {
				a = dir[2]
				b = dir[0]
			}
			for i := a; i <= b; i++ {
				if _, ok := res[i]; ok {
					if _, ok := res[i][dir[1]]; ok {
						res[i][dir[1]]++
					} else {
						res[i][dir[1]] = 1
					}
				} else {
					res[i] = map[int]int{}
					res[i][dir[1]] = 1
				}
			}

		}
	}

	return countMultiple(res)
}

func countMultiple(res map[int]map[int]int) int {
	total := 0
	for xK, _ := range res {
		for _, y := range res[xK] {
			if y > 1 {
				total++
			}
		}

	}
	return total
}

func getPart2D5(directions [500][4]int) int {
	var res = make(map[int]map[int]int)
	for _, dir := range directions {
		if dir[0] == dir[2] { // Vertical
			a := dir[1]
			b := dir[3]
			if dir[1] > dir[3] {
				a = dir[3]
				b = dir[1]
			}
			for i := a; i <= b; i++ {
				if _, ok := res[dir[0]]; ok {
					if _, ok := res[dir[0]][i]; ok {
						res[dir[0]][i]++
					} else {
						res[dir[0]][i] = 1
					}
				} else {
					res[dir[0]] = map[int]int{}
					res[dir[0]][i] = 1
				}
			}
		} else if dir[1] == dir[3] { // Horizontal
			a := dir[0]
			b := dir[2]
			if dir[0] > dir[2] {
				a = dir[2]
				b = dir[0]
			}
			for i := a; i <= b; i++ {
				if _, ok := res[i]; ok {
					if _, ok := res[i][dir[1]]; ok {
						res[i][dir[1]]++
					} else {
						res[i][dir[1]] = 1
					}
				} else {
					res[i] = map[int]int{}
					res[i][dir[1]] = 1
				}
			}
		} else { // Diagonal
			if dir[0] > dir[2] {
				if dir[1] < dir[3] { //-+
					y := dir[1]
					for x := dir[0]; x >= dir[2]; x-- {
						if _, ok := res[x]; ok {
							if _, ok := res[x][y]; ok {
								res[x][y]++
							} else {
								res[x][y] = 1
							}
						} else {
							res[x] = map[int]int{}
							res[x][y] = 1
						}
						y++
					}
				} else { // --
					y := dir[1]
					for x := dir[0]; x >= dir[2]; x-- {
						if _, ok := res[x]; ok {
							if _, ok := res[x][y]; ok {
								res[x][y]++
							} else {
								res[x][y] = 1
							}
						} else {
							res[x] = map[int]int{}
							res[x][y] = 1
						}
						y--
					}
				}
			}
			if dir[1] > dir[3] { //+-
				y := dir[1]
				for x := dir[0]; x <= dir[2]; x++ {
					if _, ok := res[x]; ok {
						if _, ok := res[x][y]; ok {
							res[x][y]++
						} else {
							res[x][y] = 1
						}
					} else {
						res[x] = map[int]int{}
						res[x][y] = 1
					}
					y--
				}
			} else { // ++
				y := dir[1]
				for x := dir[0]; x <= dir[2]; x++ {
					if _, ok := res[x]; ok {
						if _, ok := res[x][y]; ok {
							res[x][y]++
						} else {
							res[x][y] = 1
						}
					} else {
						res[x] = map[int]int{}
						res[x][y] = 1
					}
					y++
				}
			}
		}
	}

	return countMultiple(res)
}

func main() {
	filename := "resources/day5.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error1")
	}

	directions := [500][4]int{}
	rawDirs := strings.Split(string(content), "\n")
	for key, rawDir := range rawDirs {

		re := regexp.MustCompile(",| -> ")
		split := re.Split(rawDir, -1)
		for i := 0; i < 4; i++ {
			value, _ := strconv.Atoi(split[i])
			directions[key][i] = value
		}
	}

	part1 := getPart1D5(directions)
	part2 := getPart2D5(directions)
	fmt.Println(part1)
	fmt.Println(part2)
}

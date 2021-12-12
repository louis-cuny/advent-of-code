package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPart1D4(boards [100][5][5]int, inputs [100]int) int {
	for key, val := range inputs {
		if key < 4 {
			continue
		}

		for _, board := range boards {
			for i := 0; i < 5; i++ {
				if inSlice(board[0][i], inputs[:key+1]) && inSlice(board[1][i], inputs[:key+1]) && inSlice(board[2][i], inputs[:key+1]) && inSlice(board[3][i], inputs[:key+1]) && inSlice(board[4][i], inputs[:key+1]) {
					return getBoardScore(board, inputs[:key+1]) * val
				}
				if inSlice(board[i][0], inputs[:key+1]) && inSlice(board[i][1], inputs[:key+1]) && inSlice(board[i][2], inputs[:key+1]) && inSlice(board[i][3], inputs[:key+1]) && inSlice(board[i][4], inputs[:key+1]) {
					return getBoardScore(board, inputs[:key+1]) * val
				}
			}
		}
	}

	return 0
}

func getPart2D4(boards [100][5][5]int, inputs [100]int) int {
	for key, _ := range inputs {
		if key < 4 {
			continue
		}

		wBoardIds := getWinningBoards(inputs[:key+1], boards)
		if len(wBoardIds) == 99 {
			for i := 1; i < 101; i++ {
				if !inSlice(i, wBoardIds) {
					for key2, val := range inputs {
						for j := 0; j < 5; j++ {
							if inSlice(boards[i][0][j], inputs[:key2+1]) && inSlice(boards[i][1][j], inputs[:key2+1]) && inSlice(boards[i][2][j], inputs[:key2+1]) && inSlice(boards[i][3][j], inputs[:key2+1]) && inSlice(boards[i][4][j], inputs[:key2+1]) {
								return getBoardScore(boards[i], inputs[:key2+1]) * val
							}
							if inSlice(boards[i][j][0], inputs[:key2+1]) && inSlice(boards[i][j][1], inputs[:key2+1]) && inSlice(boards[i][j][2], inputs[:key2+1]) && inSlice(boards[i][j][3], inputs[:key2+1]) && inSlice(boards[i][j][4], inputs[:key2+1]) {
								return getBoardScore(boards[i], inputs[:key2+1]) * val
							}
						}
					}
				}
			}
		}
	}

	return 0
}

func getWinningBoards(inputs []int, boards [100][5][5]int) []int {
	var wBoards []int
	for bid, board := range boards {
		for i := 0; i < 5; i++ {
			if inSlice(board[0][i], inputs) && inSlice(board[1][i], inputs) && inSlice(board[2][i], inputs) && inSlice(board[3][i], inputs) && inSlice(board[4][i], inputs) {
				wBoards = append(wBoards, bid)
			} else if inSlice(board[i][0], inputs) && inSlice(board[i][1], inputs) && inSlice(board[i][2], inputs) && inSlice(board[i][3], inputs) && inSlice(board[i][4], inputs) {
				wBoards = append(wBoards, bid)
			}
		}
	}

	return removeDuplicateInt(wBoards)
}

func inSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getBoardScore(board [5][5]int, inputs []int) int {
	unmarked := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !inSlice(board[r][c], inputs) {
				unmarked += board[r][c]
			}
		}
	}
	return unmarked
}

func removeDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	var list []int
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func main() {
	input := [100]int{50, 68, 2, 1, 69, 32, 87, 10, 31, 21, 78, 23, 62, 98, 16, 99, 65, 35, 27, 96, 66, 26, 74, 72, 45, 52, 81, 60, 38, 57, 54, 19, 18, 77, 71, 29, 51, 41, 22, 6, 58, 5, 42, 92, 85, 64, 94, 12, 83, 11, 17, 14, 37, 36, 59, 33, 0, 93, 34, 70, 97, 7, 76, 20, 3, 88, 43, 47, 8, 79, 80, 63, 9, 25, 56, 75, 15, 4, 82, 67, 39, 30, 89, 86, 46, 90, 48, 73, 91, 55, 95, 28, 49, 61, 44, 84, 40, 53, 13, 24}
	filename := "resources/day4.txt"
	content, _ := ioutil.ReadFile(filename)
	boards := [100][5][5]int{}
	rawBoards := strings.Split(string(content), "\n\n")
	for bid, rawBoard := range rawBoards {
		lines := strings.Split(rawBoard, "\n")
		for lid, line := range lines {
			columns := strings.Split(strings.Replace(strings.TrimSpace(line), "  ", " ", 5), " ")
			for cid, column := range columns {
				value, _ := strconv.Atoi(column)
				boards[bid][lid][cid] = value
			}
		}
	}

	part1 := getPart1D4(boards, input)
	part2 := getPart2D4(boards, input)
	fmt.Println(part1)
	fmt.Println(part2)
}

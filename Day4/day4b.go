package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hBingo(mdArray [][]string) (bool) {
	for i, array := range mdArray {
		count := 0
		for j, _ := range array {
			if mdArray[i][j] == "X" {
				count++
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func vBingo(mdArray [][]string) (bool) {
	col1, col2, col3, col4, col5 := 0,0,0,0,0
	for i, array := range mdArray {
		for j, _ := range array {
			if mdArray[i][j] == "X" {
				if (j == 0) {
					col1++
				}
				if (j == 1) {
					col2++
				}
				if (j == 2) {
					col3++
				}
				if (j == 3) {
					col4++
				}
				if (j == 4) {
					col5++
				}
			}
		}
		if col1 == 5 || col2 == 5 || col3 == 5 || col4 == 5 || col5 == 5 {
			return true
		}
	}
	return false
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func sumOfUnmarkedNumbers(mdArray [][]string) int {
	sum := 0
	for i, array := range mdArray {
		for j, _ := range array {
			if mdArray[i][j] != "X" {
				amount, _ := strconv.Atoi(mdArray[i][j])
				sum = sum + amount
			}
		}
	}
	return sum
}

func problem2(scanner *bufio.Scanner) {

	var numbers []string

	lineCount := 0
	iterator := 1

	var values [][]string
	var boards [][][]string

	for scanner.Scan() {
		line := scanner.Text()

		if lineCount == 0 {
			numbers = strings.Split(line, ",")
		}

		if lineCount > 0 {
			if iterator % 6 == 1 {
				iterator++
				continue
			}

			if iterator % 6 != 1 {
				row := strings.Split(line, " ")

				for i, v := range row {
					if v == "" {
						row = remove(row, i)
					}
				}

				for i, v := range row {
					if v == "" {
						row = remove(row, i)
					}
				}

				values = append(values, row)
				if iterator % 6 == 0 {
					boards = append(boards, values)
					values = [][]string{}
				}
				iterator++
			}
		}

		lineCount++
	}

	var boardsHaveWon []bool
	for _, board := range boards {
		fmt.Println(board)
		boardsHaveWon = append(boardsHaveWon, false)
	}

	numberOfBoardsThatHaveWon := 0
	var losingBoard [][]string

	for _, number := range numbers {
		numberOfBoardsThatHaveWon = 0
		for boardNumber, values := range boards {
			for i, array := range values {
				for j, value := range array {
					if number == value {
						values[i][j] = "X"
					}
				}
			}
			bingoH := hBingo(values)
			bingoV := vBingo(values)
			if (bingoH || bingoV) {
				numberOfBoardsThatHaveWon++
				boardsHaveWon[boardNumber] = true
			}

		}
		if numberOfBoardsThatHaveWon == len(boards) - 1 {
			for i, v := range boardsHaveWon {
				if v == false {
					fmt.Println("Index of losing board is")
					fmt.Println(i)
					losingBoard = boards[i]
				}
			}
			break
		}
	}

	var total, finalNumber int
	for _, number := range numbers {
		for i, array := range losingBoard {
			for j, value := range array {
				if number == value {
					losingBoard[i][j] = "X"
				}
			}
		}
		bingoH := hBingo(losingBoard)
		bingoV := vBingo(losingBoard)
		if number == "13" {
			fmt.Println(losingBoard)
			fmt.Println(bingoV)
		}
		if bingoH || bingoV {
			total = sumOfUnmarkedNumbers(losingBoard)
			finalNumber, _ =  strconv.Atoi(number)
			break
		}
	}

	fmt.Println("Sum of unmarked numbers")
	fmt.Println(total)
	fmt.Println("Number just called")
	fmt.Println(finalNumber)
	fmt.Println("Product")
	fmt.Println(total * finalNumber)
}

func main() {
	f, _ := os.Open("Day4/day4.txt")

	scanner2 := bufio.NewScanner(f)

	problem2(scanner2)
}

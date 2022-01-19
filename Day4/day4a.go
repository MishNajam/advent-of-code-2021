package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bingoHorizontal(mdArray [][]string) (bool,int) {
	for i, array := range mdArray {
		count := 0
		for j, _ := range array {
			if mdArray[i][j] == "X" {
				count++
			}
		}
		if count == 5 {
			return true, i
		}
	}
	return false, -1
}

func bingoVertical(mdArray [][]string) (bool,int) {
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
			return true, i
		}
		if (i+1) % 5 == 0 {
			col1 = 0
			col2 = 0
			col3 = 0
			col4 = 0
			col5 = 0
		}
	}
	return false, -1
}

func calculateBoardRemainder(mdArray [][]string, start int, stop int) (int) {
	sum := 0
	for i, array := range mdArray {
		for j, _ := range array {
			//if i >= 10 <= 14
			if i >= start && i <= stop {
				if mdArray[i][j] != "X" {
					amount, _ :=  strconv.Atoi(mdArray[i][j])
					sum = sum + amount
				}
			}
		}
	}
	return sum
}

func problem1(scanner *bufio.Scanner) {

	var numbers []string

	lineCount := 0
	iterator := 1

	var values [][]string

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
				//fmt.Println(row)
				values = append(values, row)
				iterator++
			}
		}

		lineCount++
	}

	var boardNumber, endOfBoardIndex, startOfBoardIndex, finalNumber int

	for _, number := range numbers {
		for i, array := range values {
			for j, value := range array {
				if number == value {
					values[i][j] = "X"
				}
			}
		}

		bingoH, at := bingoHorizontal(values)
		if bingoH == true {
			bingo := fmt.Sprintf("Horizontal bingo with number %s and index %d value at %s", number, at, values[at])
			fmt.Println(bingo)
			finalNumber, _ =  strconv.Atoi(number)
			boardNumber = (at/5) + 1
			endOfBoardIndex = (boardNumber * 5) - 1
			startOfBoardIndex = endOfBoardIndex - 4
			break
		}

		bingoV, at := bingoVertical(values)
		if bingoV == true {
			fmt.Println("Bingo vertical")
			fmt.Println("Number is")
			fmt.Println(number)
			finalNumber, _ =  strconv.Atoi(number)
			fmt.Println(at)
			fmt.Println(values[at])
			boardNumber = (at/5) + 1
			endOfBoardIndex = (boardNumber * 5) - 1
			startOfBoardIndex = endOfBoardIndex - 4
			break
		}
	}

	total := calculateBoardRemainder(values, startOfBoardIndex, endOfBoardIndex)
	fmt.Println("Sum of unmarked numbers")
	fmt.Println(total) //562
	fmt.Println("Number just called")
	fmt.Println(finalNumber) //71
	fmt.Println("Product")
	fmt.Println(total * finalNumber) //39902
}

func main() {

	f, _ := os.Open("Day4/day4.txt")

	scanner := bufio.NewScanner(f)

	problem1(scanner)

}

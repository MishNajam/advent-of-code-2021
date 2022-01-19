package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem2(scanner *bufio.Scanner) {
	maxY := 0
	maxX := 0
	count := 0
	var points [][]int
	var field [][]int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		start := strings.Split(parts[0], ",")
		finish := strings.Split(parts[2], ",")

		x1, _ := strconv.Atoi(start[0])
		x2, _ := strconv.Atoi(finish[0])
		y1, _ := strconv.Atoi(start[1])
		y2, _:= strconv.Atoi(finish[1])

		if maxX < x1 { maxX = x1 }
		if maxX < x2 { maxX = x2 }
		if maxY < y1 { maxY = y1}
		if maxY < y2 { maxY = y2}

		var firstRow = []int{x1, x2, y1, y2}
		points = append(points, firstRow)
		count += 1
	}
	fmt.Println("MaxX is " + strconv.Itoa(maxX) + " MaxY is " +  strconv.Itoa(maxY))

	rowCount := 0
	for rowCount <= maxY {
		fieldRow := make([]int, maxX + 1)
		for i, _ := range fieldRow {
			fieldRow[i] = 0
		}
		field = append(field, fieldRow)
		rowCount++
	}

	for _, point := range points {
		x1 := point[0]
		x2 := point[1]
		y1 := point[2]
		y2 := point[3]

		if x1 == x2 || y1 == y2 {
			for i, array := range field { // col
				for j, _ := range array { //row
					if y1 == y2 && y1 == i {
						if (x1 >= j && x2 <= j) || (x1 <= j && x2 >= j){
							field[i][j] = field[i][j] + 1
						}
					}

					if x1 == x2 && x1 == j {
						if (y1 >= i && y2 <= i) || (y1 <= i && y2 >= i){
							field[i][j] = field[i][j] + 1
						}
					}
				}
			}
		} else {
			xOffSet := -1
			yOffSet := -1
			for i, array := range field {
				if y1 < y2 && i <= y2 && i >= y1 {
					xOffSet += 1
					yOffSet += 1
				}
				if y1 > y2 && (i <= y1 && i >= y2) {
					xOffSet += 1
					yOffSet += 1
				}
				for j, _ := range array {
					if yOffSet > -1 {
						if x1 < x2 {
							if y1 < y2 && i == y1 + yOffSet && y1 + yOffSet <= y2 && j == x1 + xOffSet && x1 + xOffSet <= x2 {
								field[i][j] = field[i][j] + 1
							}
							if y1 > y2 && i == y2 + yOffSet && y2 + yOffSet <= y1 && j == x2 - xOffSet && x2 - xOffSet >= x1 {
								field[i][j] = field[i][j] + 1
							}
						} else if x1 > x2 {

							if y1 < y2 && i == y1 + yOffSet && y1 + yOffSet <= y2 && j == x1 - xOffSet && x1 -xOffSet >= x2 {
								field[i][j] = field[i][j] + 1

							}
							if y1 > y2 && i == y2 + yOffSet && y2 + yOffSet <= y1 && j == x2 + xOffSet && x2 +xOffSet <= x1 {
								field[i][j] = field[i][j] + 1

							}
						}
					}
				}
			}
		}
	}

	dangerSpaces := 0
	for _, rows := range field{
		for _, v := range rows{
			if v > 1 {
				dangerSpaces += 1
			}
		}
	}

	fmt.Println("Lines overlap at least twice " + strconv.Itoa(dangerSpaces) + " times") //19081
}

func main() {

	f, _ := os.Open("Day5/day5.txt")

	scanner2 := bufio.NewScanner(f)

	problem2(scanner2)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem1(scanner *bufio.Scanner) {
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

		if x1 == x2  || y1 == y2 {
			var firstRow = []int{x1, x2, y1, y2}
			points = append(points, firstRow)
			count += 1
		}

	}
	fmt.Println("MaxX is " + strconv.Itoa(maxX) + " MaxY is " +  strconv.Itoa(maxY))

	rows := 0
	for rows <= maxY {
		fieldRow := make([]int, maxX + 1)
		for i, _ := range fieldRow {
			fieldRow[i] = 0
		}
		field = append(field, fieldRow)
		rows++
	}

	for _, point := range points {
		x1 := point[0]
		x2 := point[1]
		y1 := point[2]
		y2 := point[3]

		for i, array := range field {
			for j, _ := range array {
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
	}

	dangerSpaces := 0
	for _, rows := range field{
		for _, v := range rows{
			if v > 1 {
				dangerSpaces += 1
			}
		}
	}

	fmt.Println("Lines overlap at least twice " + strconv.Itoa(dangerSpaces) + " times") //6666
}

func main() {

	f, _ := os.Open("Day5/day5.txt")

	scanner := bufio.NewScanner(f)

	problem1(scanner)

}

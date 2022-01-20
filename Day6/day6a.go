package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem1(scanner *bufio.Scanner) {

	var fishes []int
	days := 80

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		for _, v := range parts {
			fish, _ := strconv.Atoi(v)
			fishes = append(fishes, fish)
		}

	}

	currentDay := 0
	for currentDay < days {
		for i, _ := range fishes {
			if fishes[i] - 1 < 0 {
				fishes[i] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[i] = fishes[i] - 1
			}

		}
		currentDay++
	}
	fmt.Println("Fish count after " + strconv.Itoa(days) + " days is " + strconv.Itoa(len(fishes))) //6666
}

func main() {
	f, _ := os.Open("Day6/day6.txt")
	scanner1 := bufio.NewScanner(f)
	problem1(scanner1)
}
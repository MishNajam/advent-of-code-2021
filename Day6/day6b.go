package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem2(scanner *bufio.Scanner) {

	days := 256

	fishes := make(map[int]int)

	counter := 8
	for counter >= 0 {
		fishes[counter] = 0
		counter--
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		for _, v := range parts {
			fish, _ := strconv.Atoi(v)
			fishes[fish] = fishes[fish] + 1
		}

	}

	currentDay := 0
	for currentDay < days {
		newSet := make(map[int]int)

		newSet[0] = fishes[1]
		newSet[1] = fishes[2]
		newSet[2] = fishes[3]
		newSet[3] = fishes[4]
		newSet[4] = fishes[5]
		newSet[5] = fishes[6]
		newSet[6] = fishes[7] + fishes[0]
		newSet[7] = fishes[8]
		newSet[8] = fishes[0]

		currentDay++
		fishes = newSet
	}

	fishCount := 0
	for i, _ := range fishes {
		fishCount += fishes[i]
	}
	fmt.Println("Fish count after " + strconv.Itoa(days) + " days is " + strconv.Itoa(fishCount)) //6666
}

func main() {
	f, _ := os.Open("Day6/day6.txt")
	scanner2 := bufio.NewScanner(f)
	problem2(scanner2)
}
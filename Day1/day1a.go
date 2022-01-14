package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func problem1(scanner *bufio.Scanner) {
	increaseCount := 0
	decreaseCount := 0
	previousNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		if number == 155 {
			previousNumber = number
			continue
		}

		if number-previousNumber > 0 {
			increaseCount = increaseCount + 1
		} else if number - previousNumber < 0 {
			decreaseCount = decreaseCount + 1
		}

		previousNumber = number
	}
	fmt.Println("Increased count is " + strconv.Itoa(increaseCount)) //1713
	fmt.Println("Decreased count is " + strconv.Itoa(decreaseCount)) //286
}

func main() {
	f, _ := os.Open("Day1/day1.txt")

	scanner := bufio.NewScanner(f)

	problem1(scanner)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem1(scanner *bufio.Scanner) {
	hP := 0
	d := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, _ :=  strconv.Atoi(parts[1])

		if direction == "forward" {
			d += distance
		}

		if direction == "down" {
			hP += distance
		}

		if direction == "up" {
			hP -= distance
		}
	}
	fmt.Println("Distance is " + strconv.Itoa(d)) //2085
	fmt.Println("Horizontal Position is " + strconv.Itoa(hP)) //785
	fmt.Println("Values multiplied is " + strconv.Itoa(d * hP)) //1636725
}

func main() {

	f, _ := os.Open("Day2/day2.txt")

	scanner := bufio.NewScanner(f)

	problem1(scanner)

}

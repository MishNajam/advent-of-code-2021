package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func problem2(scanner *bufio.Scanner) {
	//aim
	a := 0
	//horizontalPosition
	hP := 0
	//depth
	d := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, _ :=  strconv.Atoi(parts[1])

		fmt.Println("Direction is " + direction + " distance is " +  strconv.Itoa(distance))

		if direction == "forward" {
			hP += distance
			d += distance * a
		}

		if direction == "down" {
			a += distance
		}

		if direction == "up" {
			a -= distance
		}
	}
	fmt.Println("Aim is " + strconv.Itoa(a)) //785
	fmt.Println("Depth is " + strconv.Itoa(d)) //898205
	fmt.Println("Horizontal Position is " + strconv.Itoa(hP)) //2085
	fmt.Println("Final horizontal position * final depth " + strconv.Itoa(d * hP)) //1872757425
}

func main() {

	f, _ := os.Open("Day2/day2.txt")

	scanner := bufio.NewScanner(f)

	problem2(scanner)
}

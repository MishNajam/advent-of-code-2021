package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumValues(array []int) (int, int) {
	c0 := 0
	c1 := 0
	for _, v := range array {
		if v == 1 {
			c1 += 1
		} else {
			c0 += 1
		}
	}
	return c0, c1
}

func addValuesToArrays(gammaRate[] string, epsilonRate[]string, c0 int, c1 int) ([]string, []string) {
	if c0 > c1 {
		gammaRate = append(gammaRate, "0")
		epsilonRate = append(epsilonRate, "1")
	} else {
		gammaRate = append(gammaRate, "1")
		epsilonRate = append(epsilonRate, "0")
	}

	return gammaRate, epsilonRate
}

func problem1(scanner *bufio.Scanner) {

	var r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12 []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1, _ := strconv.Atoi(parts[0])
		v2, _ := strconv.Atoi(parts[1])
		v3, _ := strconv.Atoi(parts[2])
		v4, _ := strconv.Atoi(parts[3])
		v5, _ := strconv.Atoi(parts[4])
		v6, _ := strconv.Atoi(parts[5])
		v7, _ := strconv.Atoi(parts[6])
		v8, _ := strconv.Atoi(parts[7])
		v9, _ := strconv.Atoi(parts[8])
		v10, _ := strconv.Atoi(parts[9])
		v11, _ := strconv.Atoi(parts[10])
		v12, _ := strconv.Atoi(parts[11])

		r1 = append(r1, v1)
		r2 = append(r2, v2)
		r3 = append(r3, v3)
		r4 = append(r4, v4)
		r5 = append(r5, v5)
		r6 = append(r6, v6)
		r7 = append(r7, v7)
		r8 = append(r8, v8)
		r9 = append(r9, v9)
		r10 = append(r10, v10)
		r11 = append(r11, v11)
		r12 = append(r12, v12)
	}

	var gammaRate []string
	var epsilonRate []string

	c0, c1 := sumValues(r1)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)
	c0, c1 = sumValues(r2)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)
	c0, c1 = sumValues(r3)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r4)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r5)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r6)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r7)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r8)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r9)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r10)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r11)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	c0, c1 = sumValues(r12)
	gammaRate, epsilonRate = addValuesToArrays(gammaRate, epsilonRate, c0, c1)

	fmt.Println(gammaRate)
	var gR string
	for _, v := range gammaRate {
		gR = gR + v
	}
	fmt.Println(gR)
	gamma, _ := strconv.ParseInt(gR, 2, 64)
	fmt.Println(gamma)

	fmt.Println(epsilonRate)
	var eR string
	for _, v := range epsilonRate {
		eR = eR + v
	}
	fmt.Println(eR)
	epsilon, _ := strconv.ParseInt(eR, 2, 64)
	fmt.Println(epsilon)

	fmt.Println("Power rate is")
	fmt.Println(epsilon * gamma)
}

func main() {

	f, _ := os.Open("Day3/day3.txt")

	scanner := bufio.NewScanner(f)

	problem1(scanner)

}


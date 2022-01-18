package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumValuesAtIndex(array []string, index int) (int, int) {
	c0 := 0
	c1 := 0
	for _, v := range array {
		parts := strings.Split(v, "")

		if parts[index] == "1" {
			c1 += 1
		} else {
			c0 += 1
		}
	}
	return c0, c1
}

func calculateOxygenRating(oxygenRating[] string, c0 int, c1 int) ([]string, int) {
	if c0 > c1 {
		oxygenRating = append(oxygenRating, "0")
		return oxygenRating, 0
	} else {
		//if equal or less than return 1
		//c0 <= c1
		oxygenRating = append(oxygenRating, "1")
		return oxygenRating, 1
	}
}

func calculateScrubberRating(scrubberRating[] string, c0 int, c1 int) ([]string, int) {
	if c0 > c1 {
		scrubberRating = append(scrubberRating, "1")
		return scrubberRating, 1
	} else {
		//if less than return 1
		//c0 = c1 //return 0
		//c0 < c1 //return 1
		scrubberRating = append(scrubberRating, "0")
		return scrubberRating, 0
	}
}

func problem2Oxygen(scanner *bufio.Scanner) {

	var r1 []string
	var r2 []string
	var r3 []string
	var r4 []string
	var r5 []string
	var r6 []string
	var r7 []string
	var r8 []string
	var r9 []string
	var r10 []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		r1 = append(r1, v1)
	}
	var oxygenGeneratorRating []string

	c0, c1 := sumValuesAtIndex(r1, 0)
	var bit1 int
	oxygenGeneratorRating, bit1 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit1)

	h, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(h)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]

		if v1 == strconv.Itoa(bit1) {
			r2 = append(r2, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r2, 1)
	var bit2 int
	oxygenGeneratorRating, bit2 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit2)

	i, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(i)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) {
			r3 = append(r3, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r3, 2)
	var bit3 int
	oxygenGeneratorRating, bit3 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit3)

	j, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(j)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) {
			r4 = append(r4, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r4, 3)
	var bit4 int
	oxygenGeneratorRating, bit4 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit4)

	k, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(k)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) && v4 == strconv.Itoa(bit4) {
			r5 = append(r5, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r5, 4)
	var bit5 int
	oxygenGeneratorRating, bit5 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit5)

	l, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(l)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5) {
			r6 = append(r6, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r6, 5)
	var bit6 int
	oxygenGeneratorRating, bit6 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit6)


	m, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(m)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]
		v6 := parts[5]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5)  && v6 == strconv.Itoa(bit6) {
			r7 = append(r7, line)
		}
	}
	c0, c1 = sumValuesAtIndex(r7, 6)
	var bit7 int
	oxygenGeneratorRating, bit7 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit7)


	n, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(n)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]
		v6 := parts[5]
		v7 := parts[6]


		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5)  && v6 == strconv.Itoa(bit6) &&
			v7 == strconv.Itoa(bit7) {
			r8 = append(r8, line)
		}
	}
	c0, c1 = sumValuesAtIndex(r8, 7)
	var bit8 int
	oxygenGeneratorRating, bit8 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit8)


	o, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(o)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]
		v6 := parts[5]
		v7 := parts[6]
		v8 := parts[7]


		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5)  && v6 == strconv.Itoa(bit6) &&
			v7 == strconv.Itoa(bit7) && v8 == strconv.Itoa(bit8) {
			r9 = append(r9, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r9, 8)
	var bit9 int
	oxygenGeneratorRating, bit9 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit9)


	for _, v := range r9 {
		parts := strings.Split(v, "")
		v9 := parts[8]
		if v9 == strconv.Itoa(bit9) {
			r10 = append(r10, v)
		}
	}

	c0, c1 = sumValuesAtIndex(r10, 9)
	var bit10 int
	oxygenGeneratorRating, bit10 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit10)
	fmt.Println(r10)

	var r11 []string
	//var r12 []string

	for _, v := range r10 {
		parts := strings.Split(v, "")
		v10 := parts[9]
		if v10 == strconv.Itoa(bit10) {
			r11 = append(r11, v)
		}
	}

	c0, c1 = sumValuesAtIndex(r11, 10)
	var bit11 int
	oxygenGeneratorRating, bit11 = calculateOxygenRating(oxygenGeneratorRating, c0, c1)
	fmt.Println(bit11)
	fmt.Println(r11)


	if len(r11) <= 2 {
		fmt.Println(oxygenGeneratorRating)
		var oR string
		for _, v := range oxygenGeneratorRating {
			oR = oR + v
		}
		fmt.Println(oR) //01100011111
		oxygen, _ := strconv.ParseInt(oR, 2, 64)
		fmt.Println("the oxygen rating is")
		fmt.Println(oxygen) //799

		oxygen, _ = strconv.ParseInt("011000111111", 2, 64)
		fmt.Println("the oxygen rating is")
		fmt.Println(oxygen) //799
	}

}

func problem2Scrubber(scanner *bufio.Scanner) {

	var r1 []string
	var r2 []string
	var r3 []string
	var r4 []string
	var r5 []string
	var r6 []string
	var r7 []string
	var r8 []string
	var r9 []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		r1 = append(r1, v1)
	}
	var scrubberRating []string

	c0, c1 := sumValuesAtIndex(r1, 0)
	fmt.Println(c0, c1)
	fmt.Println(" ")
	fmt.Println("expect value 1")
	var bit1 int
	scrubberRating, bit1 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit1)

	h, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(h)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]

		if v1 == strconv.Itoa(bit1) {
			r2 = append(r2, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r2, 1)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 0 and count 246")
	var bit2 int
	scrubberRating, bit2 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit2)

	i, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(i)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) {
			r3 = append(r3, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r3, 2)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 1 and count 117")

	var bit3 int
	scrubberRating, bit3 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit3)

	j, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(j)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) {
			r4 = append(r4, line)
		}
	}


	c0, c1 = sumValuesAtIndex(r4, 3)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 0 and count 54")
	var bit4 int
	scrubberRating, bit4 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit4)

	k, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(k)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) && v4 == strconv.Itoa(bit4) {
			r5 = append(r5, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r5, 4)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 1 and count 22")
	var bit5 int
	scrubberRating, bit5 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit5)

	l, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(l)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5) {
			r6 = append(r6, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r6, 5)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 1 and count 10")
	var bit6 int
	scrubberRating, bit6 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit6)


	m, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(m)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]
		v6 := parts[5]

		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5)  && v6 == strconv.Itoa(bit6) {
			r7 = append(r7, line)
		}
	}
	c0, c1 = sumValuesAtIndex(r7, 6)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 0 and count 4")
	var bit7 int
	scrubberRating, bit7 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit7)


	n, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(n)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]
		v6 := parts[5]
		v7 := parts[6]


		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5)  && v6 == strconv.Itoa(bit6) &&
			v7 == strconv.Itoa(bit7) {
			r8 = append(r8, line)
		}
	}
	c0, c1 = sumValuesAtIndex(r8, 7)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 0 and count 2")
	var bit8 int
	scrubberRating, bit8 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit8)


	o, _ := os.Open("Day3/day3.txt")
	scanner = bufio.NewScanner(o)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")

		v1 := parts[0]
		v2 := parts[1]
		v3 := parts[2]
		v4 := parts[3]
		v5 := parts[4]
		v6 := parts[5]
		v7 := parts[6]
		v8 := parts[7]


		if v1 ==  strconv.Itoa(bit1) && v2 == strconv.Itoa(bit2) && v3 == strconv.Itoa(bit3) &&
			v4 == strconv.Itoa(bit4) && v5 == strconv.Itoa(bit5)  && v6 == strconv.Itoa(bit6) &&
			v7 == strconv.Itoa(bit7) && v8 == strconv.Itoa(bit8) {
			r9 = append(r9, line)
		}
	}

	c0, c1 = sumValuesAtIndex(r9, 8)
	fmt.Println(c0, c1)
	fmt.Println(c0 + c1)
	fmt.Println(" ")
	fmt.Println("expect value 0 and count 2")
	var bit9 int
	scrubberRating, bit9 = calculateScrubberRating(scrubberRating, c0, c1)
	fmt.Println(bit9)
	fmt.Println(r9)


	if len(r9) <= 2 {
		fmt.Println("Scrubber rating")
		fmt.Println(scrubberRating)
		var oR string
		for _, v := range scrubberRating {
			oR = oR + v
		}
		fmt.Println(oR) //01100011111 //101011000
		scrubber, _ := strconv.ParseInt(oR, 2, 64)
		fmt.Println("the scrubber rating is")
		fmt.Println(scrubber) //799 //344

		scrubber, _ = strconv.ParseInt("101011000100", 2, 64)
		fmt.Println("the scrubber rating is")
		fmt.Println(scrubber) //799 //344
	}

}


func main() {
	g1, _ := os.Open("Day3/day3.txt")

	scanner2 := bufio.NewScanner(g1)

	problem2Oxygen(scanner2)


	g2, _ := os.Open("Day3/day3.txt")

	scanner3 := bufio.NewScanner(g2)

	problem2Scrubber(scanner3)

	fmt.Println(1599 * 2756)
}


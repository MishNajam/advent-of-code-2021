package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func problem2(scanner *bufio.Scanner) {
	increaseCount, decreaseCount, iteration := 0, 0, 0
	sumOfA, sumOfB, sumOfC, sumOfD := 0, 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		iteration += 1

		if iteration == 1 {
			sumOfA += number
			continue
		}

		if iteration == 2 {
			sumOfA += number
			sumOfB += number
			continue
		}

		if iteration == 3 {
			sumOfA += number
			sumOfB += number
			sumOfC += number
			continue
		}

		if iteration % 4 == 0 {
			sumOfB += number
			sumOfC += number
			sumOfD = number
			if sumOfB > sumOfA {
				increaseCount += 1
			} else if sumOfB < sumOfA {
				decreaseCount += 1
			}
			continue
		}

		if iteration % 4 == 1 {
			sumOfA = number
			sumOfC += number
			sumOfD += number
			if sumOfC > sumOfB {
				increaseCount += 1
			} else if sumOfC < sumOfB {
				decreaseCount += 1
			}
			continue
		}

		if iteration % 4 == 2 {
			sumOfA += number
			sumOfB = number
			sumOfD += number
			if sumOfD > sumOfC {
				increaseCount += 1
			} else if sumOfD < sumOfC {
				decreaseCount += 1
			}
			continue
		}

		if iteration % 4 == 3 {
			sumOfA += number
			sumOfB += number
			sumOfC = number
			if sumOfA > sumOfD {
				increaseCount += 1
			} else if sumOfA < sumOfD {
				decreaseCount += 1
			}
			continue
		}

	}
	fmt.Println("Increased count is " + strconv.Itoa(increaseCount)) //1734
	fmt.Println("Decreased count is " + strconv.Itoa(decreaseCount)) //223
}

func main() {

	f, _ := os.Open("Day1/day1.txt")

	scanner2 := bufio.NewScanner(f)

	problem2(scanner2)

}

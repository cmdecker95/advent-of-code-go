package day02

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isLineSafe(line []int) bool {
	lineIncreasing := true
	safe := true

	for i := 1; i < len(line); i++ {
		delta := line[i] - line[i-1]
		increasing := delta > 0

		delta = int(math.Abs(float64(delta)))

		if delta < 1 || delta > 3 {
			safe = false
			break
		}

		if i == 1 {
			lineIncreasing = increasing
			continue
		}

		if lineIncreasing != increasing {
			safe = false
			break
		}
	}

	return safe
}

func Part1() int {
	fmt.Println("Day 2: Part 1")

	// Given a list of int arrays, return how many arrays are "safe" (meet the following):
	// - always increasing or always decreasing
	// - delta is between 1 and 3 (inclusive)

	file, _ := os.ReadFile("day02/data.txt")
	text := strings.TrimSpace(string(file))
	lines := strings.Split(text, "\n")

	safeCount := 0

	for _, lineString := range lines {
		line := []int{}
		for _, v := range strings.Fields(lineString) {
			num, _ := strconv.Atoi(v)
			line = append(line, num)
		}

		safe := isLineSafe(line)

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func Part2() int {
	fmt.Println("Day 2: Part 2")

	// Same as before, but an array is still safe if omitting
	// one value makes it safe.

	file, _ := os.ReadFile("day02/data.txt")
	text := strings.TrimSpace(string(file))
	lines := strings.Split(text, "\n")
	// lines := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}

	safeCount := 0

	for _, lineString := range lines {
		line := []int{}
		for _, v := range strings.Fields(lineString) {
			num, _ := strconv.Atoi(v)
			line = append(line, num)
		}

		safe := isLineSafe(line)

		// Second chance (brute force)
		if !safe {
			for i := 0; i < len(line); i++ {
				newLine := []int{}
				for j := 0; j < len(line); j++ {
					if i == j {
						continue
					}
					newLine = append(newLine, line[j])
				}
				if isLineSafe(newLine) {
					safe = true
					break
				}
			}
		}

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1) // 299

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2) // 364
}

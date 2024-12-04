package day04

import (
	"fmt"
	"strings"
)

func Part1() int {
	fmt.Println("Day 4: Part 1")

	test := `
    MMMSXXMASM
    MSAMXMSMSA
    AMXSXMAAMM
    MSAMASMSMX
    XMASAMXAMM
    XXAMMXXAMA
    SMSMSASXSS
    SAXAMASAAA
    MAMMMXMMMM
    MXMXAXMASX
  `

	test = strings.TrimSpace(test)
	lines := strings.Split(test, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		fmt.Println(line)
	}

	return 0
}

func Part2() int {
	fmt.Println("Day 4: Part 2")

	return 0
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1)

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2)
}

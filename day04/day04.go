package day04

import (
	"fmt"
	"os"
	"strings"
)

func Part1() int {
	fmt.Println("Day 4: Part 1")

	file, _ := os.ReadFile("day04/data.txt")
	text := string(file)
	lines := strings.Split(strings.TrimSpace(text), "\n")

	count := 0

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		for j := 0; j < len(line); j++ {
			char := string(line[j])

			if char == "X" {
				// up
				if i >= 3 && string(lines[i-1][j]) == "M" && string(lines[i-2][j]) == "A" && string(lines[i-3][j]) == "S" {
					count++
				}
				// down
				if i <= len(lines)-4 && string(lines[i+1][j]) == "M" && string(lines[i+2][j]) == "A" && string(lines[i+3][j]) == "S" {
					count++
				}
				// left
				if j >= 3 && string(lines[i][j-1]) == "M" && string(lines[i][j-2]) == "A" && string(lines[i][j-3]) == "S" {
					count++
				}
				// right
				if j <= len(line)-4 && string(lines[i][j+1]) == "M" && string(lines[i][j+2]) == "A" && string(lines[i][j+3]) == "S" {
					count++
				}
				// up left
				if i >= 3 && j >= 3 && string(lines[i-1][j-1]) == "M" && string(lines[i-2][j-2]) == "A" && string(lines[i-3][j-3]) == "S" {
					count++
				}
				// up right
				if i >= 3 && j <= len(line)-4 && string(lines[i-1][j+1]) == "M" && string(lines[i-2][j+2]) == "A" && string(lines[i-3][j+3]) == "S" {
					count++
				}
				// down left
				if i <= len(lines)-4 && j >= 3 && string(lines[i+1][j-1]) == "M" && string(lines[i+2][j-2]) == "A" && string(lines[i+3][j-3]) == "S" {
					count++
				}
				// down right
				if i <= len(lines)-4 && j <= len(line)-4 && string(lines[i+1][j+1]) == "M" && string(lines[i+2][j+2]) == "A" && string(lines[i+3][j+3]) == "S" {
					count++
				}
			}
		}
	}

	return count
}

func Part2() int {
	fmt.Println("Day 4: Part 2")

	file, _ := os.ReadFile("day04/data.txt")
	text := string(file)
	lines := strings.Split(strings.TrimSpace(text), "\n")

	count := 0

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		for j := 0; j < len(line); j++ {
			char := string(line[j])

			if char == "A" {
				if i == 0 || j == 0 || i == len(lines)-1 || j == len(line)-1 {
					continue
				}

				downMas := string(lines[i-1][j-1]) == "M" && string(lines[i+1][j+1]) == "S"
				downSam := string(lines[i-1][j-1]) == "S" && string(lines[i+1][j+1]) == "M"

				if !downMas && !downSam {
					continue
				}

				upMas := string(lines[i+1][j-1]) == "M" && string(lines[i-1][j+1]) == "S"
				upSam := string(lines[i+1][j-1]) == "S" && string(lines[i-1][j+1]) == "M"

				if !upMas && !upSam {
					continue
				}

				count++
			}
		}
	}

	return count
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1) // 2599

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2) // 1948
}

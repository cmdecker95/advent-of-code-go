package day06

import (
	"fmt"
	"os"
	"strings"
)

var sample string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func mark(i, j int, grid *[][]int) {
	(*grid)[i][j] = 1
}

func step(curr_i, curr_j int, gridMap, direction string, grid *[][]int) {

	mark(curr_i, curr_j, grid)

}

func getGridSum(grid [][]int) int {
	sum := 0
	for _, row := range grid {
		for _, val := range row {
			sum += val
		}
	}

	return sum
}

func Part1() int {
	file, _ := os.ReadFile("day06/data.txt")
	text := strings.TrimSpace(string(file))
	text = sample
	print(text)

	// Given a map grid with a guard's starting position and direction and obstacle positions,
	// return all points the guard will cover before exiting. They will travel forward,
	// turning right at each obstacle, finishing when exiting the map.

	return getGridSum(grid)
}

func Part2() int {
	file, _ := os.ReadFile("day06/data.txt")
	text := strings.TrimSpace(string(file))
	text = sample
	print(text)
	return 0
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1)

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2)
}

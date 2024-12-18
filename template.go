package main

import (
	"fmt"
	"os"
	"strings"
)

var filepath string = "FILEPATH"
var sample string = "SAMPLE"

func Part1() int {
	file, _ := os.ReadFile(filepath)
	text := strings.TrimSpace(string(file))
	text = sample
	print(text)
	return 0
}

func Part2() int {
	file, _ := os.ReadFile(filepath)
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

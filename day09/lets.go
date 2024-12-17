package day09

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filepath string = "day09/data.txt"
var sample string = "2333133121414131402"

func Part1() int {
	file, _ := os.ReadFile(filepath)
	text := strings.TrimSpace(string(file))
	text = sample

	// text is a "disk map"
	// next is a "file system"
	// final is the "checksum"

	fs := ""
	id := 0
	for i, char := range text {
		isFile := i%2 == 0
		fmt.Printf("%v %v\n", i, rune(char))

		size, _ := strconv.Atoi(string(char))
		for range size {
			if isFile {
				println(string(id))
				fs += string(id)
				id++
			} else {
				fs += "."
			}
		}
	}

	println(fs)

	return 0
}

func Part2() int {
	file, _ := os.ReadFile(filepath)
	text := strings.TrimSpace(string(file))
	text = sample
	println(text)
	return 0
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1)

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2)
}

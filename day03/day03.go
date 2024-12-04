package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sample string = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func Part1() int {
	file, _ := os.ReadFile("day03/data.txt")
	text := strings.TrimSpace(string(file))

	input := text
	sum := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, match := range re.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}

func Part2() int {
	file, _ := os.ReadFile("day03/data.txt")
	text := strings.TrimSpace(string(file))

	input := text
	sum := 0

	do := true

	mulStrings := make(map[int]string)
	muls := []int{}
	donts := []int{}
	dos := []int{}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, match := range re.FindAllStringSubmatchIndex(input, -1) {
		mulString := input[match[0]:match[1]]
		mulStrings[match[0]] = mulString

		muls = append(muls, match[0])
	}

	re = regexp.MustCompile(`don't\(\)`)
	for _, match := range re.FindAllStringSubmatchIndex(input, -1) {
		donts = append(donts, match[0])
	}

	re = regexp.MustCompile(`do\(\)`)
	for _, match := range re.FindAllStringSubmatchIndex(input, -1) {
		dos = append(dos, match[0])
	}

	for {
		if len(muls) == 0 {
			break
		}

		// Find next thing to happen
		if len(dos) > 0 {
			if len(donts) == 0 || dos[0] < donts[0] {
				if dos[0] < muls[0] {
					// do() came first
					do = true
					dos = dos[1:]
					continue
				}
			}
		}

		if len(donts) > 0 {
			if len(dos) == 0 || donts[0] < dos[0] {
				if donts[0] < muls[0] {
					// dont() came first
					do = false
					donts = donts[1:]
					continue
				}
			}
		}

		// mul() came first
		if do {
			r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
			for _, match := range r.FindAllStringSubmatch(mulStrings[muls[0]], -1) {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
			}
		}

		muls = muls[1:]
		continue

	}

	return sum
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1) // 179834255

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2) // 80570939
}

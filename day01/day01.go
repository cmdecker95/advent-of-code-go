package day01

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readArraysFromFile() ([]int, []int) {
	// get file text
	file, _ := os.ReadFile("day01/data.txt")
	text := strings.TrimSpace(string(file))

	// get pairs of numbers
	pairs := strings.Split(text, "\n")

	// initialize arrays
	n := len(pairs)
	left := make([]int, n)
	right := make([]int, n)

	// fill arrays with pairs
	for i, pair := range pairs {
		nums := strings.Fields(pair)
		l, _ := strconv.Atoi(strings.TrimSpace(nums[0]))
		r, _ := strconv.Atoi(strings.TrimSpace(nums[1]))

		left[i], right[i] = l, r
	}

	return left, right
}

func sort(arr []int) []int {
	// Bubble sort
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}

func Part1() int {
	fmt.Println("Day 1: Part 1")

	// Given two arrays, return the sum of the distances between
	// the paired elements of the sorted arrays.

	// 0. Read arrays from file
	left, right := readArraysFromFile()

	// 1. Sort the arrays
	left = sort(left)
	right = sort(right)

	// 2. Take the distances
	distances := []int{}

	for i := range left {
		distance := math.Abs(float64(left[i] - right[i]))
		distances = append(distances, int(distance))
	}

	// 3. Sum the distances
	sum := 0

	for _, v := range distances {
		sum += v
	}

	return sum
}

func Part2() int {
	fmt.Println("Day 1: Part 2")

	// Given two arrays, return the sum of the left list's similarity scores,
	// defined as the number of times the value in the left list appears in the
	// right list.

	// 0. Read arrays from file
	left, right := readArraysFromFile()

	// 1. Track scores and counts, never recounting
	scores := make(map[int]int)
	counts := make(map[int]int)

	for _, l := range left {

		// Count the number of times the value appears in the right list
		if counts[l] == 0 {
			for _, r := range right {
				if l == r {
					counts[l]++
				}
			}
		}

		scores[l] += counts[l] * l
	}

	// 2. Sum the scores
	sum := 0

	for _, v := range scores {
		sum += v
	}

	return sum
}

func Solve() {
	answer1 := Part1()
	fmt.Printf("Answer: %v\n\n", answer1) // 3508942

	answer2 := Part2()
	fmt.Printf("Answer: %v\n\n", answer2) // 26593248
}

package day05

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample string = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func start() ([]string, []string, []string) {
	file, _ := os.ReadFile("day05/data.txt")
	text := strings.TrimSpace(string(file))

	both := strings.Split(text, "\n\n")
	rules := strings.Split(both[0], "\n")
	updates := strings.Split(both[1], "\n")

	validUpdates := []string{}
	invalidUpdates := []string{}

	for _, update := range updates {
		pageStrings := strings.Split(update, ",")
		updateIsValid := true

		for _, pageString := range pageStrings {

			// Check if page passes all rules that apply to it
			for _, rule := range rules {

				// Extract rule
				ruleParts := strings.Split(rule, "|")
				left := ruleParts[0]
				right := ruleParts[1]

				// Rule does not apply to this page
				if pageString != left && pageString != right {
					continue
				}

				// Rule does not apply to this page
				if !(strings.Contains(update, left) && strings.Contains(update, right)) {
					continue
				}

				// Check if page follows rule
				for _, pageString := range pageStrings {
					if pageString == left {
						break
					} else if pageString == right {
						updateIsValid = false
						break
					}
				}

				if !updateIsValid {
					break
				}
			}

			if !updateIsValid {
				break
			}
		}

		if updateIsValid {
			validUpdates = append(validUpdates, update)
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	return validUpdates, invalidUpdates, rules
}

func Part1(validUpdates []string) int {
	answer := 0

	for _, update := range validUpdates {
		nums := strings.Split(update, ",")
		centerIdx := len(nums) / 2
		num, _ := strconv.Atoi(nums[centerIdx])
		answer += num
	}

	return answer
}

func Part2(invalidUpdates []string, rules []string) int {
	answer := 0
	for _, invalidUpdate := range invalidUpdates {
		fmt.Println("Fixing", invalidUpdate)
		pages := strings.Split(invalidUpdate, ",")

		// collect rules that apply to this update
		applies := []string{}
		for _, rule := range rules {
			ruleParts := strings.Split(rule, "|")
			left := ruleParts[0]
			right := ruleParts[1]

			hasLeft := false
			hasRight := false
			for _, page := range pages {
				if page == left {
					hasLeft = true
				}
				if page == right {
					hasRight = true
				}
			}

			if hasLeft && hasRight {
				applies = append(applies, rule)
			}
		}

		// loop until all rules pass
		for {
			allPass := true

			// go over rules, fixing when broken
			for _, rule := range applies {
				ruleParts := strings.Split(rule, "|")
				left := ruleParts[0]
				right := ruleParts[1]

				for _, page := range pages {
					if page == left {
						break
					}
					if page == right {
						// record that one broke
						allPass = false

						// fix broken
						for i, page := range pages {
							if page == right {
								pages[i] = left
							} else if page == left {
								pages[i] = right
							}
						} // done with fix/swap

						break
					}
				} // done checking (and maybe fixing)
			} // done with rules

			if allPass {
				centerIdx := len(pages) / 2
				num, _ := strconv.Atoi(pages[centerIdx])
				answer += num
				break
			} // otherwise, do it again
		} // Done fixing this update -> go to next update
		fmt.Println(pages)
	} // Done checking all updates

	return answer

}

func Solve() {
	validUpdates, invalidUpdates, rules := start()

	answer1 := Part1(validUpdates)
	fmt.Printf("Answer: %v\n\n", answer1)

	answer2 := Part2(invalidUpdates, rules)
	fmt.Printf("Answer: %v\n\n", answer2)
}

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not read file: %v\n", contents)
	}

	sum, err := FindMiddlePagesSum(string(contents))
	if err != nil {
		log.Fatalf("error finding sum: %v\n", err)
	}
	fmt.Printf("Part 1: %d\n", sum)
}

func PageOrderingRules(input string) (map[string][]string, error) {
	result := make(map[string][]string)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		splits := strings.Split(line, "|")
		if len(splits) != 2 {
			return nil, fmt.Errorf("invalid line while processing page ordering rules: %s", line)
		}

		if rule, found := result[splits[1]]; found {
			rule = append(rule, splits[0])
			result[splits[1]] = rule
		} else {
			result[splits[1]] = []string{splits[0]}
		}
	}

	return result, nil
}

func IsManualValid(manual string, rules map[string][]string) bool {
	pages := strings.Split(strings.TrimSpace(manual), ",")

	for i, page := range pages {
		rule, found := rules[page]
		if !found {
			continue
		}
		for j := i + 1; j < len(pages); j++ {
			if slices.Contains(rule, pages[j]) {
				// fmt.Printf("page %v found in rule %v\n", pages[j], rule)
				return false
			}
		}
	}

	return true
}

func FindMiddlePagesSum(input string) (int, error) {
	split := strings.Split(input, "\n\n")
	if len(split) != 2 {
		return 0, fmt.Errorf("invalid input for middle pages sum")
	}

	rules, err := PageOrderingRules(split[0])
	if err != nil {
		return 0, fmt.Errorf("invalid page ordering: %w", err)
	}

	total := 0
	for _, manual := range strings.Split(split[1], "\n") {
		if manual == "" || !IsManualValid(manual, rules) {
			continue
		}

		pages := strings.Split(manual, ",")
		middlePage, err := strconv.Atoi(pages[len(pages)/2])
		if err != nil {
			fmt.Printf("could not get middle page for pages %v: %v\n", pages, err)
		}

		total += middlePage
	}

	return total, nil
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	part1, err := Part1(file)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Println(part1)

	// Part 2
	file.Seek(0, 0) // Go back to the start of the file
	part2, err := Part2(file)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Println(part2)
}

func Part1(r io.Reader) (int, error) {
	lefts := []int{}
	rights := []int{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, "   ")
		left, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("failed to parse left number: %v\n", err)
			continue
		}
		lefts = append(lefts, left)

		right, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("failed to parse right number: %v\n", err)
			continue
		}
		rights = append(rights, right)
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error in scanner: %w", err)
	}

	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})
	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})

	difference := 0
	for i := range lefts {
		difference += max((lefts[i] - rights[i]), (rights[i] - lefts[i]))
	}

	return difference, nil
}

func Part2(r io.Reader) (int, error) {
	lefts := []int{}
	rights := make(map[int]int)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, "   ")
		left, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Printf("failed to parse left number: %v\n", err)
			continue
		}
		lefts = append(lefts, left)

		right, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Printf("failed to parse right number: %v\n", err)
			continue
		}
		rights[right] += 1
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error in scanner: %w", err)
	}

	score := 0
	for _, left := range lefts {
		similarity, present := rights[left]
		if !present {
			continue
		}
		score += left * similarity
	}

	return score, nil
}

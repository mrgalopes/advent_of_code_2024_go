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
}

func Part1(r io.Reader) (int, error) {
	lefts := []int64{}
	rights := []int64{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, "   ")
		left, err := strconv.ParseInt(numbers[0], 10, 32)
		if err != nil {
			fmt.Printf("failed to parse left number: %v\n", err)
			continue
		}
		lefts = append(lefts, left)

		right, err := strconv.ParseInt(numbers[1], 10, 32)
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
		difference += int(max((lefts[i] - rights[i]), (rights[i] - lefts[i])))
	}

	return difference, nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}

	// Part 1
	count := CountXMAS(string(contents))
	fmt.Println(count)

	// Part 2
	count = CountXMAS2(string(contents))
	fmt.Println(count)
}

type Direction struct{ x, y int }

func CountXMAS(input string) int {
	lines := strings.Split(input, "\n")
	directions := []Direction{
		{0, -1},  // Up
		{1, -1},  // Top-right
		{1, 0},   // Right
		{1, 1},   // Bottom-right
		{0, 1},   // Down
		{-1, 1},  // Bottom-left
		{-1, 0},  // Left
		{-1, -1}, // Top-left
	}

	count := 0
	for j, line := range lines {
		for i, r := range line {
			if r == 'X' {
				for _, d := range directions {
					if peek(lines, i, j, 1, d) == 'M' &&
						peek(lines, i, j, 2, d) == 'A' &&
						peek(lines, i, j, 3, d) == 'S' {
						count++
					}
				}
			}
		}
	}

	return count
}

func CountXMAS2(input string) int {
	lines := strings.Split(input, "\n")
	d1s := []Direction{
		{1, -1}, // Top-right
		{-1, 1}, // Bottom-left
	}
	d2s := []Direction{
		{1, 1},   // Bottom-right
		{-1, -1}, // Top-left
	}

	count := 0
	for j, line := range lines {
		for i, r := range line {
			if r == 'A' {
				d1 := (peek(lines, i, j, 1, d1s[0]) == 'M' && peek(lines, i, j, 1, d1s[1]) == 'S') ||
					(peek(lines, i, j, 1, d1s[0]) == 'S' && peek(lines, i, j, 1, d1s[1]) == 'M')
				if !d1 {
					continue
				}
				d2 := (peek(lines, i, j, 1, d2s[0]) == 'M' && peek(lines, i, j, 1, d2s[1]) == 'S') ||
					(peek(lines, i, j, 1, d2s[0]) == 'S' && peek(lines, i, j, 1, d2s[1]) == 'M')
				if d1 && d2 {
					count++
				}
			}
		}
	}

	return count
}

func peek(lines []string, i, j, num int, direction Direction) rune {
	newI := i + num*direction.x
	newJ := j + num*direction.y
	if newJ < 0 || newJ >= len(lines) {
		return rune(0)
	}
	if newI < 0 || newI >= len(lines[newJ]) {
		return rune(0)
	}

	return []rune(lines[newJ])[newI]
}

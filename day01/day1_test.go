package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	part1, err := Part1(strings.NewReader(input))

	if err != nil {
		t.Fatalf("want nil, got %v", err)
	}
	if part1 != 11 {
		t.Fatalf("want 11, got %d", part1)
	}
}

func TestPart2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	part2, err := Part2(strings.NewReader(input))

	if err != nil {
		t.Fatalf("want nil, got %v", err)
	}
	if part2 != 31 {
		t.Fatalf("want 31, got %d", part2)
	}
}

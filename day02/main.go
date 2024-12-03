package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	nums []int
}

func NewReport(nums []int) Report {
	return Report{nums: nums}
}

func (r Report) IsSafe() bool {
	if len(r.nums) < 2 {
		return true
	}

	isIncreasing := r.nums[1] > r.nums[0]

	for i := 1; i < len(r.nums); i++ {
		diff := r.nums[i] - r.nums[i-1]
		if diff == 0 {
			return false
		}
		if isIncreasing && !(diff >= 1 && diff <= 3) {
			return false
		}
		if !isIncreasing && !(diff <= -1 && diff >= -3) {
			return false
		}
	}

	return true
}

func (r Report) IsSafeWithDampener() bool {
	if r.IsSafe() {
		return true
	}

	for i := range r.nums {
		withoutLevel := append([]int{}, r.nums[:i]...)
		withoutLevel = append(withoutLevel, r.nums[i+1:]...)
		report := NewReport(withoutLevel)
		if report.IsSafe() {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v\n", err)
	}
	defer file.Close()

	count, countWithDampener, err := CountSafeReports(file)
	if err != nil {
		log.Fatalf("error counting safe reports: %v\n", err)
	}
	fmt.Printf("Part 1 - Safe Reports: %d\n", count)
	fmt.Printf("Part 2 - Safe Reports with Dampener: %d\n", countWithDampener)
}

func CountSafeReports(r io.Reader) (int, int, error) {
	count := 0
	countWithDampener := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		text := scanner.Text()

		nums := []int{}
		for _, s := range strings.Split(text, " ") {
			num, err := strconv.Atoi(s)
			if err != nil {
				return 0, 0, fmt.Errorf("error converting number: %w", err)
			}
			nums = append(nums, num)
		}
		report := NewReport(nums)
		if report.IsSafe() {
			count++
		}
		if report.IsSafeWithDampener() {
			countWithDampener++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, 0, fmt.Errorf("error scanning reports: %w", err)
	}

	return count, countWithDampener, nil
}

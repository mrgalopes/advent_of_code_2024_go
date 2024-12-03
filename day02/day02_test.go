package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestReportIsSafe(t *testing.T) {
	reports := []struct {
		report Report
		isSafe bool
		reason string
	}{
		{NewReport([]int{7, 6, 4, 2, 1}), true, "levels are all decreasing"},
		{NewReport([]int{1, 2, 7, 8, 9}), false, "2 7 is an increase of 5"},
		{NewReport([]int{9, 7, 6, 2, 1}), false, "6 2 is a decrease of 4"},
		{NewReport([]int{8, 6, 4, 4, 1}), false, "4 4 is neither an increase or a decrease"},
		{NewReport([]int{1, 3, 6, 7, 9}), true, "all levels are increasing by 1, 2 or 3"},
	}

	for i, r := range reports {
		t.Run(fmt.Sprintf("report %v, idx %d", r.report, i), func(t *testing.T) {
			actual := r.report.IsSafe()

			if actual != r.isSafe {
				t.Fatalf("got %v, wanted %v because %s", actual, r.isSafe, r.reason)
			}
		})
	}
}

func TestCountSafeReports(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	count, countWithDampener, err := CountSafeReports(strings.NewReader(input))
	if err != nil {
		t.Fatalf("wanted err nil, got %v", err)
	}
	if count != 2 {
		t.Fatalf("wanted 2 safe reports, got %d", count)
	}
	if countWithDampener != 4 {
		t.Fatalf("wanted 4 safe reports with dampener, got %d", count)
	}
}

func TestReportIsSafeWithDampener(t *testing.T) {
	reports := []struct {
		report Report
		isSafe bool
		reason string
	}{
		{NewReport([]int{7, 6, 4, 2, 1}), true, "levels are all decreasing"},
		{NewReport([]int{1, 2, 7, 8, 9}), false, "2 7 is an increase of 5"},
		{NewReport([]int{9, 7, 6, 2, 1}), false, "6 2 is a decrease of 4"},
		{NewReport([]int{1, 3, 2, 4, 5}), true, "removing second level, 3, makes it safe"},
		{NewReport([]int{8, 6, 4, 4, 1}), true, "removing third level, 4, makes it safe"},
		{NewReport([]int{1, 3, 6, 7, 9}), true, "all levels are increasing by 1, 2 or 3"},
	}

	for i, r := range reports {
		t.Run(fmt.Sprintf("report %v, idx %d", r.report, i), func(t *testing.T) {
			actual := r.report.IsSafeWithDampener()

			if actual != r.isSafe {
				t.Fatalf("got %v, wanted %v because %s", actual, r.isSafe, r.reason)
			}
		})
	}
}

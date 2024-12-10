package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPageOrderingRules(t *testing.T) {
	input := `47|53
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
53|13`

	want := map[string][]string{
		"53": {"47", "75", "61", "97"},
		"13": {"97", "61", "29", "47", "75", "53"},
		"61": {"97", "47", "75"},
		"47": {"97", "75"},
		"29": {"75", "97", "53", "61", "47"},
		"75": {"97"},
	}

	got, err := PageOrderingRules(input)

	if err != nil {
		t.Fatalf("want err nil, got %+v\n", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %+v, got %+v\n", want, got)
	}
}

func TestProcessRules(t *testing.T) {
	rules := map[string][]string{
		"53": {"47", "75", "61", "97"},
		"13": {"97", "61", "29", "47", "75", "53"},
		"61": {"97", "47", "75"},
		"47": {"97", "75"},
		"29": {"75", "97", "53", "61", "47"},
		"75": {"97"},
	}

	cases := []struct {
		input   string
		isValid bool
	}{
		{"75,47,61,53,29", true},
		{"97,61,53,29,13", true},
		{"75,29,13", true},
		{"75,97,47,61,53", false},
		{"61,13,29", false},
		{"97,13,75,29,47", false},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("case \"%s\"", c.input), func(t *testing.T) {
			got := IsManualValid(c.input, rules)

			if got != c.isValid {
				t.Fatalf("want %v, got %v", c.isValid, got)
			}
		})
	}
}

func TestFindMiddlePages(t *testing.T) {
	input := `47|53
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

	want := 143
	got, err := FindMiddlePagesSum(input)

	if err != nil {
		t.Fatalf("want err nil, got %+v\n", err)
	}

	if want != got {
		t.Fatalf("want %d, got %d\n", want, got)
	}
}

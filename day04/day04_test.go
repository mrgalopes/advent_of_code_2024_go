package main

import "testing"

func TestCountXMAS(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	want := 18

	got := CountXMAS(input)
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestCountXMAS2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	want := 9

	got := CountXMAS2(input)
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

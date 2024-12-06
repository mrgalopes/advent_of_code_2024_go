package main

import (
	"reflect"
	"testing"
)

func TestParseCorruptedMemory(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

	want := []Token{
		Mul{2, 4},
		Mul{5, 5},
		Mul{11, 8},
		Mul{8, 5},
	}

	got, err := ParseTokens(input)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestMultiplyAllUncorrupted(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	want := 161

	got, err := MultiplyAllUncorrupted(input)

	if err != nil {
		t.Fatalf("want err nil, got %v\n", err)
	}
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestParseCorruptedMemoryWithConditionals(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	want := []Token{
		Mul{2, 4},
		Dont{},
		Mul{5, 5},
		Mul{11, 8},
		Do{},
		Mul{8, 5},
	}

	got, err := ParseTokens(input)

	if err != nil {
		t.Fatalf("want err nil, got %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %+v, got %+v", want, got)
	}
}

func TestMultiplyAllUncorruptedWithConditionals(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	want := 48

	got, err := MultiplyAllUncorruptedWithConditionals(input)

	if err != nil {
		t.Fatalf("want err nil, got %v\n", err)
	}
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

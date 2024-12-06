package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("could not open file: %v\n", err)
	}

	// Part 1
	result, err := MultiplyAllUncorrupted(string(contents[:]))
	if err != nil {
		log.Fatalf("error multiplying uncorrupted: %v\n", err)
	}
	fmt.Printf("result 1: %d\n", result)

	// Part 2
	result, err = MultiplyAllUncorruptedWithConditionals(string(contents))
	if err != nil {
		log.Fatalf("error multiplying uncorrupted: %v\n", err)
	}
	fmt.Printf("result 2: %d\n", result)
}

func MultiplyAllUncorrupted(input string) (int, error) {
	tokens, err := ParseTokens(input)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, token := range tokens {
		if mul, ok := token.(Mul); ok {
			total += mul.A * mul.B
		}
	}

	return total, nil
}

func MultiplyAllUncorruptedWithConditionals(input string) (int, error) {
	tokens, err := ParseTokens(input)
	if err != nil {
		return 0, err
	}

	total := 0
	enabled := true
	for _, token := range tokens {
		switch t := token.(type) {
		case Mul:
			if enabled {
				total += t.A * t.B
			}
		case Do:
			enabled = true
		case Dont:
			enabled = false
		}
	}

	return total, nil
}

type Token interface{}

type Mul struct {
	A int
	B int
}

type Do struct{}
type Dont struct{}

func ParseTokens(input string) ([]Token, error) {
	var tokens []Token
	validMult := regexp.MustCompile(`(?<donot>don't\(\))|(do\(\))|mul\((?<a>\d+),(?<b>\d+)\)`)

	result := validMult.FindAllStringSubmatch(input, -1)

	for _, match := range result {
		if match[0][:3] == "mul" {
			a, err := strconv.Atoi(match[3])
			if err != nil {
				return nil, fmt.Errorf("first value invalid: %w", err)
			}

			b, err := strconv.Atoi(match[4])
			if err != nil {
				return nil, fmt.Errorf("second value invalid: %w", err)
			}

			token := Mul{a, b}
			tokens = append(tokens, token)
		} else if match[0] == "don't()" {
			tokens = append(tokens, Dont{})
		} else if match[0] == "do()" {
			tokens = append(tokens, Do{})
		}
	}

	return tokens, nil
}

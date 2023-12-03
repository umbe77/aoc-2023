package day03

import "testing"

var input = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestPart1(t *testing.T) {
	parts := parse(input)
	r := part1(parts)
	if r != 4361 {
		t.Errorf("Expected 4361, Got %d", r)
	}
}

func TestPart2(t *testing.T) {
	parts := parse(input)
	r := part2(parts)
	if r != 467835 {
		t.Errorf("Expected 467835, Got %d", r)
	}
}

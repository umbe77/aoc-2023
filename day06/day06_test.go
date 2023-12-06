package day06

import "testing"

var input = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestPart1(t *testing.T) {
	races := parse(input)
	r := part1(races)

	if r != 288 {
		t.Errorf("Expected 288, Got %d", r)
	}
}

func TestPart2(t *testing.T) {
	race := parse2(input)
	r := part2(race)

	if r != 71503 {
		t.Errorf("Expected 71503, Got %d", r)
	}
}

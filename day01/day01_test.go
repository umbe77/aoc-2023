package day01

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	result := part1(input)
	if result != 142 {
		t.Errorf("Expcetd 142 got %d", result)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}

	result := part1(input)
	if result != 281 {
		t.Errorf("Expcetd 281 got %d", result)
	}
}

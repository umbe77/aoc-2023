package day02

import "testing"

var input = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func TestPart1(t *testing.T) {

	games := make([]Game, 0)
	for _, l := range input {
		games = append(games, parseLine(l))
	}

	c := part1(games)
	if c != 8 {
		t.Errorf("Expcted 8, Got %d", c)
	}
}

func TestPart2(t *testing.T) {
	games := make([]Game, 0)
	for _, l := range input {
		games = append(games, parseLine(l))
	}

    c := part2(games)
    if c != 2286 {
        t.Errorf("Expected 2286, Got %d", c)
    }
}

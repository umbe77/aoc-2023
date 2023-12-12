package day08

import "testing"

var input1 = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

var input2 = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

var input3 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestPart1(t *testing.T) {
	nm1 := parse(input1)
	m1 := part1(nm1)
	if m1 != 2 {
		t.Errorf("Expected 2, Got %d", m1)
	}

	nm2 := parse(input2)
	m2 := part1(nm2)
	if m2 != 6 {
		t.Errorf("Expected 6, Got %d", m2)
	}
}

func TestPart2(t *testing.T) {
    nm := parse(input3)
    m := part2(nm)
    if m != 6 {
        t.Errorf("Expected 6, Got %d", m)
    }
}

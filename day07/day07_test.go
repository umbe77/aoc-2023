package day07

import "testing"

var input = [...]string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestPart1(t *testing.T) {
    hands:=make([]Hand, 0)
    for _, l := range input{
        hands = append(hands, parse(l))
    }

    r := part1(hands)
    if r != 6440 {
        t.Errorf("Expected 6440, Got %d", r)
    }
}

// func TestPart2(t *testing.T) {
//     hands:=make([]Hand, 0)
//     for _, l := range input{
//         hands = append(hands, parse2(l))
//     }
//
//     r := part1(hands)
//     if r != 5905 {
//         t.Errorf("Expected 5905, Got %d", r)
//     }
// }

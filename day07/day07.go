package day07

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	hands := make([]Hand, 0)
    hands2 := make([]Hand, 0)
	utils.ReadFile("day07/input_test.txt", func(line string) {
		hands = append(hands, parse(line))
		hands2 = append(hands2, parse2(line))
	})

	// fmt.Printf("Part 1: %d\n", part1(hands))
	fmt.Printf("Part 2: %d\n", part1(hands2))
}

var cardValues = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var cardValues2 = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 11,
	'9': 10,
	'8': 9,
	'7': 8,
	'6': 7,
	'5': 6,
	'4': 5,
	'3': 4,
	'2': 3,
	'J': 2,
}

type Draw struct {
	cards []int
	point int
}

type Hand struct {
	draw Draw
    orig string
	bid  int
}

func getPoint(p []int) int {
	switch p[0] {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		if p[1] == 2 {
			return 5
		}
		return 4
	case 2:
		if p[1] == 2 {
			return 3
		}
		return 2
	default:
		return 1
	}
}

func getPoint2(p []int, jCount int) int {
	switch p[0] {
	case 5:
		return 7
	case 4:
		return 6
	case 3:
		if p[1] == 2 {
			return 5
		}
		return 4
	case 2:
		if p[1] == 2 {
			return 3
		}
		return 2
	default:
		return 1
	}
}

func parse(input string) Hand {
	in := strings.Fields(input)
	d := in[0]
	bid := utils.Atoi(in[1])

	draw := Draw{
		cards: make([]int, 5),
	}

	counter := make(map[int]int, 5)

	for i, c := range d {
		card := cardValues[c]
		draw.cards[i] = card
		if _, ok := counter[card]; !ok {
			counter[card] = 0
		}
		counter[card] = counter[card] + 1
	}

	p := make([]int, 5)
	i := 0
	for _, v := range counter {
		p[i] = v
		i = i + 1
	}

	slices.Sort(p)
	slices.Reverse(p)

	point := getPoint(p)

	draw.point = point

	h := Hand{
		draw: draw,
        orig: input,
		bid:  bid,
	}

	return h
}

func parse2(input string) Hand {
	in := strings.Fields(input)
	d := in[0]
	bid := utils.Atoi(in[1])

	draw := Draw{
		cards: make([]int, 5),
	}

	counter := make(map[int]int, 5)

	for i, c := range d {
		card := cardValues2[c]
		draw.cards[i] = card
		if _, ok := counter[card]; !ok {
			counter[card] = 0
		}
		counter[card] = counter[card] + 1
	}
	jCount := 0
	if n, ok := counter[2]; ok {
		jCount = n
	}

	p := make([]int, 5)
	i := 0
	for _, v := range counter {
		p[i] = v
		i = i + 1
	}

	slices.Sort(p)
	slices.Reverse(p)

	point := getPoint2(p, jCount)

	draw.point = point

	h := Hand{
		draw: draw,
        orig: input,
		bid:  bid,
	}

	return h
}

func part1(hands []Hand) int {
	slices.SortFunc(hands, func(a, b Hand) int {
		if n := cmp.Compare(a.draw.point, b.draw.point); n != 0 {
			return n
		}

		for i, c := range a.draw.cards {
			if n := cmp.Compare(c, b.draw.cards[i]); n != 0 {
				return n
			}
		}
		return 0
	})

	tot := 0
	for i, v := range hands {
        fmt.Printf("%v\n", v)
		tot = tot + (v.bid * (i + 1))
	}
	return tot
}

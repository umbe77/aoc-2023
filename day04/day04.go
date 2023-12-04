package day04

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	cards := make([]*Card, 0)
	utils.ReadFile("day04/input.txt", func(line string) {
		cards = append(cards, parseLine(line))
	})

	fmt.Printf("Part 1: %d\n", part1(cards))
	fmt.Printf("Part 2: %d\n", part2(cards))
}

type Card struct {
	id   int
	wins int
	wn   []int
	hn   []int
}

func newCard() *Card {
	return &Card{
		wins: 1,
		wn:   make([]int, 0),
		hn:   make([]int, 0),
	}
}

func parseLine(row string) *Card {
	card := newCard()
	r := bufio.NewReader(strings.NewReader(row))

	lastNumRead := 0
	isReadingW := true
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}

		if c == ':' {
			card.id = lastNumRead
		}

		if c == '|' {
			isReadingW = false
		}

		if utils.IsDigit(c) {
			lastNumRead = utils.ReadInt(r)
			if card.id != 0 {
				if isReadingW {
					card.wn = append(card.wn, lastNumRead)
				} else {
					card.hn = append(card.hn, lastNumRead)
				}
			}
		}

	}

	slices.Sort(card.wn)

	return card
}

func pointPerCard(card *Card) int {
	points := 0

	for _, h := range card.hn {
		if _, ok := slices.BinarySearch(card.wn, h); ok {
			points = points + 1
		}
	}

	if points == 0 {
		return points
	}
	p := utils.PowInt(2, points-1)
	return p
}

func part1(cards []*Card) int {
	totalPoints := 0

	for _, c := range cards {
		p := pointPerCard(c)
		totalPoints = totalPoints + p
	}

	return totalPoints
}

func cardsWon(card *Card) int {
	total := 0

	for _, h := range card.hn {
		if _, ok := slices.BinarySearch(card.wn, h); ok {
			total = total + 1
		}
	}
	return total
}

func part2(cards []*Card) int {

	for i, c := range cards {
		t := cardsWon(c)
		for iw := 0; iw < c.wins; iw++ {
			for j := 1; j <= t && i+j < len(cards); j++ {
				cards[i+j].wins = cards[i+j].wins + 1
			}
		}
	}

	totalCards := 0
	for _, c := range cards {
		totalCards = totalCards + c.wins
	}
	return totalCards
}

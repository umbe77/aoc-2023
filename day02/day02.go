package day02

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	games := make([]Game, 0)
	utils.ReadFile("day02/input.txt", func(line string) {
		games = append(games, parseLine(line))
	})

	fmt.Printf("Part 1: %d\n", part1(games))
	fmt.Printf("Part 2: %d\n", part2(games))
}

type GameSet struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	sets []*GameSet
}

func (g Game) print() {
	fmt.Printf("Game: %d: ", g.id)
	for _, s := range g.sets {
		fmt.Printf("%d red, %d green, %d blue; ", s.red, s.green, s.blue)
	}
	fmt.Printf("\n")
}

func parseLine(l string) Game {

	var (
		cNum int
	)

	g := Game{}
	g.sets = make([]*GameSet, 0)

	currentSet := &GameSet{}
	g.sets = append(g.sets, currentSet)

	r := bufio.NewReader(strings.NewReader(l))
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}

		if c == ':' {
			g.id = cNum
			continue
		}

		if c == ';' {
			currentSet = &GameSet{}
			g.sets = append(g.sets, currentSet)
		}

		if utils.IsSpace(c) || c == ',' {
			continue
		}

		if utils.IsDigit(c) {
			cNum = utils.ReadInt(r)
		}

		if utils.IsAplpha(c) {
			s := utils.ReadString(r)
			if s == "Game" {
				continue
			}
			switch s {
			case "red":
				currentSet.red = currentSet.red + cNum
				break
			case "green":
				currentSet.green = currentSet.green + cNum
				break
			case "blue":
				currentSet.blue = currentSet.blue + cNum
				break
			}
		}
	}

	return g
}

const (
	valid_red   = 12
	valid_green = 13
	valid_blue  = 14
)

func isValidGame(g Game) bool {
	for _, s := range g.sets {
		if s.red > valid_red || s.green > valid_green || s.blue > valid_blue {
			return false
		}
	}
	return true
}

func part1(games []Game) int {

	var (
		total int
	)

	for _, g := range games {
        if isValidGame(g) {
            total = total + g.id
        }
	}

	return total
}

func getGamePower(g Game) int {
    maxRed := 0
    maxGreen := 0
    maxBlue := 0

    for _, s := range g.sets {
        if s.red > maxRed {
            maxRed = s.red
        }
        if s.green > maxGreen {
            maxGreen = s.green
        }
        if s.blue > maxBlue {
            maxBlue = s.blue
        }
    }

    return maxRed * maxGreen * maxBlue
}

func part2(games []Game) int {
    var (
        totalPower int
    )

    for _, g := range games {
        p := getGamePower(g)
        totalPower = totalPower + p
    }

    return totalPower
}

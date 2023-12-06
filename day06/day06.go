package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day06/input.txt", func(line string) {
		input = append(input, line)
	})

	races := parse(input)

	fmt.Printf("Part 1: %d\n", part1(races))

	race := parse2(input)
	fmt.Printf("Part 2: %d\n", part2(race))
}

type Race struct {
	time     int
	distance int
}

func parse(input []string) []Race {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	distances := strings.Fields(strings.Split(input[1], ":")[1])
	races := make([]Race, len(times))
	for i, t := range times {
		time, _ := strconv.Atoi(t)
		distance, _ := strconv.Atoi(distances[i])
		races[i] = Race{
			time:     time,
			distance: distance,
		}
	}
	return races
}

func parse2(input []string) Race {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	distances := strings.Fields(strings.Split(input[1], ":")[1])
	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(distances, ""))
	return Race{
		time:     t,
		distance: d,
	}
}

func getPossibleTimes(race Race) int {
	wins := 0

	for i := 1; i <= race.time; i++ {
		if (race.time-i)*i > race.distance {
			wins = wins + 1
		}
	}

	return wins
}

func part1(races []Race) int {
	prod := 1
	for _, race := range races {
		prod = prod * getPossibleTimes(race)
	}
	return prod
}

func part2(race Race) int {
    tot := 0
    tot = getPossibleTimes(race)
    return tot
}

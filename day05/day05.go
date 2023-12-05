package day05

import (
	"fmt"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
    rows := make([]string, 0)
    utils.ReadFile("day05/input.txt", func(line string) {
        rows = append(rows, line)
    })

    fmt.Printf("Part 1: %d\n", part1())

}

func parse(rows []string) {

}

func part1() int {
    return 0
}

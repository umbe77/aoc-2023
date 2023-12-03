package day03

import (
	"bytes"
	"fmt"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	input := make([]string, 0)
	utils.ReadFile("day03/input.txt", func(line string) {
		input = append(input, line)
	})

	p := parse(input)

	fmt.Printf("Part 1: %d\n", part1(p))
	fmt.Printf("Part 2: %d\n", part2(p))
}

type partType int

const (
	none partType = iota
	symbol
	number
)

type Part struct {
	n  int
	s  string
	id int
	t  partType
}

var (
	id int = 0
)

func readNum(idx int, h string) (int, int) {
	var (
		buf bytes.Buffer
	)
	newIdx := idx
	for i := idx; i < len(h); i++ {
		c := rune(h[i])
		if utils.IsDigit(c) {
			buf.WriteRune(c)
		} else {
			newIdx = (i - 1)
			break
		}
	}
	return utils.Atoi(buf.String()), newIdx
}

func parse(rows []string) [][]Part {
	parts := make([][]Part, len(rows))
	for y, v := range rows {
		parts[y] = make([]Part, len(v))
		for x := 0; x < len(v); x++ {
			h := rune(v[x])
			c := string(h)
			if c == "." {
				parts[y][x] = Part{t: none}
				continue
			}
			if utils.IsDigit(h) {
				n, i := readNum(x, v)
				id = id + 1
				for ; x <= i; x++ {
					parts[y][x] = Part{
						n:  n,
						t:  number,
						id: id,
					}
				}
				x = x - 1
				continue
			}
			parts[y][x] = Part{t: symbol, s: c}
		}
	}
	return parts
}

func isInList(l []int, n int) bool {
	for _, i := range l {
		if i == n {
			return true
		}
	}
	return false
}

func part1(parts [][]Part) int {

	nums := make(map[int]int, 0)
	maxY := len(parts)
	maxX := len(parts[0])
	for y, v := range parts {
		for x, h := range v {
			if h.t == symbol {
				for i := -1; i <= 1; i++ {
					if y+i >= 0 && y+i < maxY {
						for j := -1; j <= 1; j++ {
							if x+j >= 0 && x+j < maxX {
								p := parts[y+i][x+j]
								if _, ok := nums[p.id]; !ok && p.t == number {
									nums[p.id] = p.n
								}
							}
						}
					}
				}
			}
		}
	}

	total := 0
	for _, n := range nums {
		total = total + n
	}
	return total
}

func part2(parts [][]Part) int {

	total := 0
	maxY := len(parts)
	maxX := len(parts[0])
	for y, v := range parts {
		for x, h := range v {
			if h.t == symbol && h.s == "*" {
				adjacents := make(map[int]int, 0)
				for i := -1; i <= 1; i++ {
					if y+i >= 0 && y+i < maxY {
						for j := -1; j <= 1; j++ {
							if x+j >= 0 && x+j < maxX {
								p := parts[y+i][x+j]
								if _, ok := adjacents[p.id]; !ok && p.t == number {
									adjacents[p.id] = p.n
								}
							}
						}
					}
				}
				if len(adjacents) == 2 {
					prod := 1
					for _, v := range adjacents {
						prod = prod * v
					}
					total = total + prod
				}
			}
		}
	}

	return total
}

package day05

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	rows := make([]string, 0)
	utils.ReadFile("day05/input_test.txt", func(line string) {
		rows = append(rows, line)
	})

	a := parse(rows)

	fmt.Printf("Part 1: %d\n", part1(a))
	// fmt.Printf("Part 2: %d\n", part2(a))

}

type Matrix struct {
	name string
	m    [][]int
}

func newMatrix(n string) *Matrix {
	return &Matrix{
		name: n,
		m:    make([][]int, 0),
	}
}

type Almanc struct {
	seeds    []int
	matrices []*Matrix
}

func newAlmanac() Almanc {
	return Almanc{
		seeds:    make([]int, 0),
		matrices: make([]*Matrix, 0),
	}
}

func parse(rows []string) Almanc {
	a := newAlmanac()
	var current *Matrix
	for _, r := range rows {
		r := bufio.NewReader(strings.NewReader(r))
		var name string
		c, _, _ := r.ReadRune()
		if utils.IsDigit(c) {
			current.m = append(current.m, make([]int, 0))
		}
		r.UnreadRune()
		for {
			c, _, err := r.ReadRune()
			if err != nil {
				break
			}
			if utils.IsAplpha(c) {
				name = utils.ReadString(r)
				if name != "seeds" && name != "map" {
					current = newMatrix(name)
					a.matrices = append(a.matrices, current)
				}
			}
			if utils.IsDigit(c) {
				n := utils.ReadInt(r)
				if name == "seeds" {
					a.seeds = append(a.seeds, n)
				} else {
					current.m[len(current.m)-1] = append(current.m[len(current.m)-1], n)
				}
			}
		}
	}
	return a
}

func getLocation(seed int, a Almanc) int {
	current := seed
	// fmt.Printf("seed: %d - ", seed)
	for _, m := range a.matrices {
		for _, y := range m.m {
			dstRangeStart := y[0]
			srcRangeStart := y[1]
			rangeLen := y[2]
			// fmt.Println(current, srcRangeStart, rangeLen, srcRangeStart+rangeLen, current >= srcRangeStart, current < srcRangeStart+rangeLen)
			if current >= srcRangeStart && current < srcRangeStart+rangeLen {
				current = current - srcRangeStart + dstRangeStart
				break
			}
		}
		// fmt.Printf("%s: %d - ", strings.Split(m.name, "-")[2], current)
	}
	// fmt.Println()
	return current
}

func part1(a Almanc) int {
	minLoc := math.MaxInt
	for _, s := range a.seeds {
		c := getLocation(s, a)
		if c < minLoc {
			minLoc = c
		}
	}
	return minLoc
}

type Range struct {
	start, end int
}

func getRanges(rs []Range, m *Matrix) []Range {
	output := make([]Range, 0)
	for _, y := range m.m {
		for _, r := range rs {
			dstRange := y[0]
			start2 := y[1]
			end2 := start2 + y[2] - 1
			start := r.start
			if start2 > r.start {
				start = start2
			}

			end := r.end
			if end2 < r.end {
				end = end2
			}

			if end < r.start || start > end {
				continue
			}

			output = append(output, Range{
				start: start - start2 + dstRange,
				end:   (start - start2 + dstRange) + (end - start),
			})

			if r.start < start {
				output = append(output, Range{start: r.start, end: start - 1})
			}

			if r.end > end {
				output = append(output, Range{start: end + 1, end: r.end})
			}

		}
	}
	return output
}

func getLocation2(seed, seedEnd int, a Almanc) int {

	// r := Range{start: seed, end: seed}
    return 0

}

func part2(a Almanc) int {
	minLoc := math.MaxInt
	for i := 0; i < len(a.seeds); i = i + 2 {
		s := a.seeds[i]
		l := a.seeds[i+1]
		c := getLocation2(s, s+l-1, a)
		if c < minLoc {
			minLoc = c
		}
	}
	return minLoc
}

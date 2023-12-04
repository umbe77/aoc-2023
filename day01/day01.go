package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	rows := make([]string, 0)
	utils.ReadFile("day01/input.txt", func(line string) {
		rows = append(rows, line)
	})

	fmt.Printf("Part 1: %d\n", part1(rows))
	fmt.Printf("Part 2: %d\n", part2(rows))
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func getNum(row string) int {
	var (
		first, last int64
	)
	first = -1
	last = 0
	r := bufio.NewReader(strings.NewReader(row))

	for {
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}

		if isDigit(c) {
			last, _ = strconv.ParseInt(string(c), 10, 32)
			if first == -1 {
				first = last
			}
		}

	}

	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	return res
}

func part1(rows []string) int {
	total := 0

	for _, row := range rows {
		l := getNum(row)
		total = total + l
	}

	return total
}

func getNumFromLit(lit string) (int64, error) {
	switch lit {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	}
	return -1, fmt.Errorf("not a number")
}

func checkIsNumber(r *strings.Reader) (int64, error) {
	var (
		n   int64 = -1
		i   int
		buf bytes.Buffer
	)
	for i = 0; i < 5; i++ {
		c, _, err := r.ReadRune()
		if err != nil {
			if i == 0 {
				return -2, err
			}
			r.Seek(int64(i-1)*-1, io.SeekCurrent)
			return -3, nil
		}
		if i == 0 && isDigit(c) {
			n, _ = strconv.ParseInt(string(c), 10, 32)
			return n, nil
		}
		buf.WriteRune(c)
		lit := buf.String()
		n, err := getNumFromLit(lit)
		if err == nil {
			r.Seek(int64(i-1)*-1, io.SeekCurrent)
			return n, nil
		}
	}

	r.Seek(-4, io.SeekCurrent)

	return -1, nil
}

func getNum2(row string) int {
	var (
		first, last int64
	)
	first = -1
	last = 0
	r := strings.NewReader(row)

	for {

		num, err := checkIsNumber(r)
		if err != nil {
			break
		}
		if num >= 0 {
			last = num
			if first == -1 {
				first = last
			}
		}
	}
	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	return res
}

func part2(rows []string) int {
	total := 0
	for _, row := range rows {
		l := getNum2(row)
		total = total + l
	}
	return total
}

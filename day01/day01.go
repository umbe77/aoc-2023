package day01

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2023/utils"
)

func Execute() {
	rows := make([]string, 0)
	utils.ReadFile("day01/input_test.txt", func(line string) {
		rows = append(rows, line)
	})

	// fmt.Printf("Part 1: %d\n", part1(rows))
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

func checkIsNumber(input string) (int64, error) {

	switch input {
	case "one":
		return 1, nil
    case "two"
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
	return 0, fmt.Errorf("not a number")
}

func getNum2(row string) int {
	var (
		first, last int64
		buf         bytes.Buffer
	)
	first = -1
	last = 0
	r := bufio.NewReader(strings.NewReader(row))

	for {
		var num int64 = -1
		c, _, err := r.ReadRune()
		if err != nil {
			break
		}
		buf.WriteRune(c)
        fmt.Println(buf.String())
		cNum, isNum := checkIsNumber(buf.String())
		if isNum == nil {
			num = cNum
		} else if isDigit(c) {
			num, _ = strconv.ParseInt(string(c), 10, 32)
		}
        fmt.Println(num)
		if num != -1 {
			last = num
			if first == -1 {
				first = last
			}
			buf.Reset()
		}
	}
	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
	return res
}

func part2(rows []string) int {
	total := 0
	for _, row := range rows {
		l := getNum2(row)
        fmt.Println(l)
		total = total + l
	}
	return total
}

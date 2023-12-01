package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2023/day01"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "01", "day in format dd")
	flag.Parse()

	fmt.Printf("Day %s\n", day)
	switch day {
	case "01":
		day01.Execute()
	}
}

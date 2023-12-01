package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func ReadFile(fPath string, line func(line string)) {
	f, err := os.Open(fPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadAllFile(fPath string) string {
	buf, err := os.ReadFile(fPath)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func Atoi(a string) int {
	r, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return r
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

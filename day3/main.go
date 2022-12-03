package main

import (
	"bufio"
	"fmt"
	"os"
)

func runeToNumber(r rune) rune {
	switch {
	case 'a' <= r && r <= 'z':
		return r - 'a' + 1
	case 'A' <= r && r <= 'Z':
		return r - 'A' + 27
	default:
		panic(r)
	}
}

func task1(line string) int {
	firstPartChars := make(map[rune]bool)
	for _, c := range line[:len(line)/2] {
		firstPartChars[c] = true
	}
	for _, c := range line[len(line)/2:] {
		if ok := firstPartChars[c]; ok {
			return int(runeToNumber(c))
		}
	}
	panic("common rune was not found")
}

func task2(lines []string) int {
	commonChars := make(map[rune]bool)
	for i, l := range lines {
		newCommonChars := make(map[rune]bool)
		for _, c := range l {
			if i == 0 {
				commonChars[c] = true
			} else {
				if commonChars[c] {
					newCommonChars[c] = true
				}
			}
		}
		if i == 0 {
			continue
		}
		if len(newCommonChars) == 1 {
			for k, _ := range newCommonChars {
				return int(runeToNumber(k))
			}
		}
		commonChars = newCommonChars
	}
	panic("common rune was not found")
}

func main() {
	readFile, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var task1sum int
	var lines []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
		task1sum += task1(line)
	}
	fmt.Println("Day 1 task 1: ", task1sum)

	var task2sum int
	for i := 0; i < len(lines); i += 3 {
		task2sum += task2(lines[i : i+3])
	}
	fmt.Println("Day 1 task 2: ", task2sum)
}

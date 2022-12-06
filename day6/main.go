package main

import (
	"bufio"
	"fmt"
	"os"
)

func task(line string, markerSize int) int {
	var window []rune
	found := make(map[rune]int)
	for i, r := range line {
		found[r]++
		window = append(window, r)
		if len(window) > markerSize {
			popped := window[0]
			found[popped]--
			if found[popped] == 0 {
				delete(found, popped)
			}
			window = window[1:]
		}
		if len(found) == markerSize {
			return i + 1
		}
	}
	panic("marker was not found")
}

func main() {
	readFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println("task 1: ", task(line, 4))
		fmt.Println("task 2: ", task(line, 14))
	}
}

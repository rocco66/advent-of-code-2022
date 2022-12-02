package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var signScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var loseMap = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}
var drawMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func task1(line string) int {
	var winMap = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	var roundScore int
	tokens := strings.Split(line, " ")
	opponentSign := tokens[0]
	mySign := tokens[1]
	roundScore += signScore[mySign]
	if drawMap[opponentSign] == mySign {
		roundScore += 3
	} else {
		if opponentSign == winMap[mySign] {
			roundScore += 6
		}
	}
	return roundScore
}

func task2(line string) int {
	var winMap = map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	var roundScore int
	tokens := strings.Split(line, " ")
	opponentSign := tokens[0]
	resultSign := tokens[1]
	var mySign string
	switch resultSign {
	case "X":
		mySign = loseMap[opponentSign]
	case "Y":
		mySign = drawMap[opponentSign]
		roundScore += 3
	case "Z":
		mySign = winMap[opponentSign]
		roundScore += 6
	}
	roundScore += signScore[mySign]
	return roundScore
}

func main() {
	readFile, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var task1sum, task2sum int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		task1sum += task1(line)
		task2sum += task2(line)

	}
	fmt.Println("Day 1 task 1: ", task1sum)
	fmt.Println("Day 1 task 2: ", task2sum)
}

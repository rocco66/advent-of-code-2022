package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Job struct {
	First int
	Last  int
}

func NewJob(repr string) Job {
	tokens := strings.Split(repr, "-")
	first, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}
	return Job{first, second}

}

func task1(line string) bool {
	jobsRepr := strings.Split(line, ",")
	first := NewJob(jobsRepr[0])
	second := NewJob(jobsRepr[1])
	if (first.First <= second.First && first.Last >= second.Last) ||
		(second.First <= first.First && second.Last >= first.Last) {
		return true
	}
	return false
}

func task2(line string) bool {
	jobsRepr := strings.Split(line, ",")
	first := NewJob(jobsRepr[0])
	second := NewJob(jobsRepr[1])
	if (first.First >= second.First && first.First <= second.Last) ||
		(first.Last >= second.First && first.Last <= second.Last) {
		return true
	}
	return false
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
		task1Result := task1(line)
		if task1Result {
			task1sum++
		}
		if task1Result || task2(line) {
			task2sum++
		}

	}
	fmt.Println("Day 1 task 1: ", task1sum)
	fmt.Println("Day 1 task 2: ", task2sum)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Crate struct {
	Stack []rune
}

func (s *Crate) Pop(count int) []rune {
	result := s.Stack[len(s.Stack)-count:]
	s.Stack = s.Stack[:len(s.Stack)-count]
	return result
}

func (s *Crate) Push(items []rune) {
	for i := len(items) - 1; i >= 0; i-- {
		s.Stack = append(s.Stack, items[i])
	}
}

func (s *Crate) PushCrateMover9001(items []rune) {
	for i := 0; i < len(items); i++ {
		s.Stack = append(s.Stack, items[i])
	}
}

type State struct {
	Crates []Crate
}

func (s *State) Move(move Move, crateMover9001 bool) {
	items := s.Crates[move.from-1].Pop(move.count)
	if crateMover9001 {
		s.Crates[move.to-1].PushCrateMover9001(items)

	} else {
		s.Crates[move.to-1].Push(items)
	}
}

func (s *State) Push(to int, item rune) {
	s.Crates[to].Push([]rune{item})
}

func getSecondRune(str string) rune {
	return []rune(str)[1]
}

func NewState(lines []string) State {
	lastLine := lines[len(lines)-1]
	trimmedLastLine := strings.Trim(lastLine, " ")
	lastLineTokens := strings.Split(trimmedLastLine, "   ")
	cratesNumber, err := strconv.Atoi(lastLineTokens[len(lastLineTokens)-1])
	if err != nil {
		panic(err)
	}

	var result State
	result.Crates = make([]Crate, cratesNumber)
	for i := len(lines) - 2; i >= 0; i-- {
		for crateIndex := 0; crateIndex*4+3 <= len(lines[i]); crateIndex++ {
			itemRepr := lines[i][crateIndex*4 : crateIndex*4+3]
			if itemRepr == "   " {
				continue
			}
			result.Crates[crateIndex].Push([]rune{getSecondRune(itemRepr)})
		}
	}
	return result

}

type Move struct {
	from, to, count int
}

func NewMove(line string) Move {
	tokens := strings.Split(line, " ")
	count, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic("count error")
	}
	from, err := strconv.Atoi(tokens[3])
	if err != nil {
		panic("from error")
	}
	to, err := strconv.Atoi(tokens[5])
	if err != nil {
		panic("to error")
	}
	return Move{from, to, count}
}

func task(state State, moves []Move, crateMover9001 bool) string {
	for _, move := range moves {
		state.Move(move, crateMover9001)
	}
	var result string
	for _, crate := range state.Crates {
		result += string(crate.Stack[len(crate.Stack)-1])
	}
	return result
}

func main() {
	readFile, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var stateLines []string
	var moves []Move
	var movePart bool
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			movePart = true
			continue
		}
		if movePart {
			moves = append(moves, NewMove(line))
		} else {
			stateLines = append(stateLines, line)
		}
	}
	fmt.Println("task 1: ", task(NewState(stateLines), moves, false))
	fmt.Println("task 2: ", task(NewState(stateLines), moves, true))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const DiskSize = 70000000
const SpaceNeed = 30000000

type FSSpider struct {
	CurrentPath []string
	DirSizes    map[string]int
}

func NewFS() FSSpider {
	return FSSpider{nil, make(map[string]int)}
}

func (s *FSSpider) Add(fileSize int) {
	s.DirSizes["/"] += fileSize
	for i := 1; i <= len(s.CurrentPath); i++ {
		path := "/" + strings.Join(s.CurrentPath[:i], "/")
		s.DirSizes[path] += fileSize
	}
}

func (s *FSSpider) Push(line string) {
	switch {
	case len(line) >= 6 && line[:6] == "$ cd /":
		{
			s.CurrentPath = nil
		}
	case len(line) >= 4 && (line[:4] == "$ ls" || line[:4] == "dir "):
		{
			break
		}
	case len(line) >= 7 && line[:7] == "$ cd ..":
		{
			s.CurrentPath = s.CurrentPath[:len(s.CurrentPath)-1]
		}
	case len(line) >= 5 && line[:5] == "$ cd ":
		{
			s.CurrentPath = append(s.CurrentPath, line[5:])
		}
	default:
		{
			tokens := strings.Split(line, " ")
			fileSize, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			s.Add(fileSize)
		}
	}
}

func task2(fs FSSpider) int {
	free := DiskSize - fs.DirSizes["/"]
	needs := SpaceNeed - free
	var sizes []int
	for _, size := range fs.DirSizes {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	for _, s := range sizes {
		if s >= needs {
			return s
		}
	}
	panic("can't find dir for removeing")
}

func main() {
	readFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fs := NewFS()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fs.Push(line)
	}
	var task1res int
	for _, size := range fs.DirSizes {
		if size <= 100000 {
			task1res += size
		}
	}
	fmt.Println("task 1: ", task1res)
	fmt.Println("task 2: ", task2(fs))
}

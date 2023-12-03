package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	p1, err := part1(lines)
	if err != nil {
		panic("part 1: " + err.Error())
	}
	fmt.Printf("part 1: %d\n", p1)

	p2, err := part2(lines)
	if err != nil {
		panic("part 2: " + err.Error())
	}
	fmt.Printf("part 2: %d\n", p2)
}

func part1(lines []string) (total int, err error) {
	n := parse(lines)
	for _, nn := range n {
		total += nn
	}
	return total, nil
}

func part2(lines []string) (total int, err error) {
	return 0, nil
}

// incredibly janky character parser
// trying not to go full blown rsc with it
func parse(lines []string) (partNumbers []int) {
	var characters [][]string
	for _, l := range lines {
		s := strings.Split(l, "")
		characters = append(characters, s)
	}

	var (
		height               = len(characters) - 1
		width                = len(characters[0]) - 1
		currentNumber        string
		inNumber             bool
		currentTouchesSymbol bool
	)

	reset := func() {
		currentNumber = ""
		inNumber = false
		currentTouchesSymbol = false
	}

	touchesSymbol := func(y, x int) bool {
		if y == 0 && x == 0 {
			return isSymbol(characters[0][1]) || isSymbol(characters[1][0]) || isSymbol(characters[1][1])
		}
		if y == 0 && x == width {
			return isSymbol(characters[0][width-1]) || isSymbol(characters[1][width-1]) || isSymbol(characters[1][width])
		}
		if y == height && x == 0 {
			return isSymbol(characters[height-1][0]) || isSymbol(characters[height-1][1]) || isSymbol(characters[height][1])
		}
		if y == height && x == width {
			return isSymbol(characters[height][width-1]) || isSymbol(characters[height-1][width-1]) || isSymbol(characters[height-1][width])
		}
		if y == 0 {
			return isSymbol(characters[0][x-1]) || isSymbol(characters[0][x+1]) || isSymbol(characters[1][x-1]) || isSymbol(characters[1][x]) || isSymbol(characters[1][x+1])
		}
		if y == height {
			return isSymbol(characters[height][x-1]) || isSymbol(characters[height][x+1]) || isSymbol(characters[height-1][x-1]) || isSymbol(characters[height-1][x]) || isSymbol(characters[height-1][x+1])
		}
		if x == 0 {
			return isSymbol(characters[y-1][0]) || isSymbol(characters[y-1][1]) || isSymbol(characters[y][1]) || isSymbol(characters[y+1][0]) || isSymbol(characters[y+1][1])
		}
		if x == width {
			return isSymbol(characters[y-1][0]) || isSymbol(characters[y-1][width-1]) || isSymbol(characters[y][width-1]) || isSymbol(characters[y+1][width]) || isSymbol(characters[y+1][width-1])
		}
		return isSymbol(characters[y-1][x-1]) || isSymbol(characters[y-1][x]) || isSymbol(characters[y-1][x+1]) || isSymbol(characters[y][x-1]) || isSymbol(characters[y][x+1]) || isSymbol(characters[y+1][x-1]) || isSymbol(characters[y+1][x]) || isSymbol(characters[y+1][x+1])
	}

	for y, l := range characters {
		for x, c := range l {
			if (!inNumber && !isDigit(c)) || (inNumber && !currentTouchesSymbol && !isDigit(c)) {
				reset()
				continue
			}
			if inNumber && !isDigit(c) && currentTouchesSymbol {
				pn, err := strconv.Atoi(currentNumber)
				if err != nil {
					panic(fmt.Sprintf("y: %d, x: %d, err: %v", y, x, err))
				}
				partNumbers = append(partNumbers, pn)
				reset()
				continue
			}
			if inNumber && isDigit(c) {
				currentNumber += c
				if !currentTouchesSymbol && touchesSymbol(y, x) {
					currentTouchesSymbol = true
				}
				continue
			}
			if isDigit(c) {
				inNumber = true
				currentNumber += c
				currentTouchesSymbol = touchesSymbol(y, x)
			}
		}
		if inNumber && currentTouchesSymbol {
			pn, err := strconv.Atoi(currentNumber)
			if err != nil {
				panic(fmt.Sprintf("y: %d, x: %d, err: %v", y, width, err))
			}
			partNumbers = append(partNumbers, pn)
		}
		reset()
	}

	return partNumbers
}

func isDigit(c string) bool {
	return c == "1" ||
		c == "2" ||
		c == "3" ||
		c == "4" ||
		c == "5" ||
		c == "6" ||
		c == "7" ||
		c == "8" ||
		c == "9" ||
		c == "0"
}

func isSymbol(c string) bool {
	return !isDigit(c) && !(c == ".")
}

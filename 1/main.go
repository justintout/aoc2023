package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var firstDigitRegex = regexp.MustCompile("[0-9]")

var firstDigitWordRegex = regexp.MustCompile("[0-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|(zero)")
var firstDigitWordReversedRegex = regexp.MustCompile("[0-9]|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)|(orez)")
var numbers map[string]string = map[string]string{
	"one":   "1",
	"eno":   "1",
	"two":   "2",
	"owt":   "2",
	"three": "3",
	"eerht": "3",
	"four":  "4",
	"ruof":  "4",
	"five":  "5",
	"evif":  "5",
	"six":   "6",
	"xis":   "6",
	"seven": "7",
	"neves": "7",
	"eight": "8",
	"thgie": "8",
	"nine":  "9",
	"enin":  "9",
	"zero":  "0",
	"orez":  "0",
}

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
	for _, line := range lines {
		f := firstDigitRegex.FindString(line)
		l := firstDigitRegex.FindString(reverse(line))
		n, err := strconv.Atoi(f + l)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func part2(lines []string) (total int, err error) {
	for _, line := range lines {
		f := digitize(firstDigitWordRegex.FindString(line))
		l := digitize(firstDigitWordReversedRegex.FindString(reverse(line)))
		n, err := strconv.Atoi(f + l)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func digitize(s string) string {
	d, ok := numbers[s]
	if !ok {
		return s
	}
	return d
}

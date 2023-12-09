package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RaceFromEnd(time int, pressing int, distance int, direction int) int {
	run := (time - pressing) * pressing
	if run <= distance {
		return RaceFromEnd(time, pressing+direction, distance, direction)
	}
	return pressing
}

func BinaryRace(time int, low int, high int, distance int, direction int) int {
	pressing := (low + high) / 2
	run := (time - pressing) * pressing
	next := (time - (pressing + direction)) * (pressing + direction)
	prev := (time - (pressing - direction)) * (pressing - direction)

	if run > distance {
		if next <= distance {
			return pressing
		}

		l := low
		h := high
		if direction > 0 {
			l = pressing
		} else {
			h = pressing
		}
		return BinaryRace(time, l, h, distance, direction)
	}

	if prev > distance {
		return pressing - direction
	}

	l := low
	h := high
	if direction > 0 {
		h = pressing
	} else {
		l = pressing
	}
	return BinaryRace(time, l, h, distance, direction)
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	var p1t []int
	var p1d []int
	p1sum, p2sum := 1, 1

	re := regexp.MustCompile(`\b\d+\b`)

	for _, s := range re.FindAllString(lines[0], -1) {
		n, _ := strconv.Atoi(s)
		p1t = append(p1t, n)
	}

	for _, s := range re.FindAllString(lines[1], -1) {
		n, _ := strconv.Atoi(s)
		p1d = append(p1d, n)
	}

	for i, t := range p1t {
		hi := BinaryRace(t, t/2, t, p1d[i], 1)
		lo := BinaryRace(t, 1, t/2, p1d[i], -1)
		p1sum *= hi - lo + 1
	}

	p2t, _ := strconv.Atoi(strings.Join(re.FindAllString(lines[0], -1), ""))
	p2d, _ := strconv.Atoi(strings.Join(re.FindAllString(lines[1], -1), ""))

	hi := BinaryRace(p2t, p2t/2, p2t, p2d, 1)
	lo := BinaryRace(p2t, 1, p2t/2, p2d, -1)
	p2sum *= hi - lo + 1

	fmt.Printf("Part 1: %d\nPart2: %d\n", p1sum, p2sum)

}

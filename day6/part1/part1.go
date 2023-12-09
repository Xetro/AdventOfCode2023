package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func race(time int, pressing int, distance int, direction int) int {
	run := (time - pressing) * pressing
	if run <= distance {
		return race(time, pressing+direction, distance, direction)
	}
	return pressing
}

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	var time []int
	var distance []int

	re := regexp.MustCompile(`\b\d+\b`)

	for _, s := range re.FindAllString(lines[0], -1) {
		n, _ := strconv.Atoi(s)
		time = append(time, n)
	}

	for _, s := range re.FindAllString(lines[1], -1) {
		n, _ := strconv.Atoi(s)
		distance = append(distance, n)
	}

	sum := 1

	for i, t := range time {
		high := race(t, t-1, distance[i], -1)
		low := race(t, 1, distance[i], 1)

		sum *= high - low + 1
	}
	fmt.Println(sum)
}

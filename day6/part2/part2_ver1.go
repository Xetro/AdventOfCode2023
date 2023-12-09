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

	re := regexp.MustCompile(`\b\d+\b`)

	time, _ := strconv.Atoi(strings.Join(re.FindAllString(lines[0], -1), ""))
	distance, _ := strconv.Atoi(strings.Join(re.FindAllString(lines[1], -1), ""))
	sum := 1

	high := race(time, time-1, distance, -1)
	low := race(time, 1, distance, 1)

	sum *= high - low + 1
	fmt.Println(sum)
}

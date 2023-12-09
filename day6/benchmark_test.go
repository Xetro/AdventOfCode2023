package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var gLo, gHi int

func setup() (int, int) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	re := regexp.MustCompile(`\b\d+\b`)

	time, _ := strconv.Atoi(strings.Join(re.FindAllString(lines[0], -1), ""))
	distance, _ := strconv.Atoi(strings.Join(re.FindAllString(lines[1], -1), ""))

	return time, distance
}
func BenchmarkRaceFromEnd(b *testing.B) {
	time, distance := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start := RaceFromEnd(time, time-1, distance, -1)
		end := RaceFromEnd(time, 1, distance, 1)
		gLo, gHi = start, end
	}
}
func BenchmarkBinaryRace(b *testing.B) {
	time, distance := setup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		low := BinaryRace(time, 1, time/2, distance, -1)
		high := BinaryRace(time, time/2, time, distance, 1)
		gLo, gHi = low, high
	}
}

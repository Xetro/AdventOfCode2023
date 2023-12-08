package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func mapToLocation(seed int, maps []string) (loc int) {
	lines := strings.Split(maps[0], "\n")
	loc = seed

	for _, cur := range lines {
		nums := strings.Fields(cur)

		dest, _ := strconv.Atoi(nums[0])
		source, _ := strconv.Atoi(nums[1])
		r, _ := strconv.Atoi(nums[2])

		if seed >= source && seed <= (source+r) {
			loc = dest + seed - source
			break
		}
	}

	if len(maps) > 1 {
		return mapToLocation(loc, maps[1:])
	}
	return
}

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n\n")

	for i, line := range lines {
		split := strings.Split(line, ":")
		nums := strings.TrimSpace(split[1])
		lines[i] = nums
	}

	seeds := strings.Fields(lines[0])
	maps := lines[1:]
	locations := make([]float64, 0, 10)

	for _, s := range seeds {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		locations = append(locations, float64(mapToLocation(n, maps)))
	}

	fmt.Println(int(slices.Min(locations)))
}

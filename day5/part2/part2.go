package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type seed struct {
	start, end int
}

func mapToLocation(seeds []seed, maps []string) (loc []seed) {
	lines := strings.Split(maps[0], "\n")
	for _, cur := range lines {
		nums := strings.Fields(cur)

		dest, _ := strconv.Atoi(nums[0])
		source, _ := strconv.Atoi(nums[1])
		r, _ := strconv.Atoi(nums[2])

		remaining := make([]seed, 0, len(seeds)/2)

		for _, s := range seeds {
			if !(s.start <= source+r && s.end >= source) {
				remaining = append(remaining, s)
			} else {
				loc = append(loc, seed{max(dest+s.start-source, dest), min(dest+s.end-source, dest+r-1)})
				if s.start < source && s.end > source+r-1 {
					remaining = append(remaining, seed{source + r, s.end}, seed{s.start, source - 1})
				} else if s.start < source {
					remaining = append(remaining, seed{s.start, source - 1})
				} else if s.end > source+r {
					remaining = append(remaining, seed{source + r, s.end})
				}
			}
		}
		seeds = remaining
	}
	loc = append(loc, seeds...)

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

	fields := strings.Fields(lines[0])
	seeds := make([]seed, 0, len(fields))

	for i := 0; i < len(fields); i += 2 {
		s, _ := strconv.Atoi(fields[i])
		r, _ := strconv.Atoi(fields[i+1])

		seeds = append(seeds, seed{s, s + r - 1})
	}
	maps := lines[1:]

	locations := mapToLocation(seeds, maps)

	for _, s := range seeds {
		locations = append(locations, mapToLocation([]seed{s}, maps)...)
	}

	result := slices.MinFunc(locations, func(a, b seed) int {
		return cmp.Compare(a.start, b.start)
	})

	fmt.Println(result.start)
}

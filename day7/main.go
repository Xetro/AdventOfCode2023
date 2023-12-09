package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type match struct {
	count int
	char  rune
}

type hand struct {
	hand     string
	strength int
	bid      int
}

var handStrength = map[int]int{
	2: 2,
	3: 4,
	4: 6,
	5: 7,
	6: 3,
	7: 5,
}

var cardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	hands := make([]hand, 0, len(lines))
	for _, l := range lines {
		split := strings.Fields(l)
		h := split[0]
		bid, _ := strconv.Atoi(split[1])

		matches := make(map[rune]int)
		for _, c := range h {
			matches[c]++
		}

		pairs := make([]match, 0)
		for card, count := range matches {
			if count > 1 {
				pairs = append(pairs, match{count, card})
			}
		}

		var hs int
		switch len(pairs) {
		case 2:
			if pairs[0].count == pairs[1].count {
				// 2 pair
				hs = handStrength[6]
				break
			}
			// Full house
			hs = handStrength[7]
		case 1:
			hs = handStrength[pairs[0].count]
		default:
			hs = 1
		}

		hands = append(hands, hand{h, hs, bid})

		sort.Slice(hands, func(i, j int) bool {
			str1 := hands[i].strength
			str2 := hands[j].strength

			if str1 != str2 {
				return str1 < str2
			}

			for n, c := range hands[i].hand {
				str1 = cardStrength[c]
				str2 = cardStrength[rune(hands[j].hand[n])]

				if str1 != str2 {
					return str1 < str2
				}
			}
			return false
		})
	}

	sum := 0

	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	fmt.Println("Part 1: ", sum)

}

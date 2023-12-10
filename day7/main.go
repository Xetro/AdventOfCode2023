package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	TWO_PAIR   = "twopair"
	FULL_HOUSE = "fullhouse"
	FOUR_KIND  = "4"
	FIVE_KIND  = "5"
)

var handStrength = map[string]int{
	"1":         1,
	"2":         2,
	"3":         4,
	"4":         6,
	"5":         7,
	"twopair":   3,
	"fullhouse": 5,
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

type match struct {
	count int
	char  rune
}

type hand struct {
	hand     string
	strength int
	bid      int
}

type handsSlice []hand

func (s handsSlice) Less(i, j int) bool {
	strn1 := s[i].strength
	strn2 := s[j].strength

	if strn1 != strn2 {
		return strn1 < strn2
	}

	for n, c := range s[i].hand {
		strn1 = cardStrength[c]
		strn2 = cardStrength[rune(s[j].hand[n])]

		if strn1 != strn2 {
			return strn1 < strn2
		}
	}
	return false
}

func findHandStrengthPart2(sets []match, joker int, h string) int {
	jokerSets := 0
	threeKind := false

	for _, i := range sets {
		if i.char == 'J' {
			jokerSets++
		}
		if i.count == 3 {
			threeKind = true
		}
	}

	if len(sets) < 1 {
		return handStrength[strconv.Itoa(joker+1)]
	}

	if len(sets) > 1 {
		switch {
		case joker == 3, joker == 2 && threeKind:
			return handStrength[FIVE_KIND]
		case threeKind, joker == 1:
			return handStrength[FULL_HOUSE]
		case joker == 0:
			return handStrength[TWO_PAIR]
		default:
			return handStrength[FOUR_KIND]
		}
	}

	if joker == 5 {
		joker = 0
	} else if joker > 0 {
		joker = 1
	}
	return handStrength[strconv.Itoa(sets[0].count+joker)]
}

func findHandStrengthPart1(sets []match, h string) int {
	switch {
	case len(sets) == 2 && sets[0].count == sets[1].count:
		return handStrength[TWO_PAIR]
	case len(sets) == 2:
		return handStrength[FULL_HOUSE]
	case len(sets) == 1:
		return handStrength[strconv.Itoa(sets[0].count)]
	default:
		return 1
	}
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")

	handsP1 := make(handsSlice, 0, len(lines))
	handsP2 := make(handsSlice, 0, len(lines))
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

		jCount := 0
		for _, char := range h {
			if char == 'J' {
				jCount++
			}
		}

		hsP1 := findHandStrengthPart1(pairs, h)
		hsP2 := findHandStrengthPart2(pairs, jCount, h)
		handsP1 = append(handsP1, hand{h, hsP1, bid})
		handsP2 = append(handsP2, hand{h, hsP2, bid})
	}

	sort.Slice(handsP1, handsP1.Less)

	cardStrength['J'] = 0
	sort.Slice(handsP2, handsP2.Less)

	sumP1 := 0
	sumP2 := 0

	for i, h := range handsP1 {
		sumP1 += h.bid * (i + 1)
	}

	for i, h := range handsP2 {
		sumP2 += h.bid * (i + 1)
	}

	fmt.Printf("Part 1: %d\nPart2: %d\n ", sumP1, sumP2)

}

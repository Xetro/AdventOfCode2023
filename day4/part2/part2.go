package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var cards = make([]int, 0)
var winningsPerCard = make(map[int]int)
var sum = 0

func traverse(hits int, card int) int {
	if _, ok := winningsPerCard[card]; ok {
		return winningsPerCard[card]
	}

	count := 1
	for i := 1; i <= hits; i++ {
		next := card + i
		if _, ok := winningsPerCard[next]; !ok {
			count += traverse(cards[next], next)
		} else {
			count += winningsPerCard[next]
		}
	}
	winningsPerCard[card] = count
	sum += count
	return count
}

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		picked := make(map[string]bool, 8)
		hits := 0

		tokens := strings.FieldsFunc(text, func(r rune) bool {
			return r == ':' || r == '|'
		})

		for _, s := range strings.Fields(tokens[2]) {
			picked[s] = true
		}

		for _, s := range strings.Fields(tokens[1]) {
			if _, ok := picked[s]; ok {
				hits++
			}
		}
		cards = append(cards, hits)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, v := range cards {
		traverse(v, i)
	}
	fmt.Println("Sum of Ids: ", sum)
}

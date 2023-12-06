package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			continue
		}

		picked := make(map[string]bool, 8)
		points := 0

		fmt.Println("text: ", text)
		tokens := strings.FieldsFunc(text, func(r rune) bool {
			return r == ':' || r == '|'
		})

		for _, s := range strings.Fields(tokens[2]) {
			picked[s] = true
		}

		for _, s := range strings.Fields(tokens[1]) {
			if _, ok := picked[s]; ok {
				if points == 0 {
					points = 1
				} else {
					points = points << 1
				}
			}
		}
		sum += points
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of Ids: ", sum)
}

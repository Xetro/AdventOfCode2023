package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int

	for scanner.Scan() {
		var firstFound, lastFound bool
		var first, last string
		var final int

		fmt.Println(scanner.Text())
		for _, ch := range scanner.Text() {
			if unicode.IsDigit(ch) {
				if !firstFound {
					firstFound = true
					first = string(ch)
				} else {
					lastFound = true
					last = string(ch)
				}
			}
		}

		if lastFound {
			final, _ = strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		} else {
			final, _ = strconv.Atoi(fmt.Sprintf("%s%s", first, first))
		}

		sum += final
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

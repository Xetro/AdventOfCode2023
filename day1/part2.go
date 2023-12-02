package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

var textToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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
		var first, last, found string
		var final int

		fmt.Println(scanner.Text())
		line := scanner.Text()
		fmt.Println("length: ", len(line))
		for i, ch := range line {
			lenThree := int(math.Min(float64(i+3), float64(len(line))))
			lenFour := int(math.Min(float64(i+4), float64(len(line))))
			lenFive := int(math.Min(float64(i+5), float64(len(line))))

			// fmt.Println(line[i:lenThree], line[i:lenFour], line[i:lenFive])
			if unicode.IsDigit(ch) {
				found = string(ch)
				fmt.Println("found is digit: ", found, ch)
			} else if v, ok := textToInt[line[i:lenThree]]; ok {
				// one, two, six
				found = fmt.Sprint(v)
				fmt.Println("found is one two six: ", found, v)
			} else if v, ok := textToInt[line[i:lenFour]]; ok {
				// four, five, nine
				found = fmt.Sprint(v)
				fmt.Println("found is four five nine: ", found, v)
			} else if v, ok := textToInt[line[i:lenFive]]; ok {
				// three, seven, eight
				found = fmt.Sprint(v)
				fmt.Println("found is three seven eight: ", found, v)
			}

			if found != "" {
				if !firstFound {
					firstFound = true
					first = found
				} else {
					lastFound = true
					last = found
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

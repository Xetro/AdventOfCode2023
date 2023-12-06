package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type part struct {
	start  int
	end    int
	number int
}

type symbol struct {
	start int
	end   int
	row   int
	char  rune
}

func main() {
	sum := 0
	text := make([]string, 0, 140)

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			text = append(text, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	parts := make([][]part, len(text))
	symbols := make([]symbol, 0)

	for row, line := range text {

		p := part{
			start:  -1,
			end:    0,
			number: 0,
		}

		for i, r := range line {
			switch {
			case !unicode.IsDigit(r) && r != '.':
				s := symbol{
					start: max(0, i-1),
					end:   min(len(line), i+1),
					row:   row,
					char:  r,
				}
				symbols = append(symbols, s)

				if p.start != -1 {
					p.end = i - 1
					parts[row] = append(parts[row], p)
					p = part{
						start:  -1,
						end:    0,
						number: 0,
					}
				}
			case unicode.IsDigit(r):
				if p.start == -1 {
					p.start = i
				}
				p.number = (p.number * 10) + int(r-'0')

			case p.start != -1:
				p.end = i - 1
				parts[row] = append(parts[row], p)
				p = part{
					start:  -1,
					end:    0,
					number: 0,
				}
			}
		}
		if p.start != -1 {
			p.end = len(line)
			parts[row] = append(parts[row], p)
		}
	}

	for _, s := range symbols {
		adjacent := 0
		ratio := 1

		for _, partsRow := range parts[max(0, s.row-1):min(len(parts), s.row+2)] {
			for _, p := range partsRow {
				if p.start <= s.end && p.end >= s.start {
					adjacent += 1
					ratio *= p.number
				}
			}
		}

		if adjacent == 2 {
			sum += ratio
		}
	}
	fmt.Println("Sum of Ids: ", sum)
}

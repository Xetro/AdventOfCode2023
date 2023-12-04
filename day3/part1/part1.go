package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var possibleCubes map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		isValid := true
		text := scanner.Text()

		if text == "" {
			continue
		}

		for i := 0; i <= len(text); {
      if unicode.IsDigit(r) {
        for n := Min(i+1, len(text)); unicode.IsDigit(text[r+n]); {
           
      
        }
        
      

			)
		}

		for i, r := range text {
			unicode.IsDigit(r)
		}

		game := strings.FieldsFunc(text, func(r rune) bool {
			return r == ':' || r == ';'
		})

		gameId, _ := strconv.Atoi(strings.Fields(game[0])[1])

		sets := game[1:]

		for _, set := range sets {
			cubes := strings.FieldsFunc(set, func(r rune) bool {
				return r == ','
			})

			for _, cube := range cubes {
				cubeFields := strings.Fields(cube)
				count, _ := strconv.Atoi(cubeFields[0])
				color := cubeFields[1]

				if count > possibleCubes[color] {
					fmt.Printf("%s %d not valid, max posible %d\n", color, count, possibleCubes[color])
					isValid = false
					break
				}
			}

			if !isValid {
				fmt.Println("Not adding to total")
				break
			}
		}

		if isValid {
			sum += gameId
			fmt.Println("Adding", gameId, "to new total", sum)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of Ids: ", sum)
}

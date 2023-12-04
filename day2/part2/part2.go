package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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

		cubesPerColor := map[string][]int{
			"red":   {},
			"green": {},
			"blue":  {},
		}

		text := scanner.Text()

		if text == "" {
			continue
		}

		fmt.Println("text: ", text)
		game := strings.FieldsFunc(text, func(r rune) bool {
			return r == ':' || r == ';'
		})

		sets := game[1:]

		for _, set := range sets {
			cubes := strings.FieldsFunc(set, func(r rune) bool {
				return r == ','
			})

			for _, cube := range cubes {
				cubeFields := strings.Fields(cube)
				count, _ := strconv.Atoi(cubeFields[0])
				color := cubeFields[1]
				cubesPerColor[color] = append(cubesPerColor[color], count)
			}
		}

		fmt.Println("All cubeS? ", cubesPerColor)

		power := 1
		for _, cubes := range cubesPerColor {
			power *= slices.Max(cubes)
		}

		sum += power
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum of Ids: ", sum)
}

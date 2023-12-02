package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cubeCountbyColor := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	possibleGamesSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := scanner.Text()
		gameSplit := strings.Split(game, ": ")
		gameId, reveals := atoi(strings.Split(gameSplit[0], " ")[1]), strings.Split(gameSplit[1], "; ")

		feasible := true
		for ri, r := range reveals {
			colorsWithCounts := strings.Split(r, ", ")
			for ci, c := range colorsWithCounts {
				split := strings.Split(c, " ")
				count, color := atoi(split[0]), split[1]
				fmt.Printf("game %3d %3d[%3d]: %d %s <> %d\n", gameId, ri, ci, count, color, cubeCountbyColor[color])
				if count > cubeCountbyColor[color] {
					feasible = false
					break
				}
			}
		}

		if feasible {
			possibleGamesSum += gameId
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(possibleGamesSum)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tokens := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	counter := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		firstPos := 10000
		firstVal := 0
		for token, num := range tokens {
			ix := strings.Index(line, token)
			if ix != -1 && ix < firstPos {
				firstPos = ix
				firstVal = num
			}
		}
		lastPos := -1
		lastVal := 0
		for token, num := range tokens {
			ix := strings.LastIndex(line, token)
			if ix != -1 && ix > lastPos {
				lastPos = ix
				lastVal = num
			}
		}

		counter += firstVal*10 + lastVal
		fmt.Println(firstVal*10 + lastVal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(counter)
}

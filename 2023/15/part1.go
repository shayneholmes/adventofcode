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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		strs := strings.Split(line, ",")

		vals := 0
		for _, s := range strs {
			val := 0
			for _, ch := range s {
				// Determine the ASCII code for the current character of the string.
				// Increase the current value by the ASCII code you just determined.
				val += int(byte(ch))
				// Set the current value to itself multiplied by 17.
				val *= 17
				// Set the current value to the remainder of dividing itself by 256.
				val %= 256
			}
			vals += val
		}
		fmt.Println(vals)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

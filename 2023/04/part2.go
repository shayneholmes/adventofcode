package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	id := 0
	extraCopiesById := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		lineparts := strings.Split(line, ": ")
		data := lineparts[1]
		dataparts := strings.Split(data, " | ")
		mine := strings.Split(dataparts[0], " ")
		winning := strings.Split(dataparts[1], " ")
		mineMap := map[int]bool{}
		for _, m := range mine {
			if m == "" {
				continue
			}
			mineMap[atoi(m)] = true
		}
		matches := 0
		for _, w := range winning {
			if w == "" {
				continue
			}
			if _, ok := mineMap[atoi(w)]; ok {
				matches += 1
			}
		}
		score := 0
		copies := 1 + extraCopiesById[id]
		if matches >= 1 {
			score = 1 << (matches - 1)
			for i := 0; i < matches; i++ {
				extraCopiesById[id+1+i] += copies
			}
		}
		id += 1
		sum += copies
		fmt.Printf("%d after %d copies of card with score %d\n", sum, copies, score)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// cheap atoi
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func multiply(values ...int) int {
	if len(values) == 0 {
		log.Fatalf("no value in sum function")
	}

	product := 1
	for _, value := range values {
		product *= value
	}
	return product
}

func sum(values ...int) int {
	if len(values) == 0 {
		panic("no value in sum function")
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

type nullableInt struct {
	value int
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("no value in max function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || value > max.value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("no value in min function")
	}

	var min *nullableInt
	for _, value := range values {
		if min == nil || value < min.value {
			min = &nullableInt{value}
		}
	}
	return min.value
}

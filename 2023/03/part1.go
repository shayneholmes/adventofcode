package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type loc = struct {
	R    int
	C    int
	Size int
}

func isSym(c byte) bool {
	return (c < '0' || c > '9') && c != '.'
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := [][]byte{}
	numbers := map[loc]int{}

	scanner := bufio.NewScanner(file)
	r := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []byte(line))
		makingNum := false
		num := 0
		size := 0
		for i, c := range line {
			if c >= '0' && c <= '9' {
				if !makingNum {
					makingNum = true
					size = 0
				}
				size += 1
				num = num*10 + (int(c) - '0')
			} else {
				if makingNum {
					numbers[loc{R: r, C: i - size, Size: size}] = num
					makingNum = false
					num = 0
				}
			}
		}
		if makingNum { // end of row number
			fmt.Printf("Found %d at end of row: %s ", num, line[len(line)-size:])
			numbers[loc{R: r, C: len(line) - size, Size: size}] = num
			makingNum = false
			num = 0
		}
		r += 1
	}

	hasSymbolNeighbor := func(loc loc) bool {
		r1 := max(0, loc.R-1)
		r2 := min(len(lines)-1, loc.R+1)
		c1 := max(0, loc.C-1)
		c2 := min(len(lines[0])-1, loc.C+loc.Size)
		fmt.Printf("%v: searching %d-%d, %d-%d\n", loc, r1, r2, c1, c2)
		for r := r1; r <= r2; r++ {
			for c := c1; c <= c2; c++ {
				if isSym(lines[r][c]) {
					fmt.Printf("found %q at %d,%d\n", lines[r][c], r, c)
					return true
				}
			}
		}
		return false
	}

	sum := 0
	for loc, num := range numbers {
		if hasSymbolNeighbor(loc) {
			sum += num
			for i := 0; i < loc.Size; i++ {
				lines[loc.R][loc.C+i] = '9'
			}
		} else {
			for i := 0; i < loc.Size; i++ {
				lines[loc.R][loc.C+i] = '0'
			}
		}
	}

	for _, r := range lines {
		fmt.Printf("%s\n", r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
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

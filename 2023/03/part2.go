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

func isStar(c byte) bool {
	return c == '*'
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
	starNeighbors := map[int][]int{}

	hasSymbolNeighbor := func(loc loc, num int) bool {
		r1 := max(0, loc.R-1)
		r2 := min(len(lines)-1, loc.R+1)
		c1 := max(0, loc.C-1)
		c2 := min(len(lines[0])-1, loc.C+loc.Size)
		fmt.Printf("%v: searching %d-%d, %d-%d\n", loc, r1, r2, c1, c2)
		found := false
		for r := r1; r <= r2; r++ {
			for c := c1; c <= c2; c++ {
				if isStar(lines[r][c]) {
					// Add a new neighbor
					old, ok := starNeighbors[r*10000+c]
					if !ok {
						old = []int{}
					}
					starNeighbors[r*10000+c] = append(old, num)
					found = true
				}
			}
		}
		return found
	}

	sum := 0
	for loc, num := range numbers {
		if hasSymbolNeighbor(loc, num) {
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

	ratios := 0
	for _, neighbors := range starNeighbors {
		if len(neighbors) == 2 {
			ratios += neighbors[0] * neighbors[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ratios)
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

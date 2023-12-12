package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type point = struct {
	r int
	c int
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func dist(p1, p2 point) int {
	return abs(p1.r-p2.r) + abs(p1.c-p2.c)
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	space := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		space = append(space, []byte(line))
	}

	// map the universe
	galaxies := make([]point, 0, 1000)

	rowsize := len(space)
	colsize := len(space[0])
	nonemptyrows := make([]bool, rowsize)
	nonemptycols := make([]bool, colsize)
	for i := 0; i < rowsize; i++ {
		for j := 0; j < colsize; j++ {
			if space[i][j] == '#' {
				galaxies = append(galaxies, point{i, j})
				if !nonemptyrows[i] {
					nonemptyrows[i] = true
				}
				if !nonemptycols[j] {
					nonemptycols[j] = true
				}
			}
		}
	}

	// populate range tables
	emptyrowscum := make([]int, rowsize)
	{
		emptyrows := 0
		for i := 0; i < rowsize; i++ {
			if !nonemptyrows[i] {
				emptyrows += 1
			}
			emptyrowscum[i] = emptyrows
		}
	}
	emptycolscum := make([]int, colsize)
	{
		emptycols := 0
		for j := 0; j < colsize; j++ {
			if !nonemptycols[j] {
				emptycols += 1
			}
			emptycolscum[j] = emptycols
		}
	}

	gapsize := 1000000
	adjustedgalaxies := make([]point, len(galaxies))
	for i, g := range galaxies {
		rowgaps := emptyrowscum[g.r]
		colgaps := emptycolscum[g.c]
		nr := g.r + rowgaps*(gapsize-1)
		nc := g.c + colgaps*(gapsize-1)
		adjustedgalaxies[i] = point{nr, nc}
	}

	sum := 0
	for i, g1 := range adjustedgalaxies {
		for _, g2 := range adjustedgalaxies[i+1:] {
			sum += dist(g1, g2)
		}
	}
	fmt.Println(sum)
}

// return all ints in a string of text
func ints(s string) []int {
	re := regexp.MustCompile(`-?\b\d+\b`)
	intss := re.FindAllString(s, -1)
	ints := make([]int, len(intss))
	for i := 0; i < len(ints); i++ {
		ints[i] = atoi(intss[i])
	}
	return ints
}

// return all floats in a string of text
func floats(s string) []float64 {
	re := regexp.MustCompile(`-?\d+(\.\d+)?`)
	intss := re.FindAllString(s, -1)
	floats := make([]float64, len(intss))
	for i := 0; i < len(floats); i++ {
		f, err := strconv.ParseFloat(intss[i], 64)
		if err != nil {
			log.Fatalf("not an float: %q", s)
		}
		floats[i] = f
	}
	return floats
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type dir = int

const (
	north = iota
	east
	south
	west
	none
)

var dirs = []dir{north, east, south, west}

type combo struct {
	pipe byte
	dir  int
}

func getDirFromPipeGivenDir(pipe byte, dir dir) int {
	switch (combo{pipe, dir}) {
	// | is a vertical pipe connecting north and south.
	case combo{'|', south}:
		return south
	case combo{'|', north}:
		return north
	// - is a horizontal pipe connecting east and west.
	case combo{'-', west}:
		return west
	case combo{'-', east}:
		return east
	// L is a 90-degree bend connecting north and east.
	case combo{'L', west}:
		return north
	case combo{'L', south}:
		return east
	// J is a 90-degree bend connecting north and west.
	case combo{'J', east}:
		return north
	case combo{'J', south}:
		return west
	// 7 is a 90-degree bend connecting south and west.
	case combo{'7', north}:
		return west
	case combo{'7', east}:
		return south
	// F is a 90-degree bend connecting south and east.
	case combo{'F', west}:
		return south
	case combo{'F', north}:
		return east
		// . is ground; there is no pipe in this tile.
	default:
		return none
	}
}

type pos struct {
	r int
	c int
}

func getPosGivenDir(orig pos, dir dir) pos {
	switch dir {
	case north:
		return pos{r: orig.r - 1, c: orig.c}
	case south:
		return pos{r: orig.r + 1, c: orig.c}
	case east:
		return pos{r: orig.r, c: orig.c + 1}
	case west:
		return pos{r: orig.r, c: orig.c - 1}
	}
	log.Fatalf("Bad direction %q", dir)
	return pos{}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pipes := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pipes = append(pipes, []byte(line))
		fmt.Println(line)
	}

	// find start
	start := pos{}
findstart:
	for r := range pipes {
		for c := range pipes[r] {
			if pipes[r][c] == 'S' {
				start = pos{r, c}
				break findstart
			}
		}
	}
	fmt.Printf("start: %v\n", start)

	isPipe := map[pos]bool{}
	rightSide := map[pos]bool{}

	// Try starting in each direction
	{
		dir := south
		steps := 0
		p := start
		for dir != none {
			isPipe[p] = true
			p = getPosGivenDir(p, dir)
			oldDir := dir
			dir = getDirFromPipeGivenDir(pipes[p.r][p.c], dir)
			right := getPosGivenDir(p, (dir+1)%4)
			rightSide[right] = true
			rightnu := getPosGivenDir(p, (oldDir+1)%4)
			rightSide[rightnu] = true
			steps += 1
		}
		fmt.Printf("Ended after %v steps at %v (%q)\n", steps, p, pipes[p.r][p.c])
	}

	area := 0
	inside := map[pos]bool{}

	var flood func(p pos)
	flood = func(p pos) {
		if inside[p] || isPipe[p] {
			return
		}
		area += 1
		inside[p] = true
		pipes[p.r][p.c] = '&'
		for _, dir := range dirs {
			flood(getPosGivenDir(p, dir))
		}
	}

	for r := 0; r < len(pipes); r++ {
		in := false
		lastWasPipe := false
		for c := 0; c < len(pipes[r]); c++ {
			p := pos{r, c}
			thisIsPipe := isPipe[p]
			if rightSide[p] {
				// flood fill
				flood(p)
			}
			if lastWasPipe != thisIsPipe {
				in = !in
			}
			if thisIsPipe {
				pipes[r][c] = '*'
			}
			lastWasPipe = thisIsPipe
		}
		fmt.Printf("%s\n", pipes[r])
	}
	fmt.Println(area)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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

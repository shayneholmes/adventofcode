package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// location in a grid
type loc = struct {
	r int
	c int
}

type instruction struct {
	dir loc
	mag int
}

var (
	north = loc{-1, 0}
	east  = loc{0, 1}
	west  = loc{0, -1}
	south = loc{1, 0}
)

func add(i, j loc) loc {
	return loc{i.r + j.r, i.c + j.c}
}

func parsedir(s string) loc {
	switch s {
	case "U":
		return north
	case "L":
		return west
	case "R":
		return east
	case "D":
		return south
	}
	log.Fatalf("couldn't parse %v", s)
	return north
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := make([]instruction, 0)

	perimeter := 0
	scanner := bufio.NewScanner(file)
	minc, minr, maxc, maxr := 0, 0, 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		dir := parsedir(splits[0])
		mag := atoi(splits[1])
		instructions = append(instructions, instruction{dir, mag})
	}

	pos := loc{0, 0}
	locs := make([]loc, 0, 700)
	locs = append(locs, pos)
	for _, inst := range instructions {
		for i := 0; i < inst.mag; i++ {
			pos = add(pos, inst.dir)
		}
		if pos.c > maxc {
			maxc = pos.c
		}
		if pos.r > maxr {
			maxr = pos.r
		}
		if pos.c < minc {
			minc = pos.c
		}
		if pos.r < minr {
			minr = pos.r
		}
		perimeter += inst.mag
		locs = append(locs, pos)
	}

	minc -= 1 // make room for fill
	minr -= 1
	maxr += 1
	maxc += 1

	grid := make([][]int, maxr+1-minr)
	for i := range grid {
		grid[i] = make([]int, maxc+1-minc)
	}
	set := func(loc loc, x int) {
		r := loc.r - minr
		c := loc.c - minc
		grid[r][c] = x
	}
	pos = loc{0, 0}
	for _, inst := range instructions {
		set(pos, 1)
		for i := 0; i < inst.mag; i++ {
			pos = add(pos, inst.dir)
			set(pos, 1)
		}
	}

	// flood fill
	areaoutside := 1
	{
		start := loc{0, 0}
		visited := map[loc]bool{}
		tovisit := []loc{start}
		for len(tovisit) > 0 {
			cur := tovisit[0]
			tovisit = tovisit[1:]
			if visited[cur] {
				continue
			}
			visited[cur] = true
			for _, dir := range []loc{north, south, east, west} {
				nu := add(cur, dir)
				if nu.r < 0 || nu.c < 0 {
					continue
				}
				if nu.r >= len(grid) || nu.c >= len(grid[0]) {
					continue
				}
				if grid[nu.r][nu.c] == 0 && !visited[nu] {
					tovisit = append(tovisit, nu)
					grid[nu.r][nu.c] = 2
					areaoutside += 1
				}
			}
		}
	}

	for i := range grid {
		for j := range grid[0] {
			fmt.Printf("%d", grid[i][j])
		}
		fmt.Printf("\n")
	}

	areainside := (len(grid) * len(grid[0])) - areaoutside

	fmt.Println(areainside)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
}

type point = []float64

// manhattan distance
func mdist[S Number](p1, p2 []S) S {
	if len(p1) != len(p2) {
		log.Fatalf("mismatched dimensions in %v, %v", p1, p2)
	}
	var dist S
	for i := range p1 {
		dist += abs(p1[i] - p2[i])
	}
	return dist
}

// euclidean distance
func dist[S Number](p1, p2 []S) float64 {
	if len(p1) != len(p2) {
		log.Fatalf("mismatched dimensions in %v, %v", p1, p2)
	}
	var squares S
	for i := range p1 {
		diff := (p1[i] - p2[i])
		squares += diff * diff
	}
	return math.Sqrt(float64(squares))
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

// greatest common divisor
func gcd(values ...int) int {
	if len(values) == 0 {
		panic("no value in gcd function")
	}
	gcd := values[0]
	for _, i := range values[1:] {
		gcd = gcd_(gcd, i)
	}
	return gcd
}

func gcd_(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd_(b, a%b)
}

func lcm(values ...int) int {
	if len(values) == 0 {
		panic("no value in lcm function")
	}
	lcm := values[0]
	for _, i := range values[1:] {
		lcm = lcm_(lcm, i)
	}
	return lcm
}

func lcm_(a, b int) int {
	return a / gcd_(a, b) * b
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

// Types from constraints package
type (
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Integer interface{ Signed | Unsigned }
	Float   interface{ ~float32 | ~float64 }
	Complex interface{ ~complex64 | ~complex128 }
	Number  interface{ Integer | Float }
	Ordered interface{ Integer | Float | ~string }
)

// Wrong: 45160

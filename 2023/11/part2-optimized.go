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
		fmt.Println(line)
	}

	// map the universe
	galaxies := make([]point, 0, 1000)

	rowsize := len(space)
	colsize := len(space[0])
	nonemptyrows := map[int]bool{}
	nonemptycols := map[int]bool{}
	for i := 0; i < rowsize; i++ {
		for j := 0; j < colsize; j++ {
			if space[i][j] == '#' {
				galaxies = append(galaxies, point{i, j})
				nonemptyrows[i] = true
				nonemptycols[j] = true
			}
		}
	}

	// populate range tables
	emptyrowscum := make([]int, rowsize)
	{
		emptyrows := 0
		for i := 0; i < rowsize; i++ {
			if _, ok := nonemptyrows[i]; !ok {
				emptyrows += 1
			}
			emptyrowscum[i] = emptyrows
		}
	}
	emptycolscum := make([]int, colsize)
	{
		emptycols := 0
		for j := 0; j < colsize; j++ {
			if _, ok := nonemptycols[j]; !ok {
				emptycols += 1
			}
			emptycolscum[j] = emptycols
		}
	}

	const gapsize = 1000000
	extradist := func(p1, p2 point) int {
		rmin := min(p1.r, p2.r)
		rmax := max(p1.r, p2.r)
		cmin := min(p1.c, p2.c)
		cmax := max(p1.c, p2.c)

		rowgaps := emptyrowscum[rmax] - emptyrowscum[rmin]
		colgaps := emptycolscum[cmax] - emptycolscum[cmin]
		gaps := rowgaps + colgaps
		return gaps * gapsize
	}

	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			di := dist(g1, g2)
			x := extradist(g1, g2)
			d := di + x
			fmt.Printf("%v->%v: %d (%d+%d)\n", g1, g2, d, di, x)
			sum += d
		}
	}
	fmt.Println(sum)

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

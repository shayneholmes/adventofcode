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

	// expand the universe
	galaxies := map[point]bool{}
	emptyrs := map[int]bool{}
	emptycs := map[int]bool{}
	for i := 0; i < len(space); i++ {
		mightbeempty := true
		for j := 0; j < len(space); j++ {
			if space[i][j] == '#' {
				galaxies[point{i, j}] = true
				mightbeempty = false
			}
		}
		if mightbeempty {
			emptyrs[i] = true
		}
	}

	extradist := func(p1, p2 point) int {
		extras := 0
		for r := min(p1.r, p2.r); r < max(p1.r, p2.r); r++ {
			if _, ok := emptyrs[r]; ok {
				extras += 1
			}
		}
		for c := min(p1.c, p2.c); c < max(p1.c, p2.c); c++ {
			if _, ok := emptycs[c]; ok {
				extras += 1
			}
		}
		return extras
	}

	for j := 0; j < len(space); j++ {
		mightbeempty := true
		for i := 0; i < len(space); i++ {
			if space[i][j] == '#' {
				mightbeempty = false
				break
			}
		}
		if mightbeempty {
			emptycs[j] = true
		}
	}

	sum := 0
	for g1 := range galaxies {
		for g2 := range galaxies {
			d := dist(g1, g2) + extradist(g1, g2)
			fmt.Printf("%v->%v: %d\n", g1, g2, d)
			sum += d
		}
	}
	fmt.Println(sum / 2)

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

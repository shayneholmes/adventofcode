package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type nstruct = struct {
	left  string
	right string
}

type state = struct {
	instructionpos int
	node           string
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()
	nodes := map[string]nstruct{}
	startingNodes := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`[A-Z]{3}`)
		node := re.FindAllString(line, -1)
		nodes[node[0]] = nstruct{node[1], node[2]}
		name := node[0]
		if name[2] == 'A' {
			startingNodes = append(startingNodes, name)
		}
	}

	for _, start := range startingNodes {
		n := start
		visited := map[state]int{}
		zs := []int{}
		steps := 0
		times := 0
	out:
		for {
			for ipos, i := range instructions {
				if _, ok := visited[state{ipos, n}]; ok {
					fmt.Printf("loop from %v at %v (start: %v, period: %v)\n", start, n, visited[state{ipos, n}], steps-visited[state{ipos, n}])
					fmt.Printf("zs: %v\n", zs)
					break out
				}
				visited[state{ipos, n}] = steps
				if n[2] == 'Z' {
					fmt.Printf("got from %v to %v at %d\n", start, n, steps)
					zs = append(zs, steps)
					times += 1
					if times > 3 {
						break out
					}
				}
				if i == 'L' {
					n = nodes[n].left
				} else {
					n = nodes[n].right
				}
				steps += 1
			}
		}
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

package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type nstruct = struct {
	left  string
	right string
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
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`[A-Z]{3}`)
		node := re.FindAllString(line, -1)
		nodes[node[0]] = nstruct{node[1], node[2]}
	}

	n := "AAA"
	visited := map[string]int{}
	steps := 0
	for {
		for _, i := range instructions {
			if n == "ZZZ" {
				log.Fatalf("Done! %v steps\n", steps)

			}
			// fmt.Printf("visiting %v\n", n)
			if visited[n] > 0 {
				// fmt.Printf("found a loop; visited %v after %v steps and again after %v\n", n, visited[n], steps)
			}
			visited[n] = steps
			if i == 'L' {
				n = nodes[n].left
			} else {
				n = nodes[n].right
			}
			steps += 1
		}
	}
	// fmt.Println(steps)

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
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

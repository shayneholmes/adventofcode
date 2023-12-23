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

const (
	flipflop = iota
	conjunction
	broadcaster
)

type module struct {
	type_ int
	dests []string
}

const (
	lo = false
	hi = true
)

type pulse struct {
	hilo   bool
	source string
	dest   string
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	modules := map[string]module{}
	lineregex := regexp.MustCompile(`^(\S+) -> (.*)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := lineregex.FindStringSubmatch(line)[1:]
		label := matches[0]
		dests := strings.Split(matches[1], ", ")
		if label[0] == '%' {
			label = label[1:]
			modules[label] = module{flipflop, dests}
		} else if label[0] == '&' {
			label = label[1:]
			modules[label] = module{conjunction, dests}
		} else if label == "broadcaster" {
			modules[label] = module{broadcaster, dests}
		}
	}
	flipmemory := map[string]bool{}
	conjunctionmemory := map[string]map[string]bool{}

	fmt.Println(modules)
	// init conjunctions
	for label, m := range modules {
		for _, d := range m.dests {
			if modules[d].type_ == conjunction {
				if conjunctionmemory[d] == nil {
					conjunctionmemory[d] = map[string]bool{}
				}
				conjunctionmemory[d][label] = false
			}
		}
	}

	los := 0
	his := 0
	for presses := 0; presses < 1000; presses++ {
		// press the button
		q := []pulse{{lo, "button", "broadcaster"}}
		// wait for all pulses
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			fmt.Printf("%s -> %v -> %s\n", cur.source, cur.hilo, cur.dest)
			if cur.hilo == lo {
				los += 1
			} else {
				his += 1
			}
			m := modules[cur.dest]
			switch m.type_ {
			case broadcaster:
				// send pulse to all dests
				for _, d := range m.dests {
					q = append(q, pulse{cur.hilo, cur.dest, d})
				}
			case flipflop:
				if cur.hilo == hi {
					continue
				} else {
					state := flipmemory[cur.dest]
					flipmemory[cur.dest] = !state
					for _, d := range m.dests {
						q = append(q, pulse{!state, cur.dest, d})
					}
				}
			case conjunction:
				conjunctionmemory[cur.dest][cur.source] = cur.hilo
				allon := true
				for _, on := range conjunctionmemory[cur.dest] {
					if !on {
						allon = false
						break
					}
				}
				for _, d := range m.dests {
					q = append(q, pulse{!allon, cur.dest, d})
				}
			}
		}
	}

	fmt.Println(los * his)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
}

// location in a grid
type loc = []int

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

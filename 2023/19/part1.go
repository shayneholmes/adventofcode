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
	readingRules = iota
	readingParts
)

type (
	part struct {
		x, m, a, s int
	}

	label = string

	rule struct {
		condition    func(p part) bool
		conditionStr string
		label        label
	}
)

func sumpart(p part) int {
	return p.x + p.m + p.a + p.s
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	workflows := map[label][]rule{}
	sum := 0
	mode := readingRules
	for scanner.Scan() {
		line := scanner.Text()
		ruleregex := regexp.MustCompile(`(\w+)\{([^}]+)}`)
		if line == "" {
			mode = readingParts
			continue
		}
		if mode == readingRules {
			matches := ruleregex.FindStringSubmatch(line)[1:]
			label := matches[0]
			rules := strings.Split(matches[1], ",")
			workflows[label] = []rule{}
			for _, r := range rules {
				target := ""
				condition := func(p part) bool { return true }
				components := strings.Split(r, ":")
				if len(components) == 1 {
					target = r
				} else {
					target = components[1]
					conditionStr := components[0]
					val := conditionStr[0]
					cmp := conditionStr[1]
					rand := atoi(conditionStr[2:])
					getter := func(p part) int {
						log.Fatalf("undefined getter!")
						return 0
					}
					switch val {
					case 'x':
						getter = func(p part) int {
							return p.x
						}
					case 'm':
						getter = func(p part) int {
							return p.m
						}
					case 'a':
						getter = func(p part) int {
							return p.a
						}
					case 's':
						getter = func(p part) int {
							return p.s
						}
					default:
						log.Fatalf("getter not handled: %q", val)
					}
					switch cmp {
					case '<':
						condition = func(p part) bool {
							val := getter(p)
							return val < rand
						}
					case '>':
						condition = func(p part) bool {
							val := getter(p)
							return val > rand
						}
					default:
						log.Fatalf("op not handled: %q", cmp)
					}
				}
				workflows[label] = append(workflows[label],
					rule{condition, r, target},
				)
			}
		} else {
			c := ints(line)
			p := part{
				c[0],
				c[1],
				c[2],
				c[3],
			}
			workflow := "in"
		workflowstart:
			for {
				if workflow == "R" {
					break
				}
				if workflow == "A" {
					sum += sumpart(p)
					break
				}
				// fmt.Printf("evaluating workflow %v: %v\n", workflow, workflows[workflow])
				for _, r := range workflows[workflow] {
					// fmt.Printf("evaluating %v in %v\n", p, r.conditionStr)
					if r.condition(p) {
						workflow = r.label
						continue workflowstart
					}
				}
				log.Fatalf("rule not followed!")
			}
		}
	}
	fmt.Println(sum)

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

// Wrong: 465261

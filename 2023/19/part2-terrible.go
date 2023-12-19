package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
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
	mode := readingRules
	// boundaries are the lower ends of ranges.
	boundariesByField := map[rune][]int{}
	for _, field := range "xmas" {
		boundariesByField[field] = []int{1}
	}
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
					startOfNewRange := rand
					if cmp == '>' {
						startOfNewRange += 1
					}
					boundariesByField[rune(val)] = append(boundariesByField[rune(val)], startOfNewRange)
				}
				workflows[label] = append(workflows[label],
					rule{condition, r, target},
				)
			}
		} else {
			continue
		}
	}

	accepted := func(p part) bool {
		workflow := "in"
	workflowstart:
		for {
			if workflow == "R" {
				return false
			}
			if workflow == "A" {
				return true
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

	for _, field := range "xmas" {
		// deduplicate them
		slices.Sort(boundariesByField[field])
		deduped := []int{0}
		for _, r := range boundariesByField[field] {
			if r != deduped[len(deduped)-1] {
				deduped = append(deduped, r)
			}
		}
		boundariesByField[field] = deduped
	}

	// try each range
	const ratingbound = 4001 // next layer would be this, if we had one
	combos := 0

	sectionSizesByField := map[rune][]int{}
	for _, field := range "xmas" {
		sectionSizesByField[field] = make([]int, len(boundariesByField[field]))
		for i, b := range boundariesByField[field] {
			size := ratingbound - b
			if i < len(boundariesByField[field])-1 {
				size = boundariesByField[field][i+1] - b
			}
			sectionSizesByField[field][i] = size
		}
	}
	for xi, xb := range boundariesByField['x'] {
		fmt.Printf("%d/%d: %d combos\n", xi, len(boundariesByField['x']), combos)
		for mi, mb := range boundariesByField['m'] {
			for ai, ab := range boundariesByField['a'] {
				for si, sb := range boundariesByField['s'] {
					p := part{xb, mb, ab, sb}
					if accepted(p) {
						combos += sectionSizesByField['x'][xi] * sectionSizesByField['m'][mi] * sectionSizesByField['a'][ai] * sectionSizesByField['s'][si]

					}
				}
			}
		}
	}

	fmt.Println(combos)
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

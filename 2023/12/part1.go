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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.Split(line, " ")
		input, pattern := sp[0], sp[1]
		sum += waysToMatch(pattern, input)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func waysToMatch(origpattern, input string) int {
	pattern := regularizePattern(origpattern)
	// fmt.Println(pattern)

	// dynamic programming
	// Each cell [r][c] contains the number of ways to match input up to input[r]
	// with pattern up to pattern[c-1]
	prev := make([]int, len(pattern)+1)
	prev[0] = 1 // One way to match zero units of input with zero units pattern
	for i := 1; i <= len(pattern); i++ {
		if pattern[i-1] == '*' {
			prev[i] = prev[i-1]
		}
	}
	cur := make([]int, len(pattern)+1)
	curIndex := 0
	for ; curIndex < len(input); curIndex++ {
		ch := input[curIndex] // match this unit of input
		cur[0] = 0
		for i := 1; i <= len(pattern); i++ {
			switch pattern[i-1] {
			case '#': // must match a damaged one
				switch ch {
				case '.': // no match
					cur[i] = 0
					continue
				case '#': // Always matches once
					cur[i] = prev[i-1]
					continue
				case '?': // Can match one way
					cur[i] = prev[i-1]
					continue
				default:
					log.Fatalf("unmatched pair %q,%q", pattern[i-1], ch)
				}
			case '.': // must match an undamaged one
				switch ch {
				case '.': // Can match
					cur[i] = prev[i-1]
					continue
				case '#': // no match
					cur[i] = 0
					continue
				case '?': // Can match one way
					cur[i] = prev[i-1]
					continue
				default:
					log.Fatalf("unmatched pair %q,%q", pattern[i-1], ch)
				}
			case '*': // match zero or more undamaged
				// Matching zero means blindly copying cur[i-1], all cases here should
				// have that.
				switch ch {
				case '.': // match, might consume or not
					cur[i] = cur[i-1] + prev[i]
					continue
				case '?': // match, might consume or not
					cur[i] = cur[i-1] + prev[i]
					continue
				case '#': // can't consume this token
					cur[i] = cur[i-1] + 0
					continue
				default:
					log.Fatalf("unmatched pair %q,%q", pattern[i-1], ch)
				}
			default:
				log.Fatalf("unmatched pair %q,%q", pattern[i-1], ch)
			}
		}
		// fmt.Printf("%c: %v\n", ch, cur)
		cur, prev = prev, cur
	}

	fmt.Printf("%v: %v -> %q: %v\n", input, origpattern, pattern, prev[len(pattern)])

	return prev[len(pattern)]
}

func regularizePattern(pat string) string {
	var sb strings.Builder
	sb.WriteString("*")
	runs := ints(pat)
	for r, run := range runs {
		for i := 0; i < run; i++ {
			sb.WriteRune('#')
		}
		if r < len(runs)-1 {
			sb.WriteString(".*")
		}
	}
	sb.WriteString("*")
	return sb.String()
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

// Answers: 7065 is too high
// Answers: 7846 is too high
// Answers: 4060 is too low

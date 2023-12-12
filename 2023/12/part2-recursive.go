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

type parameterSet = struct {
	r int
	c int
}

func waysToMatch(origpattern, originput string) int {
	input := strings.Join([]string{originput, originput, originput, originput, originput}, "?")
	longpattern := strings.Join([]string{origpattern, origpattern, origpattern, origpattern, origpattern}, ",")
	pattern := regularizePattern(longpattern)

	// dynamic programming
	// Each cell [r][c] contains the number of ways to match input up to input[r]
	// with pattern up to pattern[c]
	cached := map[parameterSet]int{}
	var waystomatch func(r, c int) int
	waystomatch = func(r, c int) int {
		params := parameterSet{r, c}
		// check cache
		if cached, ok := cached[params]; ok {
			return cached
		}
		inner := func(r, c int) int {
			// base cases
			if r == -1 && c == -1 {
				return 1 // Seed: 1 way to match nothing with nothing
			}
			if c < 0 {
				return 0
			}
			if r == -1 && c == 0 {
				return 1 // First star always matches
			}
			if r < 0 {
				return 0
			}
			in := input[r]
			pat := pattern[c]
			switch pat {
			case '.':
				switch in {
				case '.':
					return waystomatch(r-1, c-1)
				case '#':
					return 0
				case '?':
					return waystomatch(r-1, c-1)
				}
			case '#':
				switch in {
				case '.':
					return 0
				case '#':
					return waystomatch(r-1, c-1)
				case '?':
					return waystomatch(r-1, c-1)
				}
			case '*':
				// We can match zero tokens, or keep matching as long as we don't have broken ones.
				// In other words, we fork: Either move the pattern forward (match zero
				// tokens) or move the input forward (match a token and keep the
				// pattern the same).
				switch in {
				case '.':
					return waystomatch(r, c-1) + waystomatch(r-1, c)
				case '#':
					// Match zero is the only option
					return waystomatch(r, c-1)
				case '?':
					return waystomatch(r, c-1) + waystomatch(r-1, c)
				}
			}
			log.Fatalf("bad input! %c, %c", pat, in)
			return 0
		}
		res := inner(r, c)
		cached[params] = res
		return res
	}
	return waystomatch(len(input)-1, len(pattern)-1)
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

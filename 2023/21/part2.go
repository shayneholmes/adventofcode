package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	north = loc{-1, 0}
	east  = loc{0, 1}
	west  = loc{0, -1}
	south = loc{1, 0}
)

func add(i, j loc) loc {
	return loc{i.r + j.r, i.c + j.c}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	getstart := func() tloc {
		for r := range grid {
			for c := range grid[0] {
				if grid[r][c] == 'S' {
					return tloc{loc{r, c}, 0, 0}
				}
			}
		}
		return tloc{loc{0, 0}, 0, 0}
	}

	maxr := len(grid)
	maxc := len(grid[0])

	tadd := func(i tloc, j loc) tloc {
		nu := add(i.loc, j)
		dtile := loc{0, 0}
		if nu.r >= maxr {
			nu.r = 0
			dtile.r = 1
		}
		if nu.r < 0 {
			nu.r += maxr
			dtile.r = -1
		}
		if nu.c >= maxc {
			nu.c = 0
			dtile.c = 1
		}
		if nu.c < 0 {
			nu.c += maxc
			dtile.c = -1
		}
		return tloc{nu, i.tr + dtile.r, i.tc + dtile.c}
	}

	// locs is the frontier: new places from which we may go to new places
	loccount := []int{1, 0}
	locs := map[tloc]bool{getstart(): true}
	lastlocs := map[tloc]bool{}

	// grows quadratically over sufficient periods
	lastround := 10000
	// lastround := 26501365
	// lastround := 26501365

	for round := 1; round <= lastround; round++ {
		// if round > 65 {
		// 	stategrid := make([][]int, maxr)
		// 	for r := range stategrid {
		// 		stategrid[r] = make([]int, maxc)
		// 	}
		// 	for l := range locs {
		// 		stategrid[l.r][l.c] += 1
		// 	}
		// 	for r := 0; r < maxr; r++ {
		// 		for c := 0; c < maxc; c++ {
		// 			s := stategrid[r][c]
		// 			if s > 0 {
		// 				fmt.Printf("%1d", s)
		// 			} else {
		// 				fmt.Printf("%c", grid[r][c])
		// 			}
		// 		}
		// 		fmt.Printf("\n")
		// 	}
		// 	fmt.Printf("\n")
		// }

		// polynomial: Ax^2 + bx + c = 0
		// x = 0: 3538
		// x = 1: 93156
		// x = 2: 634552

		// x -> turn # (65 + x * 131 * 2)

		// search is where x = 101150

		// c = 3538
		// A + b + c = 93156
		// 4A + 2b + c = 303494
		// 9A + 3b + c = 634552

		// A + b + c = 93156
		// A + b + 3538 = 93156
		// A = 89618 - b

		// 4A + 2b + c = 303494
		// 4(89618 - b) + 2b + 3538 = 303494
		// 358472 - 4b + 2b + 3538 = 303494 out!
		// -2b = 303494 - 358472 - 3538 = -58516
		// b = 29258
		// so A = 89618 - 29258 = 60360

		// Check: at x = 4,
		// 16*60360+4*29258+3538 = 1086330
		// it checks out!
		// 16 * 118876 + 4 * -29258 + 3538 = 1788522

		// guesses:
		// 617565585550238 (too low)
		// 617565585550239 (too low)
		// 617577796467856 (too high, we're in the right neighborhood)

		// 617565692567199 // had spit out numbers _before_ taking the steps!
		// prediction at x=20 (turn 5305) : 24732698  -> it checks out!

		candidatelocs := map[tloc]bool{}
		for l := range locs {
			for _, dir := range []loc{north, south, west, east} {
				nu := tadd(l, dir)
				if grid[nu.r][nu.c] == '#' {
					// it's a rock.
					continue
				}
				if candidatelocs[nu] {
					// we can get there from another place this turn.
					continue
				}
				if lastlocs[nu] {
					// we were just here last turn, don't go back.
					continue
				}
				candidatelocs[nu] = true
				loccount[round%2] += 1
			}
		}
		lastlocs = locs
		locs = candidatelocs
		if round%(maxr*2) == 65 {
			estimates[
			fmt.Printf("Round %d: len(locs) = %d, loccount = %d\n", round, len(locs), loccount)
		}
	}

	fmt.Println(loccount[lastround%2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
}

// location in a grid
type loc struct {
	r, c int
}

type tloc struct {
	loc
	tr, tc int
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

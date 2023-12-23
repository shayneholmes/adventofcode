package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type brick struct {
	// created so that x1 < x2, y1 < y2, z1 < z2
	x1, y1, z1, x2, y2, z2 int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bricks := []brick{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		is := ints(line)
		if len(is) != 6 {
			log.Fatalf("bad line! %v -> %v\n", line, is)
		}
		bricks = append(bricks, brick{
			min(is[0], is[3]),
			min(is[1], is[4]),
			min(is[2], is[5]),
			max(is[0], is[3]),
			max(is[1], is[4]),
			max(is[2], is[5]),
		})
	}

	// sort bricks from bottom to top
	slices.SortFunc(bricks, func(a, b brick) int {
		return cmp.Compare(a.z1, b.z1)
	})

	highestpoint := map[loc]int{}
	highestbrick := map[loc]int{}  // index into bricks
	supportedby := map[int][]int{} // supprotedby[i] is a list of indexes of bricks that are supported by brick i
	supports := map[int][]int{}    // supports[i] is a list of indexes of bricks that support brick i
	// drop bricks onto the ground
	for bi, b := range bricks {
		// compute how far this brick can fall
		zmax := 0
		for x := b.x1; x <= b.x2; x++ {
			for y := b.y1; y <= b.y2; y++ {
				l := loc{x, y}
				if z, ok := highestpoint[l]; ok && z > zmax {
					zmax = z
				}
			}
		}

		// drop the brick to zmax + 1
		dz := b.z2 - b.z1 + 1
		b.z1 = zmax + 1
		b.z2 = zmax + dz
		// update the height map
		lastsupported := -1
		for x := b.x1; x <= b.x2; x++ {
			for y := b.y1; y <= b.y2; y++ {
				l := loc{x, y}
				if zmax > 0 && highestpoint[l] == zmax && lastsupported != highestbrick[l] {
					support := highestbrick[l] // `support` supports `bi`
					supportedby[support] = append(supportedby[support], bi)
					supports[bi] = append(supports[bi], support)
					lastsupported = support
				}
				highestpoint[l] = zmax + dz
				highestbrick[l] = bi
			}
		}
	}

	numsafe := 0
	for bi := range bricks {
		safe := true
		for _, s := range supportedby[bi] {
			if len(supports[s]) == 1 {
				safe = false
			}
		}
		if safe {
			numsafe += 1
		}
	}
	fmt.Println(numsafe)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
}

// location in a grid
type loc = struct {
	x, y int
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

// guessses: 1470 too high

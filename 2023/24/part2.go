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

type stone = struct {
	pos, v    point
	slope     float64
	intercept float64
}

func futureintersection(s1, s2 stone) point {
	// ignore time
	// find a point at which [x,y] satisfies y = mx+b for both things
	// m1x + b1 = m2x + b2
	// (m1 - m2)x = b2 - b1
	// x = (b2 - b1)/(m1 - m2)
	x := (s2.intercept - s1.intercept) / (s1.slope - s2.slope)
	// y = m1x + b1
	y1 := (s1.slope)*x + s1.intercept
	y2 := (s1.slope)*x + s1.intercept
	if y1 != y2 {
		log.Fatalf("%f != %f for %v, %v\n", y1, y2, s1, s2)
	}
	// figure out if they are in the past
	// x = s1.x + t * s1.dx
	// t = (x - s1.x) / s1.dx
	// so t < 0 if the two terms are different signs
	t1 := (x - s1.pos[0]) / (s1.v[0])
	t2 := (x - s2.pos[0]) / (s2.v[0])
	// fmt.Printf("%v, %v -> %v at time %f, %f\n", s1, s2, point{x, y1}, t1, t2)
	if t1 < 0 || t2 < 0 {
		return point{math.Inf(0), math.Inf(0)}
	}
	return point{x, y1}
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stones := []stone{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		is := floats(line)
		s := stone{
			pos: point{is[0], is[1], is[2]},
			v:   point{is[3], is[4], is[5]},
		}
		// the path is traced by the function y = s1.x
		// [s1.x, s1.y] and [s1.x + dx, s1.y + dy]
		// compute the formula
		// slope is (s1.y + dy - s1.y) / (s1.x + dx - s1.x) = dy / dx
		// intercept moves left s1.x, so is s1.y - (dy / dx * s1.x)
		// y = mx + b @ s1.x, s1.y
		// s1.y = (dy/dx)*s1.x + b
		// b = s1.y - (dy / dx * s1.x)
		// for some position (x, y), s1.x +
		s.slope = s.v[1] / s.v[0]
		s.intercept = s.pos[1] - s.slope*s.pos[0]
		stones = append(stones, s)
		// fmt.Println(s)
	}

	intersections := 0
	for i, s1 := range stones {
		for _, s2 := range stones[i+1:] {
			// fmt.Printf("%v, %v\n", s1, s2)
			intersection := futureintersection(s1, s2)
			if intersection[0] >= xmin && intersection[0] <= xmax &&
				intersection[1] >= ymin && intersection[1] <= ymax {
				intersections += 1
			}
		}
	}
	fmt.Println(intersections)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

const (
	// xmin = 7
	// xmax = 27
	xmin = 200000000000000
	xmax = 400000000000000
	ymin = xmin
	ymax = xmax
)

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
}

var (
	north = loc{-1, 0}
	east  = loc{0, 1}
	west  = loc{0, -1}
	south = loc{1, 0}
)

// location in a grid
type loc = struct {
	r int
	c int
}

func add(i, j loc) loc {
	return loc{i.r + j.r, i.c + j.c}
}

var dirs = []loc{north, south, east, west}

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

// solve 9 equations
// 275325627102914, 177556324137106, 279758114394131 @ 249, 405, -531
// 284428334220238, 231958436807561, 189800593445547 @ 237, 140, -111
// 208260774362545, 354915461185166, 308973039318009 @ 128, -159, -65
//
// xs+t1*dxs=275325627102914+t1*249
// ys+t1*dys=177556324137106+t1*405
// zs+t1*dzs=279758114394131+t1*(-531)
// xs+t2*dxs=284428334220238+t2*237
// ys+t2*dys=231958436807561+t2*140
// zs+t2*dzs=189800593445547+t2*(-111)
// xs+t3*dxs=208260774362545+t3*128
// ys+t3*dys=354915461185166+t3*(-159)
// zs+t3*dzs=308973039318009+t3*(-65)
//
// Not linear, so we have to do something else...
// xs+t1*dxs=275325627102914+t1*249
// xs+t1*dxs=275325627102914+t1*249

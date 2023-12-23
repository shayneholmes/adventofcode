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

type pathstate struct {
	loc
	dist int
}
type edge struct {
	start, end loc
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

	rmax := len(grid)
	cmax := len(grid[0])

	start := loc{0, 1}
	end := loc{140, 139}
	visited := map[loc]bool{}

	getneighbors := func(spot loc) (res []loc) {
		dirs := []loc{north, south, east, west}
		for _, dir := range dirs {
			nu := add(spot, dir)
			if nu.r < 0 || nu.r >= rmax {
				continue
			}
			if nu.c < 0 || nu.c >= cmax {
				continue
			}
			if grid[nu.r][nu.c] == '#' {
				continue
			}
			if visited[nu] {
				continue
			}
			res = append(res, nu)
		}
		return res
	}

	// construct a graph of intersections
	nodes := map[loc]bool{start: true, end: true}
	// explore the graph
	var exploregraph func(loc loc)
	exploregraph = func(loc loc) {
		visited[loc] = true
		neighbors := getneighbors(loc)
		if len(neighbors) > 1 {
			nodes[loc] = true
		}
		for _, n := range neighbors {
			exploregraph(n)
		}
	}
	exploregraph(start)

	dists := map[loc]map[loc]int{}
	fmt.Printf("nodes: %v\n", nodes)

	// compute dists to adjacent nodes
	for n1 := range nodes {
		visited = map[loc]bool{}
		dists[n1] = map[loc]int{}
		q := []pathstate{{n1, 0}}
		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			visited[n.loc] = true
			if nodes[n.loc] && n.dist > 0 {
				// adjacent node
				dists[n1][n.loc] = n.dist
			} else {
				neighbors := getneighbors(n.loc)
				for _, neighbor := range neighbors {
					q = append(q, pathstate{neighbor, n.dist + 1})
				}
			}
		}
	}

	fmt.Printf("%d nodes\n", len(nodes))
	fmt.Println(dists)

	// compute longest path to end node
	visited = map[loc]bool{}
	var longestpath func(loc loc) int
	longestpath = func(from loc) int {
		// fmt.Printf("%s%v\n", strings.Repeat(" ", len(visited)), loc)
		if from == end {
			return 0
		}
		visited[from] = true
		res := -1
		for n, d := range dists[from] {
			// fmt.Printf("Evaluating %v->%v (%d)\n", loc, n, d)
			if !visited[n] {
				r := longestpath(n)
				if r > -1 {
					one := r + d
					if one > res {
						res = one
					}
				}
			}
		}
		delete(visited, from)
		return res
	}

	fmt.Println(longestpath(start))

}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
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

// 6794 is too high

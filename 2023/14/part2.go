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

	scanner := bufio.NewScanner(file)
	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
		fmt.Println(line)
	}
	positions := map[string]int{}
	foundCycle := false
	for rotations := 1_000_000_000 - 1; rotations >= 0; rotations-- {
		if !foundCycle {
			var sb strings.Builder
			for _, li := range grid {
				sb.Write(li)
			}
			key := sb.String()
			if pos, ok := positions[key]; ok {
				foundCycle = true
				fmt.Printf("Cycle! We saw this at move %d and again at move %d\n", pos, rotations)
				cycleLength := abs(rotations - pos)
				for ; rotations > cycleLength; rotations -= cycleLength {
				}
			}
			positions[key] = rotations
		}
		// North
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] == 'O' {
					// roll the rock north
					nr := r - 1
					for ; nr >= 0 && grid[nr][c] == '.'; nr-- {
					}
					nr++ // move south from the obstacle
					if nr != r {
						grid[nr][c], grid[r][c] = grid[r][c], grid[nr][c]
					}
				}
			}
		}
		// West
		for c := 0; c < len(grid[0]); c++ {
			for r := 0; r < len(grid); r++ {
				if grid[r][c] == 'O' {
					// roll the rock west
					nc := c - 1
					for ; nc >= 0 && grid[r][nc] == '.'; nc-- {
					}
					nc++ // move east from the obstacle
					if nc != c {
						grid[r][nc], grid[r][c] = grid[r][c], grid[r][nc]
					}
				}
			}
		}
		// South
		for r := len(grid) - 1; r >= 0; r-- {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] == 'O' {
					// roll the rock south
					nr := r + 1
					for ; nr < len(grid) && grid[nr][c] == '.'; nr++ {
					}
					nr-- // move south from the obstacle
					if nr != r {
						grid[nr][c], grid[r][c] = grid[r][c], grid[nr][c]
					}
				}
			}
		}
		// East
		for c := len(grid[0]) - 1; c >= 0; c-- {
			for r := len(grid) - 1; r >= 0; r-- {
				if grid[r][c] == 'O' {
					// roll the rock east
					nc := c + 1
					for ; nc < len(grid) && grid[r][nc] == '.'; nc++ {
					}
					nc-- // move west from the obstacle
					if nc != c {
						grid[r][nc], grid[r][c] = grid[r][c], grid[r][nc]
					}
				}
			}
		}
	}
	weight := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 'O' {
				weight += len(grid) - r
			}
		}
	}

	for _, line := range grid {
		fmt.Println(string(line))
	}
	fmt.Println(weight)

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
	r int
	c int
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

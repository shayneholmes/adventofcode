package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// location in a grid
type loc = struct {
	r int
	c int
}

type instruction struct {
	dir loc
	mag int
}

var (
	north = loc{-1, 0}
	east  = loc{0, 1}
	west  = loc{0, -1}
	south = loc{1, 0}
)

func add(i, j loc) loc {
	return loc{i.r + j.r, i.c + j.c}
}

func mult(i loc, m int) loc {
	return loc{i.r * m, i.c * m}
}

func parsedir(s byte) loc {
	switch s {
	case '0':
		return east
	case '1':
		return south
	case '2':
		return west
	case '3':
		return north
	}
	log.Fatalf("couldn't parse %v", s)
	return north
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := make([]instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		hex := strings.TrimPrefix(splits[2], "(#")
		mag, _ := strconv.ParseInt(hex[0:5], 16, 64)
		dir := parsedir(hex[5])
		instructions = append(instructions, instruction{dir, int(mag)})
	}

	area := 0
	perimeter := 0
	pos := loc{0, 0}
	for _, inst := range instructions {
		perimeter += inst.mag
		nu := add(pos, mult(inst.dir, inst.mag))
		area += (nu.c - pos.c) * (nu.r + pos.r) / 2
		pos = nu
	}
	area = abs(area)

	fmt.Printf("area: %v\n", area+(perimeter)/2+1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
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

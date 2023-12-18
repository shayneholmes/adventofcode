package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	perimeter := 0
	pos := loc{0, 0}
	locs := make([]loc, 0, 700)
	locs = append(locs, pos)
	rbounds := map[int]bool{0: true}
	cbounds := map[int]bool{0: true}
	for _, inst := range instructions {
		pos = add(pos, mult(inst.dir, inst.mag))
		rbounds[pos.r] = true
		cbounds[pos.c] = true
		perimeter += inst.mag
		locs = append(locs, pos)
	}

	rboundlist := []int{}
	for r := range rbounds {
		rboundlist = append(rboundlist, r)
	}
	slices.Sort(rboundlist)

	cboundlist := []int{}
	for c := range cbounds {
		cboundlist = append(cboundlist, c)
	}
	slices.Sort(cboundlist)

	inside := func(test loc) bool {
		intersections := 0
		last := locs[0]
		for _, l := range locs[1:] {
			horizontal := last.r == l.r
			above := last.r <= test.r

			left := min(l.c, last.c)
			right := max(l.c, last.c)
			leftandright := (left <= test.c) && (right > test.c)
			intersects := horizontal && above && leftandright
			if intersects {
				intersections += 1
			}
			last = l
		}
		return intersections%2 == 1
	}

	area := 0
	lastr := rboundlist[0]
	for _, rb := range rboundlist[1:] {
		lastc := cboundlist[0]
		for _, cb := range cboundlist[1:] {
			test := loc{lastr, lastc}
			if inside(test) {
				// compute area inside this thing, including left and top walls
				area += (cb - lastc) * (rb - lastr)
			}
			lastc = cb
		}
		lastr = rb
	}
	fmt.Printf("area: %v\n", area+(perimeter)/2+1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

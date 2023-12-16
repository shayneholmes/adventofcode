package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	north = loc{-1, 0}
	east  = loc{0, 1}
	west  = loc{0, -1}
	south = loc{1, 0}
)

type stage = struct {
	loc loc
	dir loc
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tovisit := []stage{{loc{0, 0}, east}}
	stagesseen := map[stage]bool{}
	visited := map[loc]bool{}
	for len(tovisit) > 0 {
		currentstage := tovisit[len(tovisit)-1]
		tovisit = tovisit[:len(tovisit)-1]
		if stagesseen[currentstage] {
			continue
		}
		spot := currentstage.loc
		dir := currentstage.dir
		if spot.r < 0 || spot.r >= len(grid) || spot.c < 0 || spot.c >= len(grid[0]) {
			// out of bounds
			continue
		}
		visited[spot] = true
		stagesseen[currentstage] = true

		godir := func(dir loc) {
			tovisit = append(tovisit, stage{add(spot, dir), dir})
		}

		switch grid[spot.r][spot.c] {
		case '.':
			// If the beam encounters empty space (.), it continues in the same
			// direction.
			godir(dir)
			continue
		case '/':
			// If the beam encounters a mirror (/ or \), the beam is reflected 90
			// degrees depending on the angle of the mirror. For instance, a
			// rightward-moving beam that encounters a / mirror would continue upward
			// in the mirror's column, while a | rightward-moving beam that encounters
			// a \ mirror would continue downward from the mirror's column.
			switch dir {
			case north:
				dir = east
			case south:
				dir = west
			case east:
				dir = north
			case west:
				dir = south
			}
			godir(dir)
			continue
		case '\\':
			switch dir {
			case north:
				dir = west
			case south:
				dir = east
			case east:
				dir = south
			case west:
				dir = north
			}
			godir(dir)
			continue

		case '|':
			switch dir {
			// If the beam encounters the pointy end of a splitter (| or -), the beam passes through the splitter as if the splitter were empty space. For instance, a rightward-moving beam that encounters a - splitter would continue in the same direction.
			case north:
				fallthrough
			case south:
				godir(dir)
			// If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.
			case east:
				fallthrough
			case west:
				godir(north)
				godir(south)
			}
			continue
		case '-':
			switch dir {
			// If the beam encounters the pointy end of a splitter (| or -), the beam passes through the splitter as if the splitter were empty space. For instance, a rightward-moving beam that encounters a - splitter would continue in the same direction.
			case east:
				fallthrough
			case west:
				godir(dir)
			// If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.
			case north:
				fallthrough
			case south:
				godir(east)
				godir(west)
			}
			continue
		}
	}
	fmt.Println(len(visited))
}

// location in a grid
type loc = struct {
	r int
	c int
}

func add(i, j loc) loc {
	return loc{i.r + j.r, i.c + j.c}
}

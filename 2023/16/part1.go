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
		// for i := range grid {
		// 	for j := range grid[0] {
		// 		if visited[loc{i, j}] {
		// 			fmt.Printf("#")
		// 		} else {
		// 			fmt.Printf("%c", grid[i][j])
		// 		}
		// 	}
		// 	fmt.Printf("\n")
		// }
		// fmt.Printf("\n")
		currentstage := tovisit[0]
		tovisit = tovisit[1:]
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

		switch grid[spot.r][spot.c] {
		case '.':
			// If the beam encounters empty space (.), it continues in the same
			// direction.
			tovisit = append(tovisit, stage{loc{spot.r + dir.r, spot.c + dir.c}, dir})
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
			tovisit = append(tovisit, stage{loc{spot.r + dir.r, spot.c + dir.c}, dir})
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
			tovisit = append(tovisit, stage{loc{spot.r + dir.r, spot.c + dir.c}, dir})
			continue

		case '|':
			switch dir {
			// If the beam encounters the pointy end of a splitter (| or -), the beam passes through the splitter as if the splitter were empty space. For instance, a rightward-moving beam that encounters a - splitter would continue in the same direction.
			case north:
				fallthrough
			case south:
				tovisit = append(tovisit, stage{loc{spot.r + dir.r, spot.c + dir.c}, dir})
			// If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.
			case east:
				fallthrough
			case west:
				tovisit = append(tovisit, stage{loc{spot.r + north.r, spot.c + north.c}, north}, stage{loc{spot.r + south.r, spot.c + south.c}, south})
			}
			continue
		case '-':
			switch dir {
			// If the beam encounters the pointy end of a splitter (| or -), the beam passes through the splitter as if the splitter were empty space. For instance, a rightward-moving beam that encounters a - splitter would continue in the same direction.
			case east:
				fallthrough
			case west:
				tovisit = append(tovisit, stage{loc{spot.r + dir.r, spot.c + dir.c}, dir})
			// If the beam encounters the flat side of a splitter (| or -), the beam is split into two beams going in each of the two directions the splitter's pointy ends are pointing. For instance, a rightward-moving beam that encounters a | splitter would split into two beams: one that continues upward from the splitter's column and one that continues downward from the splitter's column.
			case north:
				fallthrough
			case south:
				tovisit = append(tovisit, stage{loc{spot.r + east.r, spot.c + east.c}, east}, stage{loc{spot.r + west.r, spot.c + west.c}, west})
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

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
	loc
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

	countenergizedstartingat := func(seed stage) int {
		tovisit := []stage{seed}
		stagesseen := map[stage]bool{}
		visited := map[loc]bool{}
		for len(tovisit) > 0 {
			spot := tovisit[len(tovisit)-1]
			tovisit = tovisit[:len(tovisit)-1]
			if stagesseen[spot] {
				continue
			}
			godir := func(dir loc) {
				tovisit = append(tovisit, stage{add(spot.loc, dir), dir})
			}

			dir := spot.dir
			if spot.r < 0 || spot.r >= len(grid) || spot.c < 0 || spot.c >= len(grid[0]) {
				// out of bounds
				continue
			}
			stagesseen[spot] = true
			visited[spot.loc] = true

			switch grid[spot.r][spot.c] {
			case '.':
				// If the beam encounters empty space (.), it continues in the same
				// direction.
				godir(dir)
			case '/':
				// If the beam encounters a mirror (/ or \), the beam is reflected 90
				// degrees depending on the angle of the mirror. For instance, a
				// rightward-moving beam that encounters a / mirror would continue upward
				// in the mirror's column, while a | rightward-moving beam that encounters
				// a \ mirror would continue downward from the mirror's column.
				switch dir {
				case north:
					godir(east)
				case south:
					godir(west)
				case east:
					godir(north)
				case west:
					godir(south)
				}
			case '\\':
				switch dir {
				case north:
					godir(west)
				case south:
					godir(east)
				case east:
					godir(south)
				case west:
					godir(north)
				}
			case '|':
				switch dir {
				// If the beam encounters the pointy end of a splitter (| or -), the
				// beam passes through the splitter as if the splitter were empty
				// space. For instance, a rightward-moving beam that encounters a -
				// splitter would continue in the same direction.
				case north:
					fallthrough
				case south:
					godir(dir)
				// If the beam encounters the flat side of a splitter (| or -), the
				// beam is split into two beams going in each of the two directions the
				// splitter's pointy ends are pointing. For instance, a
				// rightward-moving beam that encounters a | splitter would split into
				// two beams: one that continues upward from the splitter's column and
				// one that continues downward from the splitter's column.
				case east:
					fallthrough
				case west:
					godir(north)
					godir(south)
				}
			case '-':
				switch dir {
				// If the beam encounters the pointy end of a splitter (| or -), the
				// beam passes through the splitter as if the splitter were empty
				// space. For instance, a rightward-moving beam that encounters a -
				// splitter would continue in the same direction.
				case east:
					fallthrough
				case west:
					godir(dir)
				// If the beam encounters the flat side of a splitter (| or -), the
				// beam is split into two beams going in each of the two directions the
				// splitter's pointy ends are pointing. For instance, a
				// rightward-moving beam that encounters a | splitter would split into
				// two beams: one that continues upward from the splitter's column and
				// one that continues downward from the splitter's column.
				case north:
					fallthrough
				case south:
					godir(east)
					godir(west)
				}
			}
		}
		return len(visited)
	}

	seeds := func() []stage {
		stages := []stage{}
		rmax := len(grid) - 1
		cmax := len(grid[0]) - 1
		for i := range grid {
			stages = append(stages,
				stage{loc{i, 0}, east},
				stage{loc{i, cmax}, west},
			)
		}
		for j := range grid[0] {
			stages = append(stages,
				stage{loc{0, j}, south},
				stage{loc{rmax, j}, north},
			)
		}
		return stages
	}

	max := 0
	for _, seed := range seeds() {
		en := countenergizedstartingat(seed)
		if en > max {
			max = en
		}
	}
	fmt.Println(max)
}

// location in a grid
type loc = struct {
	r int
	c int
}

func add(i, j loc) loc {
	return loc{i.r + j.r, i.c + j.c}
}

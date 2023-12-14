package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	north := loc{-1, 0}
	west := loc{0, -1}
	south := loc{1, 0}
	east := loc{0, 1}

	scanner := bufio.NewScanner(file)
	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	rollRocks := func(diff loc) {
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[0]); c++ {
				if grid[r][c] == 'O' { // it's a rock.
					// roll the rock until we hit a barrier
					// if it's a rock, move "through" it, and we'll backtrack in a second
					rr, rc := r, c
					for {
						nr, nc := rr+diff[0], rc+diff[1]
						if nr < 0 || nr >= len(grid) {
							break
						}
						if nc < 0 || nc >= len(grid[0]) {
							break
						}
						if grid[nr][nc] == '#' {
							break
						}
						rr, rc = nr, nc
					}
					// We've moved as far as we can go, but maybe we moved "through" some
					// rocks. Backtrack through them until we find a suitable landing
					// spot.
					for {
						if rr == r && rc == c {
							// we're back where we started, we can definitely land here.
							break
						}
						if grid[rr][rc] == '.' {
							// this spot is empty, it will do nicely.
							break
						}
						nr, nc := rr-diff[0], rc-diff[1]
						if nr < 0 || nr >= len(grid) {
							log.Fatalf("backtracking landed us off the grid! %v, %v", nr, nc)
						}
						if nc < 0 || nc >= len(grid[0]) {
							log.Fatalf("backtracking landed us off the grid! %v, %v", nr, nc)
						}
						// backtrack further
						rr, rc = nr, nc
					}
					if !(rr == r && rc == c) {
						grid[rr][rc], grid[r][c] = grid[r][c], grid[rr][rc]
					}
				}
			}
		}
	}

	keyForGrid := func() string {
		var keybuilder strings.Builder
		for _, li := range grid {
			keybuilder.Write(li)
		}
		return keybuilder.String()
	}

	rotations := 1_000_000_000
	firstTurnSeen := map[string]int{}
	foundCycle := false
	for rot := 0; rot < rotations; rot++ {
		rollRocks(north)
		rollRocks(west)
		rollRocks(south)
		rollRocks(east)

		if !foundCycle {
			key := keyForGrid()
			if firstSeen, ok := firstTurnSeen[key]; !ok {
				firstTurnSeen[key] = rot
			} else {
				foundCycle = true
				cycleLength := rot - firstSeen
				fmt.Printf("Cycle! We saw this at move %d and again at move %d -> length %d\n", firstSeen, rot, cycleLength)
				// fast-forward
				rot += (rotations - rot) / cycleLength * cycleLength
			}
		}
	}

	// sum weights
	weight := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 'O' {
				weight += len(grid) - r
			}
		}
	}

	fmt.Println(weight)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// location in a grid
type loc = []int

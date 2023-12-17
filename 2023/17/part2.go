package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

type polarity = int

const (
	eastwest   = polarity(0)
	northsouth = polarity(1)
)

func polarityfromdir(dir loc) polarity {
	switch dir {
	case north:
		fallthrough
	case south:
		return northsouth
	case east:
		fallthrough
	case west:
		return eastwest
	default:
		return eastwest
	}
}

type params = struct {
	loc
	dir polarity
}

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
	}

	goal := loc{len(grid) - 1, len(grid[0]) - 1}
	start := loc{0, 0}

	seeds := map[params]int{
		{start, northsouth}: 0,
		{start, eastwest}:   0,
	}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(seeds))
	i := 0
	for value, priority := range seeds {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	dists := map[params]int{}

	for pq.Len() > 0 {
		cur := heap.Pop(&pq).(*Item)
		val := cur.value
		cost := cur.priority
		if _, ok := dists[val]; ok {
			// already have this one
			continue
		} else {
			// record this one
			dists[val] = cost
		}

		for _, nudir := range []loc{south, east, north, west} {
			if polarityfromdir(nudir) == val.dir {
				// can't double back or keep going
				continue
			}
			nucost := cost
			nu := val.loc
			for i := 1; i <= 10; i++ {
				nu = add(nu, nudir)
				if nu.r < 0 || nu.r >= len(grid) || nu.c < 0 || nu.c >= len(grid[0]) {
					// out of bounds
					continue
				}
				nucost += int(grid[nu.r][nu.c] - byte('0'))
				if i >= 4 {
					loc := params{nu, polarityfromdir(nudir)}
					// fmt.Printf("%v -> %v: %d\n", val, loc, nucost)
					heap.Push(&pq, &Item{value: loc, priority: nucost})
				}
			}
		}
	}

	polarities := []polarity{northsouth, eastwest}
	best := 10000000
	for _, pol := range polarities {
		res := dists[params{goal, pol}]
		if res < best {
			best = res
		}
	}
	fmt.Printf("goal: %d\n", best)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// location in a grid
type loc = struct {
	r int
	c int
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    params
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Return lower cost first
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

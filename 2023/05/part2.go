package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type seedRange = struct {
	Start int
	Size  int
}

type mappingRange = struct {
	Source int
	Target int
	Size   int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seeds := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
	seedRanges := []seedRange{}
	for len(seeds) > 0 {
		seedRanges = append(seedRanges, seedRange{Start: atoi(seeds[0]), Size: atoi(seeds[1])})
		seeds = seeds[2:]
	}

	maps := make([]map[int]mappingRange, 7)
	curMap := 0

	scanner.Scan()
	scanner.Scan()
	maps[curMap] = map[int]mappingRange{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			curMap += 1
			maps[curMap] = map[int]mappingRange{}
			scanner.Scan()
			continue
		}
		parsed := strings.Split(line, " ")
		dest, source, size := atoi(parsed[0]), atoi(parsed[1]), atoi(parsed[2])
		maps[curMap][source] = mappingRange{Source: source, Target: dest, Size: size}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	applyMap := func(i int, m mappingRange) int {
		if i < m.Source {
			log.Fatalf("%d outside range %v", i, m)
		}
		if i >= m.Source+m.Size {
			log.Fatalf("%d outside range %v", i, m)
		}
		res := i - m.Source + m.Target
		if res < 0 {
			log.Fatalf("%d mapped to a negative number: %d (mapping: %v)", i, i-m.Source+m.Target, m)
		}
		return res
	}

	assertEquals := func(i, j int) {
		if i != j {
			log.Fatalf("Expected %d == %d", i, j)
		}
	}

	applyMapToRange := func(mapping mappingRange, chunk seedRange) (mapped []seedRange, remaining []seedRange) {
		// Four cases to consider
		clo := chunk.Start
		chi := chunk.Start + chunk.Size
		mlo := mapping.Source
		mhi := mapping.Source + mapping.Size

		if mhi <= clo {
			// no overlap; leave this chunk intact for the next mapping
			remaining = []seedRange{chunk}
			return
		}
		if mlo >= chi {
			// no overlap; leave this chunk intact for the next mapping
			remaining = []seedRange{chunk}
			return
		}
		if mlo <= clo && mhi >= chi {
			// mapping overlaps chunk entirely
			// map the whole chunk: [clo, chi]
			mapped = []seedRange{{Start: applyMap(clo, mapping), Size: chi - clo}}
			return
		}
		if mlo > clo && mhi < chi {
			// mapping splits chunk
			// map the map range in its entirety [mlo, mhi]
			mapped = []seedRange{{Start: applyMap(mlo, mapping), Size: mhi - mlo}}

			assertEquals(applyMap(mlo, mapping)+mhi-mlo-1, applyMap(mhi-1, mapping))
			// split chunk into [clo, mlo] and [mhi, chi]
			remaining = []seedRange{
				{Start: clo, Size: mlo - clo},
				{Start: mhi, Size: chi - mhi},
			}
			return
		}
		if mhi >= chi {
			// mapping intersects right side of chunk
			mapped = []seedRange{{Start: applyMap(mlo, mapping), Size: chi - mlo}}
			remaining = []seedRange{{Start: clo, Size: mlo - clo}}
			return
		}
		if mlo <= clo {
			// mapping intersects left side of chunk
			mapped = []seedRange{{Start: applyMap(clo, mapping), Size: mhi - clo}}
			remaining = []seedRange{{Start: mhi, Size: chi - mhi}}
			return
		}
		log.Fatalf("Couldn't figure out what to do with c[%d,%d] and m[%d,%d]", clo, chi, mlo, mhi)
		return
	}

	// Go through each layer, and map all the chunks (or don't map them)
	for mi, mappingLayer := range maps {
		// Map all chunks in seedRanges in this layer, place them in nextGenSeedRanges
		fmt.Printf("=== Layer %d: %v chunks\n", mi, len(seedRanges))
		mappedRanges := []seedRange{}
		for _, mapping := range mappingLayer {
			unmappedRanges := []seedRange{}
			for _, chunk := range seedRanges {
				mapped, remaining := applyMapToRange(mapping, chunk)
				mappedRanges = append(mappedRanges, mapped...)
				unmappedRanges = append(unmappedRanges, remaining...)
			}
			seedRanges = unmappedRanges
		}
		// Unmapped pieces get mapped as-is
		seedRanges = append(mappedRanges, seedRanges...)
	}

	minSeed := 100000000000
	for _, r := range seedRanges {
		if r.Start < minSeed {
			minSeed = r.Start
		}
	}
	fmt.Println(minSeed)

}

// cheap atoi
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
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

type nullableInt struct {
	value int
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("no value in max function")
	}

	var max *nullableInt
	for _, value := range values {
		if max == nil || value > max.value {
			max = &nullableInt{value}
		}
	}
	return max.value
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("no value in min function")
	}

	var min *nullableInt
	for _, value := range values {
		if min == nil || value < min.value {
			min = &nullableInt{value}
		}
	}
	return min.value
}

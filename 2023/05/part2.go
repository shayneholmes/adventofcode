package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	fmt.Println(seedRanges)

	maps := map[int]map[int]mappingRange{}
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

	consolidateRanges := func(rs []seedRange) []seedRange {
		sort.Slice(rs, func(i, j int) bool {
			return rs[i].Start < rs[j].Start
		})
		nextGen := []seedRange{}
		i := 0
		size := 0
		for i < len(rs) {
			start := rs[i].Start
			end := rs[i].Start + rs[i].Size
			// look for any items we can fold in
			j := i + 1
			for ; j < len(rs) && rs[j].Start <= end; j++ {
				jEnd := rs[j].Start + rs[j].Size
				if jEnd > end {
					end = jEnd
				}
			}
			nextGen = append(nextGen, seedRange{start, end - start})
			i = j
			size += end - start
		}
		for i := 0; i < len(nextGen)-1; i++ {
			if !(nextGen[i].Start+nextGen[i].Size < nextGen[i+1].Start) {
				log.Fatalf("Expected disjoint ranges, but %d ends at %d, and %d starts at %d",
					i,
					nextGen[i].Start+nextGen[i].Size,
					i+1,
					nextGen[i+1].Start,
				)
			}
		}
		fmt.Printf("Consolidated to size %d\n", size)
		return nextGen
	}

	assertEquals := func(i, j int) {
		if i != j {
			log.Fatalf("Expected %d == %d", i, j)
		}
	}

	for mi, mappingLayer := range maps {
		seedRanges = consolidateRanges(seedRanges)
		fmt.Printf("=== Layer %d: %v\n", mi, len(seedRanges))
		nextGenSeedRanges := []seedRange{}
		for _, r := range seedRanges {
			chunks := []seedRange{r} // seeds that need to find a home
			// Each bit of this range needs to be mapped into a new range, and hopefully they don't overlap
			// Find map ranges that intersect, and map those
			// Assume everything only gets mapped once
			for start, mapping := range mappingLayer {
				chunksRemaining := []seedRange{} // chunks that haven't been mapped should go here
				for _, chunk := range chunks {
					// Four cases to consider
					clo := chunk.Start
					chi := chunk.Start + chunk.Size
					mlo := start
					mhi := start + mapping.Size
					if mhi <= clo {
						// no overlap; leave this chunk intact for the next mapping
						chunksRemaining = append(chunksRemaining, chunk)
						continue
					}
					if mlo >= chi {
						// no overlap; leave this chunk intact for the next mapping
						chunksRemaining = append(chunksRemaining, chunk)
						continue
					}
					if mlo <= clo && mhi >= chi {
						// mapping overlaps chunk entirely
						// map the whole chunk: [clo, chi]
						nextGenSeedRanges = append(nextGenSeedRanges, seedRange{
							Start: applyMap(clo, mapping),
							Size:  chi - clo,
						})
						continue
					}
					if mlo > clo && mhi < chi {
						// mapping splits chunk
						// map the map range in its entirety [mlo, mhi]
						nextGenSeedRanges = append(nextGenSeedRanges, seedRange{
							Start: applyMap(mlo, mapping),
							Size:  mhi - mlo,
						})

						assertEquals(applyMap(mlo, mapping)+mhi-mlo-1, applyMap(mhi-1, mapping))
						// split chunk into [clo, mlo] and [mhi, chi]
						chunksRemaining = append(chunksRemaining,
							seedRange{Start: clo, Size: mlo - clo},
							seedRange{Start: mhi, Size: chi - mhi},
						)
						assertEquals(chunk.Size, mhi-mlo+mlo-clo+chi-mhi)
						continue
					}
					if mhi >= chi {
						// mapping intersects right side of chunk
						// map [mlo, chi]
						nextGenSeedRanges = append(nextGenSeedRanges, seedRange{
							Start: applyMap(mlo, mapping),
							Size:  chi - mlo,
						})
						assertEquals(applyMap(mlo, mapping)+chi-mlo-1, applyMap(chi-1, mapping))
						// leave [clo, mlo]
						chunksRemaining = append(chunksRemaining,
							seedRange{Start: clo, Size: mlo - clo},
						)
						assertEquals(chunk.Size, chi-mlo+mlo-clo)
						continue
					}
					if mlo <= clo {
						// mapping intersects left side of chunk
						// map [clo, mhi]
						nextGenSeedRanges = append(nextGenSeedRanges, seedRange{
							Start: applyMap(clo, mapping),
							Size:  mhi - clo,
						})
						assertEquals(applyMap(clo, mapping)+mhi-clo-1, applyMap(mhi-1, mapping))
						// leave [mhi, chi]
						chunksRemaining = append(chunksRemaining,
							seedRange{Start: mhi, Size: chi - mhi},
						)
						assertEquals(chunk.Size, mhi-clo+chi-mhi)
						continue
					}
					log.Fatalf("Couldn't figure out what to do with c[%d,%d] and m[%d,%d]", clo, chi, mlo, mhi)
				}
				chunks = chunksRemaining
			}
			// Unmapped pieces get mapped as-is
			nextGenSeedRanges = append(nextGenSeedRanges, chunks...)
			// nextGenSeedRanges = consolidateRanges(nextGenSeedRanges)
		}
		// nextGenSeedRanges = consolidateRanges(nextGenSeedRanges)
		seedRanges = nextGenSeedRanges
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

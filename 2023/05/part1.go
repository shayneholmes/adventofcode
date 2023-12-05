package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type mappingRange = struct {
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

	maps := map[int]map[int]mappingRange{}
	curMap := 0

	scanner.Scan()
	scanner.Scan()
	maps[curMap] = map[int]mappingRange{}

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%v\n", line)
		if line == "" {
			curMap += 1
			maps[curMap] = map[int]mappingRange{}
			scanner.Scan()
			fmt.Println("skipping empty line")
			continue
		}
		parsed := strings.Split(line, " ")
		dest, source, size := atoi(parsed[0]), atoi(parsed[1]), atoi(parsed[2])
		maps[curMap][source] = mappingRange{Target: dest, Size: size}
	}

	mapValue := func(source int, layer int) int {
		for start, mapping := range maps[layer] {
			if source >= start && source < start+mapping.Size {
				return mapping.Target + source - start
			}
		}
		return source // unmapped values stay the same
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minSeed := 100000000000
	for _, sstr := range seeds {
		s := atoi(sstr)
		for i := 0; i <= curMap; i++ {
			s = mapValue(s, i)
		}
		fmt.Println(s)
		if s < minSeed {
			minSeed = s
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

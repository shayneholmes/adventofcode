package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grids := [][]string{}
	grids = append(grids, []string{})
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			grids = append(grids, []string{})
		} else {
			grids[len(grids)-1] = append(grids[len(grids)-1], line)
		}
	}

	sum := 0
	fmt.Printf("%v grids\n", len(grids))
	for _, g := range grids {
		lr := len(g)
		lc := len(g[0])
		// true if the reflection is before this row (i.e. there are r rows before the line)
		isHorizReflectionAboutAxis := func(axis int) bool {
			rmax := min(lr-axis, axis) // how many rows we'll compare
			for r := 0; r < rmax; r++ {
				for c := 0; c < lc; c++ {
					if g[axis-1-r][c] != g[axis+r][c] {
						return false
					}
				}
			}
			return true
		}
		isVertReflectionAboutAxis := func(axis int) bool {
			cmax := min(lc-axis, axis) // how many rows we'll compare
			for c := 0; c < cmax; c++ {
				for r := 0; r < lr; r++ {
					if g[r][axis-1-c] != g[r][axis+c] {
						return false
					}
				}
			}
			return true
		}
		done := false
		for r := 1; r <= len(g)-1; r++ {
			if isHorizReflectionAboutAxis(r) {
				sum += r * 100
				for _, r := range g {
					fmt.Println(r)
				}
				fmt.Printf("Horizontal: %v\n", r)
				if done {
					log.Fatalf("two lines!")
				}
				done = true
			}
		}
		for c := 1; c <= len(g[0])-1; c++ {
			if isVertReflectionAboutAxis(c) {
				sum += c
				for _, r := range g {
					fmt.Println(r)
				}
				fmt.Printf("Vertical: %v\n", c)
				if done {
					log.Fatalf("two lines!")
				}
				done = true
			}
		}
		if !done {
			for _, r := range g {
				fmt.Println(r)
			}
			log.Fatalf("No lines!")
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

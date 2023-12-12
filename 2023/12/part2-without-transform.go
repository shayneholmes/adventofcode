package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.Split(line, " ")
		input, runs := sp[0], sp[1]
		input = strings.Join([]string{input, input, input, input, input}, "?")
		runs = strings.Join([]string{runs, runs, runs, runs, runs}, ",")
		sum += waysToMatch(ints(runs), input)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type parameterSet = struct {
	r int // run index
	i int // input index
}

func waysToMatch(runs []int, input string) int {
	li := len(input)
	lr := len(runs)
	cached := map[parameterSet]int{}
	var waystomatch func(r, i int) int
	waystomatch = func(r, i int) int {
		params := parameterSet{r, i}
		// check cache
		if cached, ok := cached[params]; ok {
			return cached
		}
		inner := func(r, i int) int {
			// Compute ways to match, based on recursive calls

			if i >= li {
				// end of the input, or match
				if r >= lr {
					// We have run out of runs, too. Success!
					return 1
				}
				return 0
			}

			{
				// bound: do we have enough room for the remaining runs?
				remaining := li - i
				necessary := -1 // no padding needed at the end
				for _, r := range runs[r:] {
					necessary += r + 1 // A run needs padding around it, too.
				}
				if necessary > remaining {
					return 0
				}
			}
			// It might fit.
			ways := 0
			in := input[i]
			if r < lr {
				// There are remaining runs of damaged springs. One potential option is to start one now.
				canUseRun := true
				runlength := runs[r]
				for j := 0; j < runlength; j++ {
					if input[i+j] == '.' {
						canUseRun = false
					}
					// A run must also have padding, unless it's at the end
					if i+runs[r] < li && input[i+runs[r]] == '#' {
						canUseRun = false
					}
				}
				if canUseRun {
					ways += waystomatch(r+1, i+runs[r]+1)
				}
			}
			if in != '#' {
				// We also have the option of _not_ starting a run.
				ways += waystomatch(r, i+1)
			}
			return ways
		}
		res := inner(r, i)
		cached[params] = res
		return res
	}
	return waystomatch(0, 0)
}

// return all ints in a string of text
func ints(s string) []int {
	re := regexp.MustCompile(`-?\b\d+\b`)
	intss := re.FindAllString(s, -1)
	ints := make([]int, len(intss))
	for i := 0; i < len(ints); i++ {
		ints[i] = atoi(intss[i])
	}
	return ints
}

// cheap atoi
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

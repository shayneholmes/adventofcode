package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hands := []hand{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ps := strings.Split(line, " ")
		fmt.Printf("%q -> %d\n", ps[0], classify(ps[0]))
		hands = append(
			hands,
			hand{
				item:     ps[0],
				bid:      atoi(ps[1]),
				handType: classify(ps[0]),
			},
		)
	}

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]
		if h1.handType != h2.handType {
			return h1.handType > h2.handType
		}
		s1 := h1.item
		s2 := h2.item
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return orderByChar[s1[i]] > orderByChar[s2[i]]
			}
		}

		return false
	})

	var sum int64 = 0
	for rank, hand := range hands {
		sum += (int64)((rank + 1) * hand.bid)
		fmt.Printf("%v %d %9d %v*%4v\n", hand.item, hand.handType, sum, rank+1, hand.bid)
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type hand = struct {
	item     string
	bid      int
	handType int
}

const (
	fiveofkind = iota
	fourofkind
	fullhouse
	threeofkind
	twopair
	onepair
	highcard
)

func classify(s string) int {
	jokers := 0
	ranks := map[rune]int{}
	for _, ch := range s {
		if ch == 'J' {
			jokers += 1
		} else {
			ranks[ch] += 1
		}
	}
	best := 0
	num := 1
	bestrank := 'o'
	for rank, item := range ranks {
		if item > best {
			best = item
			bestrank = rank
			num = 1
		} else if item == best {
			num += 1
		}
	}
	best += jokers
	ranks[bestrank] += jokers
	fmt.Printf("%d %d %v\n", best, num, ranks)
	switch best {
	case 5:
		return fiveofkind
	case 4:
		return fourofkind
	case 3:
		nextbest := 0
		for _, item := range ranks {
			if item == best {
				continue
			}
			if item > nextbest {
				nextbest = item
			}
		}
		fmt.Printf("%d %d %d %v\n", best, num, nextbest, ranks)
		if nextbest == 2 {
			return fullhouse
		}
		return threeofkind
	case 2:
		if num == 2 {
			return twopair
		}
		fmt.Println("onepair")
		fmt.Printf("%d %d %v\n", best, num, ranks)
		return onepair
	case 1:
		return highcard
	}
	log.Fatalf("no way to handle %v\n", s)
	return 11111
}

var orderByChar = map[byte]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 13, // weak joker
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
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

// return all floats in a string of text
func floats(s string) []float64 {
	re := regexp.MustCompile(`-?\d+(\.\d+)?`)
	intss := re.FindAllString(s, -1)
	floats := make([]float64, len(intss))
	for i := 0; i < len(floats); i++ {
		f, err := strconv.ParseFloat(intss[i], 64)
		if err != nil {
			log.Fatalf("not an float: %q", s)
		}
		floats[i] = f
	}
	return floats
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

package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	readingRules = iota
	readingParts
)

type (
	attrrange struct {
		lo int // inclusive
		hi int // exclusive
	}

	partrange = map[rune]attrrange

	state = struct {
		partrange partrange
		label     label
	}

	label = string

	rule struct {
		field  rune
		op     rune
		rand   int
		target label
	}
)

func getattr(p partrange, field rune) attrrange {
	if attr, ok := p[field]; ok {
		return attr
	}
	log.Fatalf("bad field %q", field)
	return attrrange{}
}

// returns a new partrange with the attribute set differently.
func setattr(p partrange, field rune, a attrrange) partrange {
	nu := maps.Clone(p)
	nu[field] = a
	return nu
}

// Split a range at the indicated field, return lo and hi
func splitrange(p partrange, field rune, newlo int) (partrange, partrange) {
	attr := getattr(p, field)
	if attr.lo >= newlo {
		return nil, p
	}
	if attr.hi <= newlo {
		return p, nil
	}

	attrlo := attr
	attrlo.hi = newlo
	plo := setattr(p, field, attrlo)

	attrhi := attr
	attrhi.lo = newlo
	phi := setattr(p, field, attrhi)

	return plo, phi
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	workflows := map[label][]rule{}
	mode := readingRules
	for scanner.Scan() {
		line := scanner.Text()
		ruleregex := regexp.MustCompile(`(\w+)\{([^}]+)}`)
		if line == "" {
			mode = readingParts
			continue
		}
		if mode == readingRules {
			matches := ruleregex.FindStringSubmatch(line)[1:]
			label := matches[0]
			rules := strings.Split(matches[1], ",")
			workflows[label] = []rule{}
			for _, r := range rules {
				target := ""
				op := rune(0)
				field := rune(0)
				rand := 0
				components := strings.Split(r, ":")
				if len(components) == 1 {
					target = r
				} else {
					target = components[1]
					conditionStr := components[0]
					field = rune(conditionStr[0])
					op = rune(conditionStr[1])
					rand = atoi(conditionStr[2:])
				}
				workflows[label] = append(workflows[label],
					rule{field, op, rand, target},
				)
			}
		} else {
			continue
		}
	}

	var acceptedcombos func(partrange, label) int
	acceptedcombos = func(range_ partrange, label string) int {
		if range_ == nil {
			return 0
		}
		if label == "A" {
			rangesize := 1
			for _, attr := range range_ {
				rangesize *= attr.hi - attr.lo
			}
			return rangesize
		}
		if label == "R" {
			return 0
		}
		res := 0
		for _, rule := range workflows[label] {
			if range_ == nil { // nothing to evaluate the rule against
				break
			}
			switch rule.op {
			case rune(0): // no predicate, move the whole range
				res += acceptedcombos(range_, rule.target)
				range_ = nil
			case '<':
				lo, hi := splitrange(range_, rule.field, rule.rand)
				res += acceptedcombos(lo, rule.target)
				range_ = hi
			case '>':
				lo, hi := splitrange(range_, rule.field, rule.rand+1)
				res += acceptedcombos(hi, rule.target)
				range_ = lo
			}
		}
		return res
	}

	// Create a starting range of everything
	attrLo := 1
	attrHi := 4001 // exclusive
	fullrange := partrange{
		'x': attrrange{attrLo, attrHi},
		'm': attrrange{attrLo, attrHi},
		'a': attrrange{attrLo, attrHi},
		's': attrrange{attrLo, attrHi},
	}
	fmt.Printf("%d\n", acceptedcombos(fullrange, "in"))
}

// cheap atoi
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

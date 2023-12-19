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

const (
	readingRules = iota
	readingParts
)

type (
	attrrange struct {
		lo int // inclusive
		hi int // exclusive
	}

	partrange = struct {
		x attrrange
		m attrrange
		a attrrange
		s attrrange
	}

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

var (
	empty_attr = attrrange{0, 0}
	empty      = partrange{empty_attr, empty_attr, empty_attr, empty_attr}
)

func getattr(p partrange, field rune) attrrange {
	switch field {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	default:
		log.Fatalf("bad field %q", field)
		return attrrange{}
	}
}

func setattr(p partrange, field rune, a attrrange) partrange {
	switch field {
	case 'x':
		p.x = a
	case 'm':
		p.m = a
	case 'a':
		p.a = a
	case 's':
		p.s = a
	default:
		log.Fatalf("bad field %q", field)
	}
	return p
}

// Split a range at the indicated field, return lo and hi
func splitrange(p partrange, field rune, newlo int) (partrange, partrange) {
	attr := getattr(p, field)
	if attr.lo >= newlo {
		return empty, p
	}
	if attr.hi <= newlo {
		return p, empty
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

	// Create a starting range of everything
	attrLo := 1
	attrHi := 4001 // exclusive
	r := partrange{
		x: attrrange{attrLo, attrHi},
		m: attrrange{attrLo, attrHi},
		a: attrrange{attrLo, attrHi},
		s: attrrange{attrLo, attrHi},
	}

	acceptedcombos := 0
	q := []state{{r, "in"}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		range_ := cur.partrange
		if cur.label == "A" {
			combos := (range_.x.hi - range_.x.lo) * (range_.m.hi - range_.m.lo) *
				(range_.a.hi - range_.a.lo) * (range_.s.hi - range_.s.lo)
			acceptedcombos += combos
		}
		if cur.label == "R" {
			continue
		}
		moveRangeToLabel := func(range_ partrange, l label) {
			if range_ == empty {
				return
			}
			q = append(q, state{range_, l})
		}
		for _, rule := range workflows[cur.label] {
			if range_ == empty { // nothing to evaluate the rule against
				break
			}
			switch rule.op {
			case rune(0): // no predicate, move the whole range
				moveRangeToLabel(range_, rule.target)
				range_ = empty
			case '<':
				lo, hi := splitrange(range_, rule.field, rule.rand)
				moveRangeToLabel(lo, rule.target)
				range_ = hi
			case '>':
				lo, hi := splitrange(range_, rule.field, rule.rand+1)
				moveRangeToLabel(hi, rule.target)
				range_ = lo
			}
		}
	}

	fmt.Printf("%d\n", acceptedcombos)
}

// cheap atoi
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

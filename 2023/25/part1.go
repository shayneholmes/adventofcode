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

	nodes := map[string]bool{}
	edges := map[string]map[string]bool{}

	scanner := bufio.NewScanner(file)
	edgecount := 0
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, " ")
		one := l[0][:len(l[0])-1]
		nodes[one] = true
		for _, other := range l[1:] {
			edgecount += 1
			nodes[other] = true
			if edges[one] == nil {
				edges[one] = map[string]bool{}
			}
			edges[one][other] = true
			if edges[other] == nil {
				edges[other] = map[string]bool{}
			}
			edges[other][one] = true
		}
	}
	fmt.Printf("nodes: %d\n", len(nodes))
	fmt.Printf("edges: %v\n", edgecount)

	// According to this problem, removing three edges can make a cut. That means
	// that all nodes are connected by at least three edges. (If they were
	// connected by less, that would be the min cut number.) I think that this
	// means that we can find the shortest path between them, and then remove the edges used
	// from consideration, and find another path that doesn't use those edges,
	// and do it again. If we can't find a fourth path, then it must be on the
	// other side of the cut. Otherwise, they're in the same group.

	groups := map[string]string{} // map from element to the earliest group leader

	getgroup := func(s string) string {
		group := s
		for {
			if g, ok := groups[group]; ok && group != g {
				group = g
			} else {
				groups[s] = group
				return group
			}
		}
	}

	setgroup := func(el, g string) {
		for groups[el] != g {
			nu := groups[el]
			groups[el] = g
			el = nu
		}
	}

	// Find the number of disjoint paths between s1 and s2
	hasmorepathsthan := func(s1, s2 string, paths int) bool {
		edgekey := func(s1, s2 string) string {
			first := min(s1, s2)
			last := max(s1, s2)
			return fmt.Sprintf("%s-%s", first, last)
		}
		usededges := map[string]bool{}
		pathmatches := 0
		edgecount := 0
	findpath:
		for {
			q := []string{s1}
			arrivedfrom := map[string]string{}
			for len(q) > 0 {
				edgecount += 1
				node := q[0]
				q = q[1:]
				if node == s2 {
					// Found a path
					pathmatches += 1
					if pathmatches > paths {
						return true
					}
					for n := s2; n != s1; n = arrivedfrom[n] {
						usededges[edgekey(n, arrivedfrom[n])] = true
					}
					continue findpath
				}
				for neighbor := range edges[node] {
					if _, ok := arrivedfrom[neighbor]; !ok && !usededges[edgekey(node, neighbor)] {
						arrivedfrom[neighbor] = node
						q = append(q, neighbor)
					}
				}
			}
			// ran out of nodes in our exploration
			return false
		}
	}

	joingroup := func(one, other string) {
		g1 := getgroup(one)
		g2 := getgroup(other)
		group := min(g1, g2)
		setgroup(one, group)
		setgroup(other, group)
	}

	// merge groups as we're able
	for one := range nodes {
		// its neighbors should be fast to search
		for other := range edges[one] {
			if getgroup(one) == getgroup(other) {
				// we already know they're in the same group
				continue
			}
			if hasmorepathsthan(one, other, 3) {
				// They're in the same group.
				joingroup(one, other)
			}
		}
	}

	membership := map[string]int{}
	for el := range groups {
		membership[getgroup(el)]++
	}
	fmt.Println(membership)

	prod := uint64(1)
	for _, s := range membership {
		prod *= uint64(s)
	}
	fmt.Println(prod)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

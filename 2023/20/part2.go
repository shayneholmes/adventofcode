package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

const (
	flipflop = iota
	conjunction
	broadcaster
)

type module struct {
	type_ int
	dests []string
}

const (
	lo = false
	hi = true
)

type pulse struct {
	hilo   bool
	source string
	dest   string
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	modules := map[string]module{}
	sources := map[string][]string{}
	lineregex := regexp.MustCompile(`^(\S+) -> (.*)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := lineregex.FindStringSubmatch(line)[1:]
		label := matches[0]
		dests := strings.Split(matches[1], ", ")
		if label[0] == '%' {
			label = label[1:]
			modules[label] = module{flipflop, dests}
		} else if label[0] == '&' {
			label = label[1:]
			modules[label] = module{conjunction, dests}
		} else if label == "broadcaster" {
			modules[label] = module{broadcaster, dests}
		}
		for _, d := range dests {
			sources[d] = append(sources[d], label)
		}
	}

	flipmemory := map[string]bool{}
	conjunctionmemory := map[string]map[string]bool{}

	fmt.Println(modules)
	// init conjunctions
	for label, m := range modules {
		for _, d := range m.dests {
			if modules[d].type_ == conjunction {
				if conjunctionmemory[d] == nil {
					conjunctionmemory[d] = map[string]bool{}
				}
				conjunctionmemory[d][label] = false
			}
		}
	}

	modulelabels := make([]string, 0, len(modules))
	for m := range modules {
		modulelabels = append(modulelabels, m)
	}
	slices.Sort(modulelabels)

	rootModule := "rx"
	combinator := sources[rootModule][0]

	modulestate := func(label string) string {
		switch modules[label].type_ {
		case broadcaster:
			return "Broadcaster"
		case flipflop:
			return fmt.Sprintf("F:%s:%v", label, flipmemory[label])
		case conjunction:
			var sb strings.Builder
			cm := conjunctionmemory[label]
			for _, l := range modulelabels {
				if v, ok := cm[l]; ok {
					sb.WriteString(fmt.Sprintf("%v=%v;", l, v))
				}
			}
			return fmt.Sprintf("C:%s:%s:", label, sb.String())
		default:
			log.Fatalf("bad: %v", label)
			return "BAD"
		}
	}

	presses := 1
	lastChange := map[string]int{}
	cycles := map[string]int{}
	press := func() {
		// press the button
		q := []pulse{{lo, "button", "broadcaster"}}
		// wait for all pulses
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			m := modules[cur.dest]
			if cur.dest == combinator && cur.hilo != conjunctionmemory[combinator][cur.source] {
				if cur.hilo == hi {
					// change to hi
					if l, ok := lastChange[cur.source]; !ok {
						lastChange[cur.source] = presses
					} else {
						cycles[cur.source] = presses - l
						if presses-l != l {
							log.Fatalf("Cycle didn't start at zero! %d, %d", l, presses)
						}
					}
					if len(cycles) == len(sources[combinator]) {
						// we found cycles for each
						c := make([]int, 0, len(cycles))
						for _, v := range cycles {
							c = append(c, v)
						}
						fmt.Printf("%d\n", lcm(c...))
					}
				}
				fmt.Printf("%d: %s -> %v -> %s\n", presses, cur.source, cur.hilo, cur.dest)
			}
			switch m.type_ {
			case broadcaster:
				// send pulse to all dests
				for _, d := range m.dests {
					q = append(q, pulse{cur.hilo, cur.dest, d})
				}
			case flipflop:
				if cur.hilo == hi {
					continue
				} else {
					state := flipmemory[cur.dest]
					flipmemory[cur.dest] = !state
					for _, d := range m.dests {
						q = append(q, pulse{!state, cur.dest, d})
					}
				}
			case conjunction:
				conjunctionmemory[cur.dest][cur.source] = cur.hilo
				allon := true
				for _, on := range conjunctionmemory[cur.dest] {
					if !on {
						allon = false
						break
					}
				}
				for _, d := range m.dests {
					q = append(q, pulse{!allon, cur.dest, d})
				}
			}
		}
	}

	// find some cycles
	lastState := ""
	for ; presses < 10000000; presses++ {
		press()
		st := modulestate(combinator)
		if st != lastState {
			fmt.Printf("%d: Changed state! %s\n", presses, st)
			lastState = st
		}
	}

	// // use the cycle info gathered above?
	// var cycle func(label string) int
	// cycle = func(label string) int {
	// 	m := modules[label]
	// 	switch m.type_ {
	// 	case broadcaster:
	// 		return 1
	// 	case flipflop:
	// 		prod := 1
	// 		for _, s := range sources[label] {
	// 			prod = lcm(prod, cycle(s))
	// 		}
	// 		return 2 * prod
	// 	case conjunction:
	// 		prod := 1
	// 		for _, s := range sources[label] {
	// 			prod = lcm(prod, cycle(s))
	// 		}
	// 		return prod
	// 	}
	// 	return 0
	// }
	// fmt.Println(cycle("rx"))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// absolute value
func abs[V Number](i V) V {
	return max(i, -i)
}

// greatest common divisor
func gcd(values ...int) int {
	if len(values) == 0 {
		panic("no value in gcd function")
	}
	gcd := values[0]
	for _, i := range values[1:] {
		gcd = gcd_(gcd, i)
	}
	return gcd
}

func gcd_(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd_(b, a%b)
}

func lcm(values ...int) int {
	if len(values) == 0 {
		panic("no value in lcm function")
	}
	lcm := values[0]
	for _, i := range values[1:] {
		lcm = lcm_(lcm, i)
	}
	return lcm
}

func lcm_(a, b int) int {
	return a / gcd_(a, b) * b
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

// Types from constraints package
type (
	Signed interface {
		~int | ~int8 | ~int16 | ~int32 | ~int64
	}
	Unsigned interface {
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
	}
	Integer interface{ Signed | Unsigned }
	Float   interface{ ~float32 | ~float64 }
	Complex interface{ ~complex64 | ~complex128 }
	Number  interface{ Integer | Float }
	Ordered interface{ Integer | Float | ~string }
)

// flip-flops only send pulses when they get lows
// * two flip-flops halve the volume of low pulses
// conjunctions send them all the time, if they remember a high for all then they send low, otherwise they're high
// rx requires zh to send a low pulse
// zh requires all of
// zp sends a lo every 3761

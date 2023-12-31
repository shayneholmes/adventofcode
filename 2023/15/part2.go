package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hash(s string) int {
	val := 0
	for _, ch := range s {
		val += int(byte(ch))
		val *= 17
		val %= 256
	}
	return val
}

type label = string
type focallength = int
type pos = int

type lens = struct {
	label string
	f     int
}
type box = []lens

func findLabelInBox(label string, box box) int {
	for i, l := range box {
		if label == l.label {
			return i
		}
	}
	return -1
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, ",")

		boxes := make([]box, 256)
		vals := 0
		for _, s := range strs {
			if strings.Contains(s, "=") {
				v := strings.Split(s, "=")
				label := v[0]
				// If the operation character is an equals sign (=), it will be followed by a
				// number indicating the focal length of the lens that needs to go into the
				// relevant box; be sure to use the label maker to mark the lens with the label
				// given in the beginning of the step so you can find it later. There are two
				// possible situations:
				box := hash(label)
				focallength := atoi(v[1])

				if i := findLabelInBox(label, boxes[box]); i > -1 {
					// If there is already a lens in the box with the same label, replace the old
					// lens with the new lens: remove the old lens and put the new lens in its
					// place, not moving any other lenses in the box.
					boxes[box][i] = lens{label, focallength}
				} else {
					// If there is not already a lens in the bx with the same label, add the lens
					// to the box immediately behind any lenses already in the box. Don't move any
					// of the other lenses when you do this. If there aren't any lenses in the box,
					// the new lens goes all the way to the front of the box.
					boxes[box] = append(boxes[box], lens{label, focallength})
				}
			} else {
				// If the operation character is a dash (-), go to the relevant box and remove the
				// lens with the given label if it is present in the box. Then, move any remaining
				// lenses as far forward in the box as they can go without changing their order,
				// filling any space made by removing the indicated lens. (If no lens in that box
				// has the given label, nothing happens.)

				label := strings.TrimSuffix(s, "-")
				box := hash(label)
				if i := findLabelInBox(label, boxes[box]); i > -1 {
					boxes[box][i] = lens{}
				}
			}
			vals += hash(s)
		}
		// To confirm that all of the lenses are installed correctly, add up the
		// focusing power of all of the lenses. The focusing power of a single lens
		// is the result of multiplying together:

		// One plus the box number of the lens in question.
		// The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
		// The focal length of the lens.
		sum := 0
		for b, box := range boxes {
			lensId := 1
			for _, lens := range box {
				if lens.label != "" {
					// lens is in here
					sum += (b + 1) * lensId * lens.f
					lensId += 1
				}
			}
		}
		fmt.Println(sum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// cheap atoi
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("not an int: %q", s)
	}
	return i
}

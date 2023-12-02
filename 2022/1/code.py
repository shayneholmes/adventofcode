#!/usr/bin/env python

elves = []
cur = 0
with open("input.txt") as f:
    for line in f:
        if line == "\n":
            elves.append(cur)
            cur = 0
            continue
        cur += int(line)
print(sum(sorted(elves)[-3:]))

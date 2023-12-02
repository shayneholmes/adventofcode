import re
import itertools
import functools
import math

rocks = [ # list of bytes, bottom to top, lined up so their first bit is at position 2 (lines up with walls)
    # ####
    [
        0b11110000,
    ],

    #
    # .#.
    # ###
    # .#.
    [
        0b01000000,
        0b11100000,
        0b01000000,
    ],

    #
    # ..#
    # ..#
    # ###
    [
        0b11100000,
        0b00100000,
        0b00100000,
    ],

    # #
    # #
    # #
    # #
    [
        0b10000000,
        0b10000000,
        0b10000000,
        0b10000000,
    ],

    # ##
    # ##
    [
        0b11000000,
        0b11000000,
    ],
]

with open("input.txt") as f:
    pattern = [-1 if c == '<' else 1 for c in f.readline().strip()]

width = 7
grid = []
def newrow():
    # has walls
    return 0b100000001

# return highest row of rock
# the floor is -1
def highestrow():
    for i in range(len(grid)-1, -1, -1):
        if grid[i] > 0b100000001:
            return i
    return -1

print(highestrow())

# true if rock collides with rocks or walls in this position
def collision(rock, rr, rc):
    for (nr, row) in enumerate(rock):
        if rc < 0: # bumps into wall
            return True
        r = nr + rr
        if r < 0: # into floor
            return True
        while len(grid) <= r:
            grid.append(newrow())
        if row >> rc & grid[nr + rr]:
            return True
    return False

def fits(rock, rr, rc):
    return not collision(rock, rr, rc)

def draw(rock, rr, rc):
    if not fits(rock, rr, rc):
        raise Exception("Trying to draw but it doesn't fit: %d,%d"%(rr,rc))
    # print("drawing rock at %d,%d"%(rr,rc))
    for (nr, row) in enumerate(rock):
        r = nr + rr
        while len(grid) <= r:
            grid.append(newrow())
        grid[r] |= (row >> rc)

def printgrid():
    print("------")
    for i in range(len(grid)-1,max(-1,len(grid)-10),-1):
        print(bin(grid[i]))
    print("...")
    for i in range(min(len(grid)-1,10),-1,-1):
        print(bin(grid[i]))

turns = len(pattern) * len(rocks)
turns = 1000000000000
print(turns)
print(1000000000000 // turns)
# turns = 1000000000000 % (len(pattern) * len(rocks))
print(turns)
gascounter = 0
last = 0
heightbyturn = {0: 0}
last10state = {} # set -> turn
for turn in range(128+2*1700):
    if turn % 100000 == 0:
        print("progress: %f%%"%(100.0 * turn / turns))
    # summon the rock
    rock = rocks[turn % len(rocks)]
    rr = highestrow() + 4 # leave three spaces
    rc = 2 # start two from left wall
    while True:
        # print(rr,rc)
        # blow the rock
        dc = pattern[gascounter % len(pattern)]
        # print("blow by %d"%(dc))
        gascounter += 1
        if fits(rock, rr, rc + dc):
            rc += dc

        # drop the rock
        if fits(rock, rr - 1, rc):
            rr -= 1
        else:
            # land the rock
            draw(rock, rr, rc)
            break
    # period check
    if len(grid) > 100:
        hash = (tuple(grid[-100:]), turn % len(rocks), gascounter % len(pattern))
        if hash in last10state:
            print("repeat: %d has same result as at %d: %s"%(turn, last10state[hash], hash))
        else:
            last10state[hash] = turn
    if gascounter % (len(pattern)) == 0 and turn % len(rocks) == 0:
        print("period")
        print(turn)
        print(highestrow() + 1)
        last = now
        printgrid()


    # heightbyturn[turn + 1] = highestrow() + 1
    # if (turn + 1) % 10 == 0:
    #     heightby10 = heightbyturn[turn + 1]
    #     matches = True
    #     for i in range(10):
    #         if i > 1:
    #             print(turn + 1, i, heightby10 * i, heightbyturn[(turn + 1) // 10 * i])
    #         if heightby10 // 10 * i != heightbyturn[(turn + 1) // 10 * i]:
    #             matches = False
    #             break
    #     if matches:
    #         print("Period: %d turns -> %d lines"%((turn + 1) / 10, heightby10 / 10))
    #         raise "Error"


# find how tall the tower is
print(highestrow() + 1)
printgrid()

# 3152 is too low
# 3168

# part two:
# 1554553562591 is too high

# 1000000000000

# There's a repeating pattern: 5 rocks, and the pattern repeats, so it's period
# Each 50455 turns, generates 78435
# Next 50455 turns makes 156851, which is a bit less (78416 new)
# Next is 235246, 78395 new

# 128 -> 214
# 128 + 1700 (1828) -> 2856
# 128 + 2 * 1700 -> 5498

# After 1700, there are 2650
# After 3400 (two cycles) -> 5292
# Each new 1700 adds 2642
# 1828 looks like 128
# After 1000000000000 turns:
#  128 + (588235294 * 1700) + 72
# There are these number of rows
# 214 + 588235294 * 2642 + 108
# = 1554117646864 (too low, got the starting count wrong)
# = 1554117647070  (right!)

# 128 -> 200 (+72) adds 108 new lines

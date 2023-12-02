import re
import itertools
import functools
import math

rocks = [ # (row,column), but row goes 0 low
    # ####
    frozenset([
        (0,0),
        (0,1),
        (0,2),
        (0,3),
    ]),

    #
    # .#.
    # ###
    # .#.
    frozenset([
        (0,1),
        (1,0),
        (1,1),
        (1,2),
        (2,1),
    ]),

    #
    # ..#
    # ..#
    # ###
    frozenset([
        (0,0),
        (0,1),
        (0,2),
        (1,2),
        (2,2),
    ]),

    # #
    # #
    # #
    # #
    frozenset([
        (0,0),
        (1,0),
        (2,0),
        (3,0),
    ]),

    # ##
    # ##
    frozenset([
        (0,0),
        (0,1),
        (1,0),
        (1,1),
    ]),
]

with open("input.txt") as f:
    pattern = [-1 if c == '<' else 1 for c in f.readline().strip()]

width = 7
grid = []
def newrow():
    return [0 for i in range(width)]

# return highest row of rock
# the floor is -1
def highestrow():
    for i in range(len(grid)-1, -1, -1):
        if sum(grid[i]) > 0:
            return i
    return -1

print(highestrow())

# true if rock collides with rocks or walls in this position
def collision(rock, rr, rc):
    for (nr, nc) in rock: # node coordinates
        r = nr + rr
        c = nc + rc
        if c < 0 or c >= width:
            return True
        if r < 0:
            return True
        if r < len(grid) and grid[r][c]:
            return True
    return False

def fits(rock, rr, rc):
    return not collision(rock, rr, rc)

def draw(rock, rr, rc):
    if not fits(rock, rr, rc):
        raise Exception("Trying to draw but it doesn't fit: %d,%d"%(rr,rc))
    # print("drawing rock at %d,%d"%(rr,rc))
    for (nr, nc) in rock: # node coordinates
        r = nr + rr
        c = nc + rc
        while r >= len(grid):
            grid.append(newrow())
        if grid[r][c] > 0:
            print(r,c,grid[r][c])
            raise "Error"
        grid[r][c] = 1

def printgrid():
    print("------")
    for i in range(len(grid)-1,max(-1,len(grid)-5),-1):
        print("".join(["X" if x else " " for x in grid[i]]))
    print("...")
    for i in range(min(len(grid)-1,10),-1,-1):
        print("".join(["X" if x else " " for x in grid[i]]))

turns = len(pattern) * len(rocks)
turns = 1000000000000
print(turns)
print(1000000000000 // turns)
# turns = 1000000000000 % (len(pattern) * len(rocks))
print(turns)
gascounter = 0
last = 0
heightbyturn = {0: 0}
for turn in range(2022):
    printgrid()
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

# 3152 is too low
# 3168

# part two:
# 1554553562591 is too high

# 1000000000000

# There's a repeating pattern: 5 rocks, and the pattern repeats, so it's period
# Each 50455 turns, generates 78435
# Next 50455 turns makes 156851, which is a bit less (78416 new)
# Next is 235246, 78395 new

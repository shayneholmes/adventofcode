import re
import itertools
import functools
import math

def get(r,c):
    if r >= len(map):
        return " "
    if c >= len(map[r]):
        return " "
    return map[r][c]

path = None
map = []
with open("input.txt") as f:
    for line in f:
        m = re.match("[ .#]+$",line)
        if m:
            map.append(m.group(0))
            continue
        m = re.match("((\d+)[LR])+", line.strip())
        if m:
            path = line.strip()

print(map)
print(path)

facelen = 50
firstinrow = {}
firstincol = {}
lastinrow = {}
lastincol = {}
for r in range(len(map)):
    for c in range(len(map[0])):
        if get(r,c) != " ":
            firstinrow[r] = c
            break
for r in range(len(map)):
    for c in reversed(range(len(map[0]))):
        if get(r,c) != " ":
            lastinrow[r] = c
            break
for c in range(len(map[0])):
    for r in range(len(map)):
        if get(r,c) != " ":
            firstincol[c] = r
            break
for c in range(len(map[0])):
    for r in reversed(range(len(map))):
        if get(r,c) != " ":
            lastincol[c] = r
            break

print(firstinrow)

visited = {}

def char(dir):
    if dir == 0:
        return ">"
    if dir == 1:
        return "v"
    if dir == 2:
        return "<"
    if dir == 3:
        return "^"

def printmap():
    for r in range(len(map)):
        print("".join([char(visited[(r,c)]) if (r,c) in visited else get(r,c) for c in range(len(map[0]))]))

def moveAlongPath(r,c,dir,path):
    global map, visited
    visited = {}

    while path:
        m = re.match("^(\d+)(.*)$", path)
        if not m:
            break
        dist = int(m.group(1))
        path = m.group(2)
        print(path)
        print(r,c,dir)
        print("move",dist)

    #  XX
    #  X
    # XX
    # X

    #   aAbB  0
    #   AaBb  49
    #   cC    50
    #   Cc    99
    # dDeE    100
    # DdEe    149
    # fF      150
    # Ff      199
    # 045911
    #  90904
    #     09

        for i in range(dist):
            print(r,c,dir)
            # move direction
            (nur, nuc, nudir) = (r,c,dir)
            if dir == 0: #right
                nuc += 1
                if nuc > lastinrow[r]:
                    if r // facelen == 0: # B -> E
                        # (0,149) => (149,99)
                        # (49,149) => (100,99)
                        nur = 149 - r
                        nuc = 99
                        nudir = 2 # left
                    elif r // facelen == 1: # C -> B
                        # (50,50) -> (49,100)
                        # (99,50) -> (49,149)
                        nur = 49
                        nuc = r + 50
                        nudir = 3 # up
                    elif r // facelen == 2: # E-> B
                        # (100,100) -> (49,149)
                        # (149,100) -> (0,149)
                        nur = 149 - r
                        nuc = 149
                        nudir = 2 # left
                    elif r // facelen == 3: # F -> E
                        # (150,50) -> (149,50)
                        # (199,50) -> (149,99)
                        nur = 149
                        nuc = r - 100
                        nudir = 3 # up
            elif dir == 1: #down
                nur += 1
                if nur > lastincol[c]:
                    if c // facelen == 0: # F -> B (top)
                        nur = 0
                        nuc += 100
                    elif c // facelen == 1: # E -> F (right)
                        nuc = 49
                        nur = c + 100
                        nudir = 2 # left
                    elif c // facelen == 2: # B -> C (right)
                        nur = c - 50
                        nuc = 99
                        nudir = 2 # left
            elif dir == 2: #lef
                nuc -= 1
                if nuc < firstinrow[r]:
                    if r // facelen == 0: # A -> D (left)
                        nuc = 0
                        nur = 149 - r
                        nudir = 0
                    elif r // facelen == 1: # C -> D (top)
                        nudir = 1 # down
                        nur = 100
                        nuc = r - 50
                    elif r // facelen == 2: # D -> A (left)
                        nudir = 0
                        nuc = 50
                        nur = 149 - r
                    elif r // facelen == 3: # F -> A (top)
                        nur = 0
                        nuc = r - 100
                        nudir = 1 #down
            elif dir == 3: #up
                nur -= 1
                if nur < firstincol[c]:
                    if c // facelen == 0: # D -> C (left)
                        nudir = 0 # right
                        nuc = 50
                        nur = c + 50
                    elif c // facelen == 1: # A -> F (left)
                        nudir = 0 # right
                        nuc = 0
                        nur = c + 100
                    elif c // facelen == 2: # B -> F (bottom)
                        nudir = 3 #up
                        nur = 199
                        nuc = c - 100
            if get(nur, nuc) == "#":
                print("hit a wall")
                # hit a wall, stop here
                break
            if get(nur, nuc) != ".":
                raise Exception("Error: (%d,%d) -> %s"%(nur,nuc,get(nur,nuc)))
            (r,c,dir) = (nur,nuc,nudir)
            visited[(r,c)] = dir

        # turn
        m = re.match("^([LR])(.*)$", path)
        if not m:
            break
        rot = m.group(1)
        print("rotate", rot)
        path = m.group(2)
        if rot == "L":
            dir = (dir - 1) % 4
        else:
            dir = (dir + 1) % 4
        visited[(r,c)] = dir
    return (r,c,dir)

(r, c, dir) = moveAlongPath(0, firstinrow[0], 0, path)

# ## UNIT TESTS
# for r in range(len(map)):
#     c = lastinrow[r]
#     (nur, nuc, dir) = moveAlongPath(r, c, 0, "1R0R1R0R")
#     if (r,c,dir) != (nur,nuc,dir):
#         raise Exception("%d,%d -> %d,%d"%(r,c, nur,nuc))

# for r in range(len(map)):
#     c = firstinrow[r]
#     (nur, nuc, dir) = moveAlongPath(r, c, 2, "1R0R1R0R")
#     if (r,c,dir) != (nur,nuc,dir):
#         raise Exception("%d,%d -> %d,%d"%(r,c, nur,nuc))

# for c in range(150):
#     r = firstincol[c]
#     (nur, nuc, dir) = moveAlongPath(r, c, 3, "1R0R1R0R")
#     if (r,c,dir) != (nur,nuc,dir):
#         raise Exception("%d,%d -> %d,%d"%(r,c, nur,nuc))

# for c in range(150):
#     r = lastincol[c]
#     (nur, nuc, dir) = moveAlongPath(r, c, 1, "1R0R1R0R")
#     if (r,c,dir) != (nur,nuc,dir):
#         raise Exception("%d,%d -> %d,%d"%(r,c, nur,nuc))

# You begin the path in the leftmost open tile of the top row of tiles. Initially, you are facing to the right (from the perspective of how the map is drawn).

printmap()
print(r+1,c+1,dir)
print(1000 * (r+1) + 4 * (c+1) + dir)

# wrong guess: 109102

# part 2 wrong:
# 172099 : too high
# 104214 : too high
# 21302: too low (whoops, I forgot to replace the walls in the input after my tests)

import re
import itertools
import functools
import math

cubes = set()
with open("input.txt") as f:
    for line in f:
        m = re.match(r"(\d+),(\d+),(\d+)", line.strip())
        (x,y,z) = (int(m.group(1)),int(m.group(2)),int(m.group(3)))
        cubes.add((x,y,z))

xrange = (min([c[0] for c in cubes])-1 , max([c[0] for c in cubes])+1 )
yrange = (min([c[1] for c in cubes])-1 , max([c[1] for c in cubes])+1 )
zrange = (min([c[2] for c in cubes])-1 , max([c[2] for c in cubes])+1 )

print(xrange, yrange, zrange)

# flood from fill
tovisit = [(0,0,0)]
visited = set((0,0,0))
while tovisit:
    c = tovisit.pop()
    for d in [
            (1,0,0),
            (-1,0,0),
            (0,1,0),
            (0,-1,0),
            (0,0,1),
            (0,0,-1),
    ]:
        neighbor = (c[0]+d[0],c[1]+d[1],c[2]+d[2])
        if neighbor[0] < xrange[0]:
            continue
        if neighbor[0] > xrange[1]:
            continue
        if neighbor[1] < yrange[0]:
            continue
        if neighbor[1] > yrange[1]:
            continue
        if neighbor[2] < zrange[0]:
            continue
        if neighbor[2] > zrange[1]:
            continue
        if neighbor not in cubes and neighbor not in visited:
            visited.add(neighbor)
            tovisit.append(neighbor)

exposed = 0
for c in cubes:
    for d in [
            (1,0,0),
            (-1,0,0),
            (0,1,0),
            (0,-1,0),
            (0,0,1),
            (0,0,-1),
    ]:
        if (c[0]+d[0],c[1]+d[1],c[2]+d[2]) in visited:
            exposed += 1

print(exposed)


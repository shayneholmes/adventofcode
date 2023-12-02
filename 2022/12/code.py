elevations = []
with open('input.txt') as f:
    for line in f:
        line = line.strip()
        elevations.append([25 if ch == 'E' else 0 if ch == 'S' else ord(ch)-ord('a') for ch in line])

# hack: manually compute these
start = (0, 20)
goal = (43, 20)

def getElevation(pos):
    x = pos[0]
    y = pos[1]
    if x < 0 or x >= len(elevations[0]):
        return None
    if y < 0 or y >= len(elevations):
        return None
    return elevations[y][x]

def cost(pos):
    el = getElevation(pos)
    return abs(goal[0]-pos[0])**2 + abs(goal[1]-pos[1])**2 + 20 * abs(el-25)**2

# return neighbors in order of desirability: higher is better, less distance is better
def getNeighbors(pos):
    x = pos[0]
    y = pos[1]
    height = getElevation(pos)
    north = (x,y-1)
    south = (x,y+1)
    east = (x+1,y)
    west = (x-1,y)
    return sorted([i for i in [north, south, east, west] if (getElevation(i) != None and getElevation(i) <= height + 1)], key=cost)

def printmap():
    global path, elevations
    for i in range(len(elevations)):
        line = []
        for j in range(len(elevations[0])):
            if (j,i) in visited:
                line.append(".")
            else:
                line.append(chr(ord("a")+elevations[i][j]))
        print("".join(line))

# BFS to guarantee it's shortest
best = 9999999999
for startx in range(len(elevations[0])):
    for starty in range(len(elevations)):
        mystart=(startx,starty)
        if getElevation(mystart) != 0:
            continue

        path = {}
        visited = set()
        path[mystart] = mystart
        toVisit = [mystart] # queue
        while toVisit:
            # printmap()
            # print(toVisit)
            cur = toVisit.pop(0)
            visited.add(cur)
            if cur == goal:
                # Done!
                moves = 0
                i = goal
                visited = set()
                while i != mystart:
                    visited.add(i)
                    i = path[i]
                    moves += 1
                if moves < best:
                    best = moves
                    # printmap()
                break
            for n in reversed(getNeighbors(cur)):
                if n not in path:
                    path[n] = cur
                    toVisit.append(n)


print(best)

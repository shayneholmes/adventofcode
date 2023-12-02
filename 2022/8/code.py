rows = []
with open("input.txt") as f:
    for line in f:
        rows.append([int(i) for i in line.strip()])

# part 1

visible = [[False for i in j] for j in rows]
height = len(rows)
width = len(rows[0])
for i in range(height):
    # look from left
    maxSoFar = -1
    for j in range(width):
        v = rows[i][j]
        if v > maxSoFar: # tree is visible
            visible[i][j] = True
            maxSoFar = v
    # look from left
    maxSoFar = -1
    for j in reversed(range(width)):
        v = rows[i][j]
        if v > maxSoFar: # tree is visible
            visible[i][j] = True
            maxSoFar = v
for j in range(width):
    # look from top
    maxSoFar = -1
    for i in range(height):
        v = rows[i][j]
        if v > maxSoFar: # tree is visible
            visible[i][j] = True
            maxSoFar = v
    # look from left
    maxSoFar = -1
    for i in reversed(range(height)):
        v = rows[i][j]
        if v > maxSoFar: # tree is visible
            visible[i][j] = True
            maxSoFar = v

print(sum([1 if i else 0 for j in visible for i in j]))

def lookInDirection(y, x, dy, dx):
    startHeight = rows[y][x]
    distance = 0
    y += dy
    x += dx
    while 0 <= y and y < height and 0 <= x and x < width:
        distance += 1
        if rows[y][x] >= startHeight:
            # This is the last tree we can see, though we can see it.
            break
        y += dy
        x += dx
    return distance

best = 0
for i in range(height):
    for j in range(width):
        cur = \
            lookInDirection(i, j, -1, 0) * \
            lookInDirection(i, j,  1, 0) * \
            lookInDirection(i, j, 0, -1) * \
            lookInDirection(i, j, 0,  1)
        if cur > best:
            best = cur

print(best)

import re

def makenode(parent):
    return {'parent': parent, 'size': 0, 'children': {}}

def addchild(parent, name):
    child = makenode(parent)
    parent['children'][name] = child
    return child

root = makenode(None)
cur = root
with open("input.txt") as f:
    for line in f:
        line = line.strip()
        if line == "$ cd /":
            continue
        if line == "$ ls":
            continue
        if line == "$ cd ..":
            cur = cur['parent']
            continue
        m = re.match(r"\$ cd (\w+)", line)
        if m:
            cur = addchild(cur, m.group(1))
            continue
        m = re.match(r"dir (\w+)", line)
        if m:
            # Assume we'll catch this later
            continue
        m = re.match(r"(\d+) (\w+)", line)
        if m:
            size = int(m.group(1))
            # name = m.group(2)
            cur['size'] += size
            continue
print(root)

threshold = 100000
sum = 0

# Now the tree is built. Compute the size with a depth-first search
def computesize(node):
    global sum
    size = node['size']
    for i in node['children'].keys():
        size += computesize(node['children'][i])
    if size <= threshold:
        sum += size
    return size

updaterequiredspace = 30000000
totalspace = 70000000
spacetaken = computesize(root)
currentfreespace = totalspace - spacetaken
spacetofree = updaterequiredspace - currentfreespace

print(spacetaken)
print(spacetofree)

# Find the smallest dir that satisfies the condition
best = spacetaken # root node
def findsmallest(node):
    global best
    size = node['size']
    for i in node['children'].keys():
        size += findsmallest(node['children'][i])
    if size >= spacetofree and size < best:
        best = size
    return size

findsmallest(root)

print(best)

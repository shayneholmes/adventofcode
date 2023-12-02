import re
import itertools
import functools
import math
import queue

w = None
h = None
rbs = {}
lbs = {}
ubs = {}
dbs = {}
with open("input.txt") as f:
    r = -1
    for line in f:
        if r >= 0:
            if w == None:
                w = len(line.strip()) - 2
            if line[2] == "#": # last line:
                continue
            for (i, ch) in enumerate(line.strip()):
                if i == 0:
                    continue
                if i == 101:
                    continue
                c = i - 1
                if r not in rbs:
                    rbs[r] = set()
                if r not in lbs:
                    lbs[r] = set()
                if c not in ubs:
                    ubs[c] = set()
                if c not in dbs:
                    dbs[c] = set()
                dir = None
                if ch == ".":
                    continue
                if ch == ">": # right
                    rbs[r].add(c)
                    continue
                if ch == "<":
                    lbs[r].add(c)
                    continue
                if ch == "^":
                    ubs[c].add(r)
                    continue
                if ch == "v":
                    dbs[c].add(r)
                    continue
                raise Exception("error at %s, saw ch %s"%((r,c),ch))
        r += 1
    h = r

print(h,w)

goal = (35,99)

print(goal)

def dist(a,b): #manhattan dist
    return abs(a[0]-b[0]) + abs(a[1]-b[1])

def isopen(pos, turn):
    (r,c) = pos
    # check for right-going blizzards
    if r == -1 and c == 0:
        return True
    if r == h and c == w - 1:
        return True
    if r < 0 or r >= h:
        return False
    if c < 0 or c >= w:
        return False
    if (c + turn) % w in lbs[r]:
        return False
    if (c - turn) % w in rbs[r]:
        return False
    if (r + turn) % h in ubs[c]:
        return False
    if (r - turn) % h in dbs[c]:
        return False
    return True

def printmap(turn):
    global w, h
    for r in range(h):
        print("".join(["." if isopen((r,c),turn) else "X" for c in range(w)]))

best = 10000
i = 0
def go(turn,start,goal):
    done = set()
    global i
    pq = queue.PriorityQueue()
    pq.put((0,turn,start)) # pri, turn, pos
    while not pq.empty():
        (pri, turn, pos) = pq.get()
        i += 1
        if i % 10000 == 0:
            print("stack:",pq.qsize(),"turn:",turn,pos)
        if pos == goal:
            print("Found!", turn, pos)
            return turn
            continue
        for dir in [(1,0),(0,1),(0,0),(0,-1),(-1,0)]:
            cand = (pos[0]+dir[0],pos[1]+dir[1])
            # print("considering",cand)
            if isopen(cand, turn + 1):
                if (turn + 1,cand) in done:
                    continue
                pq.put((turn + dist(cand,goal), turn + 1, cand))
                done.add((turn + 1,cand))


leg1 = go(0,(-0,1),(h,w-1))
leg2 = go(leg1,(h,w-1),(-0,1))
leg3 = go(leg2,(-0,1),(h,w-1))
# wrong guesses: 920
# Found! 242 (35, 99)
# Found! 478 (-1, 0)
# Found! 720 (35, 99)

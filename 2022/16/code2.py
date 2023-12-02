import re
import itertools
import math

with open("input.txt") as f:
    valves = {}
    for line in f:
        line = line.strip()
        m = re.match(r"Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$", line)
        [valve, rate, neighbors] = [m.group(1), int(m.group(2)), m.group(3).split(", ")]
        valves[valve] = (rate, neighbors)

    valvesWithFlow = set([v for v in valves if valves[v][0] > 0])

    # first, find the shortest distance between all the nodes
    dists = {} # (src,dest) => number of turns
    for src in valves:
        left = set(valves.keys())
        states = [(src, 0)]
        while left:
            (cur, dist) = states.pop(0)
            if cur in left:
                left.remove(cur)
                dists[(src,cur)] = dist
                for n in valves[cur][1]:
                    if n in left:
                        states.append((n, dist+1))

    best = 0 # start with nothing
    bestsolution = None

    # When turns to get there is zero, we've done what we set out to do and turned on the tap
    # (We add a turn for turning on the thing)
    states = [('AA', 26, 0, set(), [], 1)] # dest and turns to get there (me), remaining moves, score, open, visited, playersafter

    def bound(state):
        moves = state[1]
        score = state[2]
        open = set(state[3])
        player = state[5]
        return score + (moves + 26 * player) * sum([valves[v][0] for v in valves]) / 2

    i = 0
    while states:
        state = states.pop()
        if bound(state) <= best:
            continue
        i += 1
        if i % 100000 == 0:
            print("---")
            print(state)
            print("bound: %d"%(best))
            print("states: %d"%(len(states)))
        cur = state[0]
        moves = state[1]
        score = state[2]
        open = set(state[3])
        progress = state[4]
        player = state[5]
        if score > best:
            best = score
            bestsolution = state
            print("New best: ", state)

        c = cur
        # look for options of where to go next.
        ni = 0
        leftToOpen = [v for v in valvesWithFlow if v not in open and dists[(c, v)] + 1 <= moves]
        if not leftToOpen:
            if player > 0:
                # Start new player at AA with 26 moves to go
                states.append(('AA', 26, score, open, progress, player - 1))
        for dest in leftToOpen:
            dist = dists[(c, dest)]
            nuopen = set(open)
            nuopen.add(dest)
            nuscore = (moves - dist - 1) * valves[dest][0] # turned on the tap, add the score going forward.
            states.append((dest, moves - dist - 1, score + nuscore, nuopen, progress + [dest], player))
            ni += 1
    print("%d cases"%(i))
    print(best)
    print(bestsolution)


#     # 1178 is too low
#     # 1243 is too low
#     # Part two:
#     # 1360 is wrong
#     # 1455 is wrong


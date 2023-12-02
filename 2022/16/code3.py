import re
import itertools
import functools
import math

with open("input.txt") as f:
    valves = {}
    for line in f:
        line = line.strip()
        m = re.match(r"Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)$", line)
        [valve, rate, neighbors] = [m.group(1), int(m.group(2)), m.group(3).split(", ")]
        valves[valve] = (rate, neighbors)

    valvesWithFlow = frozenset([v for v in valves if valves[v][0] > 0])

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

    # Recursive search with cache
    @functools.cache
    def bestscore(timeleft, position, opened, players):
        if timeleft < 0:
            raise "error!"
        if timeleft == 0:
            return (0, []) # no points!
        pointsthisturn = sum([valves[v][0] for v in opened])
        options = []
        # option 0: don't do anything, leave existing valves alone
        if players > 0:
            [points, path] = bestscore(26, 'AA', opened, players - 1)
            options.append(((timeleft - 26) * pointsthisturn + points, ["-- elephant --"] + path))
        else:
            options.append((timeleft * pointsthisturn, ['wait %d'%(timeleft)]))
        # option 1: turn on the valve this turn and get points for that
        for n in [v for v in valvesWithFlow if v not in opened]:
            turnsToAct = dists[(position, n)] + 1
            if turnsToAct > timeleft:
                continue
            [points, path] = bestscore(timeleft - turnsToAct, n, opened | frozenset([n]), players)
            options.append((pointsthisturn * turnsToAct + points, ["%s(%d)"%(n,turnsToAct)] + path))
        return max(options)

    print(bestscore(26, 'AA', frozenset(), 1))



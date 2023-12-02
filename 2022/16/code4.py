import re
import itertools
import functools
import math

with open("input.txt") as f:
    maxtime = 30
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

    # Recursive search, pass down score so we can save all results at the end
    results = {} # (opened) -> score
    def compute(timeleft, position, opened, score):
        if timeleft < 0:
            raise "error!"
        pointsthisturn = sum([valves[v][0] for v in opened])
        options = []
        # option 0: don't do anything, leave existing valves alone
        endscore = score + timeleft * pointsthisturn
        if opened in results:
            results[opened] = max(results[opened], endscore)
        else:
            results[opened] = endscore
        # print(opened, results[opened])
        # option 1: go somewhere and turn on the valve there
        for n in [v for v in valvesWithFlow if v not in opened]:
            turnsToAct = dists[(position, n)] + 1 # extra turn to turn on the valve
            if turnsToAct > timeleft:
                # can't make it there in time
                continue
            compute(timeleft - turnsToAct, n, opened | frozenset([n]), score + pointsthisturn * turnsToAct)

    compute(26, 'AA', frozenset(), 0)
    best = 0
    bestsolution = None
    for me in results:
        for el in results:
            if not me & el: # No overlap, awesome
                score = results[me] + results[el]
                if score > best:
                    best = score
                    bestsolution = (me, el)

    print(best, bestsolution)



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
    cur = 'AA'
    best = 0 # guess
    bestsolution = None
    states = [(('AA', 'AA'), 26, 0, set(), {})] # node, remaining moves, score, open, visited
    # DFS: Try moves
    def bound(state): # best we can get from here
        # we can open a valve every two moves if we're lucky, and 25 at the most, then descending
        moves = state[1]
        score = state[2]
        open = state[3]
        visited = state[4]
        # we might get moves more valves open
        return score + min(moves,(14 - len(state[3]))) * 25 * moves / 2
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
        # print(state)
        (cur, el) = state[0]
        moves = state[1]
        score = state[2]
        open = state[3]
        visited = state[4]
        if score > best:
            best = score
            bestsolution = state
            print("New best: ", state)
        if moves == 0:
            continue
        for n in [v[1] for v in sorted([(valves[v][0], v) for v in valves[cur][1]])]:
            # I move to a neighbor
            nuvisited = dict(visited)
            if n in visited and visited[n] > 3:
                continue
            if n not in visited:
                nuvisited[n] = 0
            nuvisited[n] += 1
            for eln in [v[1] for v in sorted([(valves[v][0], v) for v in valves[el][1]])]:
                # el moves to a neighbor
                nunuvisited = dict(nuvisited)
                if eln in nuvisited and nuvisited[eln] > 3:
                    continue
                if eln not in nuvisited:
                    nunuvisited[eln] = 0
                nunuvisited[eln] += 1
                states.append(((n, eln), moves-1, score, open, nunuvisited))
            if valves[el][0] > 0 and not el in open:
                # el opens a valve
                newopen = set(open)
                newopen.add(el)
                states.append(((n, el), moves-1, score+valves[el][0] * (moves - 1), newopen, nuvisited))
        if valves[cur][0] > 0 and not cur in open:
            # I open a valve
            newopen = set(open)
            newopen.add(cur)
            for eln in valves[el][1]:
                # el moves to a neighbor
                nunuvisited = dict(visited)
                if eln in nuvisited and nuvisited[eln] > 3:
                    continue
                if eln not in nunuvisited:
                    nunuvisited[eln] = 0
                nunuvisited[eln] += 1
                states.append(((cur, eln), moves-1, score+valves[cur][0] * (moves - 1), newopen, nunuvisited))
            if valves[el][0] > 0 and not el in open:
                # el opens a valve
                newnewopen = set(newopen)
                newnewopen.add(el)
                states.append(((n, el), moves-1, score+(valves[cur][0] + valves[el][0]) * (moves - 1), newnewopen, nuvisited))
    print(best)
    print(bestsolution)


    # 1178 is too low
    # 1243 is too low
    # Part two:
    # 1360 is wrong

import re
import itertools
import functools
import math

turns = 32
plans = {}
with open("input.txt") as f:
    for line in f:
        m = re.match(r"Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.", line.strip())
        (id, orecost, clayinore, obsidianore, obsidianclay, geodeore, geodeobsidian) = (int(m.group(1)),int(m.group(2)),int(m.group(3)),int(m.group(4)),int(m.group(5)),int(m.group(6)),int(m.group(7)))
        if id <= 3:
            plans[id] = (id, orecost, clayinore, obsidianore, obsidianclay, geodeore, geodeobsidian)

def turnsToBuild(amt, perturn):
    if amt <= 0:
        return 1 # we already have enough, just one turn to build
    turns = amt // perturn + 1 # at least this many turns (including one to build)
    if amt % perturn > 0:
        turns += 1 # will take an extra turn to get enough
    return turns

i = 0
best = {}
@functools.cache
def mostgeodes(plan, timeleft, state):
    global i, best
    i += 1
    if i % 1000000 == 0:
        print(timeleft, state)

    (id, orecost, clayinore, obsidianore, obsidianclay, geodeore, geodeobsidian) = plan
    (ore, clay, obsidian, geodes, orebots, claybots, obsidianbots, geodebots) = state
    options = []

    if timeleft < 2:
        res = geodes + timeleft * geodebots
        if plan not in best or best[plan] < res:
            best[plan] = res
        return res

    # bounds check
    if plan in best and geodes + timeleft * geodebots + timeleft * timeleft // 2 < best[plan]:
        # short-circuit
        return best[plan]

    if orebots > 0 and obsidianbots > 0:
        # save up and build a geode bot
        turns = max(turnsToBuild(geodeore - ore, orebots), turnsToBuild(geodeobsidian - obsidian, obsidianbots))
        if timeleft >= turns:
            options.append((mostgeodes(plan, timeleft - turns, (ore + turns * orebots - geodeore, clay + turns * claybots, obsidian + turns * obsidianbots - geodeobsidian, geodes + turns * geodebots, orebots, claybots, obsidianbots, geodebots + 1)), "build a geode bot"))
    if orebots > 0 and claybots > 0 and obsidianbots < geodeobsidian:
        # save up and build an obsidian bot
        turns = max(turnsToBuild(obsidianore - ore, orebots), turnsToBuild(obsidianclay - clay, claybots))
        if timeleft >= turns:
            options.append((mostgeodes(plan, timeleft - turns, (ore + turns * orebots - obsidianore, clay + turns * claybots - obsidianclay, obsidian + turns * obsidianbots, geodes + turns * geodebots, orebots, claybots, obsidianbots + 1, geodebots)), "build an obsidian bot"))
    if orebots > 0 and claybots < obsidianclay:
        # save up and build a clay bot
        turns = turnsToBuild(clayinore - ore, orebots)
        if timeleft >= turns:
            options.append((mostgeodes(plan, timeleft - turns, (ore + turns * orebots - clayinore, clay + turns * claybots, obsidian + turns * obsidianbots, geodes + turns * geodebots, orebots, claybots + 1, obsidianbots, geodebots)), "build a clay bot"))
    if orebots > 0 and orebots < max(clayinore, obsidianore, geodeore):
        # save up and build an ore bot
        turns = turnsToBuild(orecost - ore, orebots)
        if timeleft >= turns:
            options.append((mostgeodes(plan, timeleft - turns, (ore + turns * orebots - orecost, clay + turns * claybots, obsidian + turns * obsidianbots, geodes + turns * geodebots, orebots + 1, claybots, obsidianbots, geodebots)), "build an ore bot"))
    # do nothing for the rest of the game
    turns = timeleft
    options.append((geodes + turns * geodebots, "wait"))
    # if max(options)[0] > 9:
    #     print(max(options), timeleft, state)
    res = max(options)[0]
    if plan not in best or best[plan] < res:
        best[plan] = res
    return res

sum = 0
for pi in plans:
    geodes = mostgeodes(plans[pi], turns, (0, 0, 0, 0, 1, 0, 0, 0))
    quality = geodes * pi
    sum += quality
    print(pi, geodes, quality)

print(sum)


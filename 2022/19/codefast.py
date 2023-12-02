import re
import itertools
import functools
import math

turns = 24
plans = {}
with open("input.txt") as f:
    for line in f:
        m = re.match(r"Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.", line.strip())
        (id, orecost, clayinore, obsidianore, obsidianclay, geodeore, geodeobsidian) = (int(m.group(1)),int(m.group(2)),int(m.group(3)),int(m.group(4)),int(m.group(5)),int(m.group(6)),int(m.group(7)))
        plans[id] = (id, orecost, clayinore, obsidianore, obsidianclay, geodeore, geodeobsidian) = (int(m.group(1)),int(m.group(2)),int(m.group(3)),int(m.group(4)),int(m.group(5)),int(m.group(6)),int(m.group(7)))

i = 0
@functools.cache
def mostgeodes(plan, timeleft, state):
    global i
    i += 1
    if i % 100000 == 0:
        print(state)
    (id, orecost, clayinore, obsidianore, obsidianclay, geodeore, geodeobsidian) = plan
    (ore, clay, obsidian, geodes, orebots, claybots, obsidianbots, geodebots) = state
    if timeleft == 0:
        if geodes > 5:
            print(geodes, state)
        return geodes
    options = []
    if orebots > 0 and obsidianbots > 0:
        # save up and build a geode bot
        turns = math.ceil(max(0, float(geodeore - ore) / orebots, float(geodeobsidian - obsidian) / obsidianbots)) + 1
        if timeleft >= turns:
            options.append(mostgeodes(plan, timeleft - turns, (ore + turns * orebots - geodeore, clay + turns * claybots, obsidian + turns * obsidianbots - geodeobsidian, geodes + turns * geodebots, orebots, claybots, obsidianbots, geodebots + 1)))
    if orebots > 0 and claybots > 0:
        # save up and build an obsidian bot
        turns = math.ceil(max(0, float(obsidianore - ore) / orebots, float(obsidianclay - clay) / claybots)) + 1
        if timeleft >= turns:
            options.append(mostgeodes(plan, timeleft - turns, (ore + turns * orebots - obsidianore, clay + turns * claybots - obsidianclay, obsidian + turns * obsidianbots, geodes + turns * geodebots, orebots, claybots, obsidianbots + 1, geodebots)))
    if orebots > 0:
        # save up and build a clay bot
        turns = math.ceil(max(0, float(clayinore - ore) / orebots)) + 1
        if timeleft >= turns:
            options.append(mostgeodes(plan, timeleft - turns, (ore + turns * orebots - clayinore, clay + turns * claybots, obsidian + turns * obsidianbots, geodes + turns * geodebots, orebots, claybots + 1, obsidianbots, geodebots)))
    if orebots > 0:
        # save up and build an ore bot
        turns = math.ceil(max(0, float(orecost - ore) / orebots)) + 1
        if timeleft >= turns:
            options.append(mostgeodes(plan, timeleft - turns, (ore + turns * orebots - orecost, clay + turns * claybots, obsidian + turns * obsidianbots, geodes + turns * geodebots, orebots + 1, claybots, obsidianbots, geodebots)))
    if timeleft > 0:
        # do nothing for the rest of the game
        turns = timeleft
        options.append(geodes + turns * geodebots)
    return max(options)

sum = 0
for pi in plans:
    geodes = mostgeodes(plans[pi], turns, (0, 0, 0, 0, 1, 0, 0, 0))
    quality = geodes * pi
    sum += quality
    print(pi, geodes, quality)

print(sum)

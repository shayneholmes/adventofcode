import re
import itertools
import functools
import math

elves = set()
with open("input.txt") as f:
    r = 0
    for line in f:
        c = 0
        for ch in line.strip():
            if ch == "#":
                elves.add((r,c))
            c += 1
        r += 1

def getbounds():
    global elves
    rlo = 9999
    rhi = -9999
    clo = 9999
    chi = -9999
    for (r,c) in elves:
        if r < rlo:
            rlo = r
        if r > rhi:
            rhi = r
        if c < clo:
            clo = c
        if c > chi:
            chi = c
    return (rlo,clo,rhi,chi)

def printelves():
    print("--")
    global elves
    (rlo,clo,rhi,chi) = getbounds()
    for r in range(rlo,rhi+1):
        print("".join(["#" if (r,c) in elves else "." for c in range(clo,chi+1)]))

round = 0
while True:
    # printelves()
    # scope out spots
    proposals = set()
    dests = set()
    aborted = set()
    for (r,c) in elves:
        n = (r-1,c)
        ne = (r-1,c+1)
        nw = (r-1,c-1)
        e = (r,c+1)
        w = (r,c-1)
        s = (r+1,c)
        se = (r+1,c+1)
        sw = (r+1,c-1)
        hasn = n in elves
        hasne = ne in elves
        hasnw = nw in elves
        hase = e in elves
        hasw = w in elves
        hass = s in elves
        hasse = se in elves
        hassw = sw in elves
        if not hasn and not hasne and not hasnw and not hase and not hasw and not hass and not hasse and not hassw:
            # no motion
            continue
        myproposals = set() # (pri,r,c)
        if not hasn and not hasne and not hasnw:
            myproposals.add(((0 - round) % 4,n))
            # print("north:",(r,c),((0 - round) % 4,n))
        if not hass and not hasse and not hassw:
            myproposals.add(((1 - round) % 4,s))
            # print("south:",(r,c),((1 - round) % 4,s))
        if not hasw and not hasnw and not hassw:
            myproposals.add(((2 - round) % 4,w))
            # print("west:",(r,c),((2 - round) % 4,w))
        if not hase and not hasne and not hasse:
            myproposals.add(((3 - round) % 4,e))
            # print("east:",(r,c),((3 - round) % 4,e))
        if not myproposals: # no options
            continue
        (pri,nu) = sorted(myproposals)[0]
        # print("chose",pri,nu)
        if nu not in dests:
            proposals.add(((r,c),nu))
            dests.add(nu)
        else:
            aborted.add(nu)
    # Make the moves
    madeamove = False
    for (old,nu) in proposals:
        if nu not in aborted:
            madeamove = True
            elves.remove(old)
            elves.add(nu)
    round += 1
    if not madeamove:
        print(round)
        break

# printelves()
# (rlo,clo,rhi,chi) = getbounds()
# print((rhi - rlo + 1) * (chi - clo + 1) - len(elves))

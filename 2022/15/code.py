import re

def union(ranges):
    res = []
    last = None
    for r in sorted(ranges):
        # print("---")
        # print(res)
        # print(last)
        # print(r)
        if last == None:
            last = r
            continue
        if last[1] < r[0] - 1: # no intersection
            # print("no intersection, saving", last)
            res.append(last)
            last = r
        else:
            if last[1] >= r[1]: #totally swallowed
                continue
            last = (last[0],r[1]) # merge
    res.append(last)
    return res


with open("input.txt") as f:
    sensors = set()
    beacons = set()
    for line in f:
        m = re.match(r"^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$", line.strip())
        sx = int(m.group(1))
        sy = int(m.group(2))
        bx = int(m.group(3))
        by = int(m.group(4))
        dist = abs(sx-bx)+abs(sy-by) # manhattan
        sensors.add((sx,sy,dist))
        beacons.add((bx,by))

    # step 1: Find all places on y=2000000 that can't have beacons
    for target in range(3100000,4000000):
        if target % 100000 == 0:
            print(target)
        ranges = []

        for s in sensors:
            (x,y,dist) = s
            # how far is the sensor away from this line?
            dy = abs(y - target)
            remainingdist = dist - dy
            if remainingdist < 0:
                # out of range of the target
                continue
            ranges.append((x-remainingdist,x+remainingdist))

        numbeaconsOnTarget = sum([1 for i in beacons if i[1] == target])
        ans = (union(sorted(ranges)))
        if len(ans) > 1:
            print(target, ans)

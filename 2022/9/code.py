def getdelta(dir): # [dx, dy]
    if dir == "D":
        return [0, -1]
    if dir == "U":
        return [0, 1]
    if dir == "R":
        return [1, 0]
    if dir == "L":
        return [-1, 0]
    raise "bad dir %s"%(dir)

with open("input.txt") as f:
    knots = [(0,0) for i in range(10)]
    tpositions = set() # set of tuples of where the tail has been
    for cmd in f:
        [dir, num] = cmd.split(" ")
        num = int(num)
        [dx, dy] = getdelta(dir)
        for i in range(num):
            # move head
            knots[0] = (knots[0][0] + dx, knots[0][1] + dy)
            for i in range(1, len(knots)):
                # catch knots up
                (hx,hy) = knots[i-1]
                (tx,ty) = knots[i]
                if abs(hx-tx) > 1:
                    tx = (hx+tx)//2
                    ty = hy
                if abs(hy-ty) > 1:
                    ty = (hy+ty)//2
                    tx = hx
                knots[i] = (tx, ty)
            tpositions.add(knots[len(knots)-1])

print((tpositions))
print(len(tpositions))

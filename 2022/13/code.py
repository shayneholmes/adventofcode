import math
import functools

# return negative if l is lower than r, 0 if same.
def compare(l, r):
    # Simple case: Two numbers
    if isinstance(l, int) and isinstance(r, int):
        return l - r
    # Convert
    if isinstance(l, int):
        l = [l]
    if isinstance(r, int):
        r = [r]
    if isinstance(l, list) and isinstance(r, list):
        # compare each element of the list
        for i in range(min(len(l), len(r))):
            res = compare(l[i], r[i])
            if res != 0:
                return res
        return len(l) - len(r)


divider1 = [[2]]
divider2 = [[6]]
packets = [divider1, divider2] # start with divider packets
with open("input.txt") as f:
    index = 0 # start at 1, we'll increment in a second
    while True:
        try:
            index += 1
            l = eval(f.readline().strip())
            r = eval(f.readline().strip())
            packets.append(l)
            packets.append(r)
            f.readline()
        except:
            break

packets = sorted(packets, key=functools.cmp_to_key(compare))
print(math.prod([i+1 for [i,v] in enumerate(packets) if v == divider1 or v == divider2]))

# 20972 is too low, because it's zero-indexed D:

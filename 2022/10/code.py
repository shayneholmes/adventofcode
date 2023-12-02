import re

xs = [0]
pixels = [False]
x = 1
with open("input.txt") as f:
    def addpixel():
        global xs, x, pixels
        pos = (len(xs)-1) % 40
        pixels.append(abs(pos-x) <= 1)
        xs.append(x)
    for cmd in f:
        cmd = cmd.strip()
        if cmd == "noop":
            addpixel()
            continue
        m = re.match(r"^addx (-?\d+)$", cmd)
        if m:
            addpixel()
            addpixel()
            dx = int(m.group(1))
            x += dx
            continue
        raise "error"

print(xs)

print(sum([i*v for [i,v] in enumerate(xs) if i % 40 == 20]))
print(str("".join([(("#" if v else ".") + ("\n" if i % 40 == 0 else "")) for [i,v] in enumerate(pixels)])))


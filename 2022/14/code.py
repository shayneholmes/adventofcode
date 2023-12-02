
xmin = -1000
xmax = 5000
ymin = 0
ymax = 172

arrwidth = xmax - xmin + 1
arrheight = ymax - ymin + 1
print(arrheight,arrwidth)

myarray = [[0 for i in range(arrwidth)] for j in range(arrheight)]

def print_path():
    global myarray
    def ch(i):
        if i == 0:
            return " "
        if i == 1: # rock
            return "#"
        if i == 2: # sand
            return "o"
        return "X"
    for r in myarray:
        print("".join([ch(i) for i in r]))

def get(x,y):
    global myarray
    if x < xmin:
        return None
    if x > xmax:
        return None
    if y < ymin:
        return None
    if y > ymax:
        return None
    return myarray[y-ymin][x-xmin]

def write(x,y,v):
    global myarray
    # assume that callers provide in-bound values when they write
    # print(y-ymin, x-xmin, v)
    myarray[y-ymin][x-xmin] = v

with open("input.txt") as f:
    for path in f:
        coords = path.strip().split(" -> ")
        xi, yi = None, None
        for c in coords:
            [x, y] = [int(i) for i in c.split(",")]
            if not xi:
                write(x,y,1)
                xi, yi = x, y
                continue
            if x > xi:
                for i in range(xi, x+1):
                    write(i,y,1)
                xi, yi = x, y
                continue
            if y > yi:
                for i in range(yi, y+1):
                    write(x,i,1)
                xi, yi = x, y
                continue
            if x < xi:
                for i in range(x, xi+1):
                    write(i,y,1)
                xi, yi = x, y
                continue
            if y < yi:
                for i in range(y, yi+1):
                    write(x,i,1)
                xi, yi = x, y
                continue
            raise "ERRORR"

# drop grains
grains = 0
while True:
    y = 0
    x = 500
    while True:
        if y == ymax:
            # floor
            break
        if get(x,y) == None:
            # off the map
            raise f"ERRROR at %d, %d"%(x,y)
        # try down
        if get(x,y+1) == 0:
            y = y+1
            continue
        # try left
        if get(x-1,y+1) == 0:
            x = x-1
            y = y+1
            continue
        # right
        if get(x+1,y+1) == 0:
            x = x+1
            y = y+1
            continue
        # stuck!
        break
    if get(x,y) == None:
        # off the map
        raise "Off the map!"
    write(x,y,2)
    grains += 1
    if grains % 10000 == 0:
        print_path()
    if y == 0: #source is plugged, stop
        break
print(grains)



def priority(ch):
    if ord('a') <= ord(ch) and ord('z') >= ord(ch):
        return ord(ch) - ord('a') + 1
    return ord(ch) - ord('A') + 27

# part 1
with open('input.txt') as f:
    pri = 0
    for line in f:
        s = len(line)//2
        left = set(line[:s])
        right = set(line[s:])
        pri += priority(list(left&right)[0])
    print(pri)

# part 2
with open('input.txt') as f:
    pri = 0
    try:
        while True:
            line1 = set(f.readline()[:-1])
            line2 = set(f.readline()[:-1])
            line3 = set(f.readline()[:-1])
            print(line1&line2&line3)
            pri += priority(list(line1&line2&line3)[0])
    except:
        pass
    print(pri)

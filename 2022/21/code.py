import re
import itertools
import functools
import math

monkeys = {}
results = {}
with open("input.txt") as f:
    for line in f:
        m = re.match("([a-z]+): (\d+)", line.strip())
        if m:
            mon = m.group(1)
            if mon == "humn":
                continue
            num = int(m.group(2))
            results[mon] = num
            continue
        m = re.match("([a-z]+): (\w+) (.) (\w+)", line.strip())
        mon = m.group(1)
        a = (m.group(2))
        op = (m.group(3))
        b = (m.group(4))
        monkeys[mon] = (op, a, b)

@functools.cache
def get(id):
    if id in results:
        return results[id]
    if id == "humn":
        return "hi"
    (op, a, b) = monkeys[id]
    ares = get(a)
    bres = get(b)
    if isinstance(ares, int) and isinstance(bres, int):
        if op == "*":
            return ares * bres
        if op == "+":
            return ares + bres
        if op == "-":
            return ares - bres
        if op == "/":
            return ares // bres
    return "(%s %s %s)"%(ares,op,bres)

@functools.cache
def makeitequal(id, eq):
    if id in results:
        return results[id]
    if id == "humn":
        print(eq)
        return eq
    (op, a, b) = monkeys[id]
    ares = get(a)
    bres = get(b)
    if isinstance(ares, int) and isinstance(bres, int):
        if op == "*":
            return ares * bres
        if op == "+":
            return ares + bres
        if op == "-":
            return ares - bres
        if op == "/":
            return ares // bres
    if isinstance(ares, int):
        # find b
        if op == "*":
            return makeitequal(b, eq // ares)
        if op == "+":
            return makeitequal(b, eq - ares)
        if op == "-":
            return makeitequal(b, ares - eq)
        if op == "/":
            return makeitequal(b, eq * ares)
    if isinstance(bres, int):
        # find a
        if op == "*": # a * b = eq
            return makeitequal(a, eq // bres)
        if op == "+": # a + b = eq
            return makeitequal(a, eq - bres)
        if op == "-": # a - b = eq
            return makeitequal(a, eq + bres)
        if op == "/": # a / b = eq
            return makeitequal(a, eq * bres)
    return "(%s %s %s)"%(ares,op,bres)


print(get('fglq'))
makeitequal('fglq',42130890593816)

# fzbp: 42130890593816

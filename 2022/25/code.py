import re
import itertools
import collections
import functools
import math

base = 5

def dec2snafu(dec):
    res = collections.deque()
    value = 1
    while dec:
        dig = dec % (value * base)
        dec -= dig
        dec //= base
        if dig > 2: # We can't represent 3 or 4
            dig -= base
            dec += 1 # need to up the next
        res.appendleft(dig2ch(dig))
    return "".join(res)

def dig2ch(dig):
    if dig == 2:
        return "2"
    if dig == 1:
        return "1"
    if dig == 0:
        return "0"
    if dig == -1:
        return "-"
    if dig == -2:
        return "="
    raise Exception("unknown digit %d"%(dig))

def ch2dig(ch):
    if ch == "1":
        return 1
    if ch == "2":
        return 2
    if ch == "-":
        return -1
    if ch == "=":
        return -2
    if ch == "0":
        return 0
    raise Exception("unknown digit %s"%(ch))

def snafu2dec(snafu):
    snafu = list(snafu)
    res = 0
    value = 1
    while snafu:
        ch = snafu.pop()
        res += ch2dig(ch) * value
        value *= base
    return res

print(snafu2dec("1121-1110-1=0"))
print(dec2snafu(15))

with open("input.txt") as f:
    sum = 0
    for line in f:
        sum += snafu2dec(line.strip())
    print(dec2snafu(sum))

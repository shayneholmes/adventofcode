import re
import itertools
import functools
import math

l = {}
key = 811589153
zero = None
with open("input.txt") as f:
    i = 0
    last = None
    for line in f:
        n = int(line.strip()) * 811589153
        node = {'value': n, 'last': last, 'next': None}
        if n == 0:
            zero = node
        if last:
            last['next'] = node
        l[i] = node
        last = node
        i += 1
    # circular list
    l[i-1]['next'] = l[0]
    l[0]['last'] = l[i-1]

length = len(l)
for x in range(10):
    for i in l:
        node = l[i]
        distance = node['value']
        distance %= (length - 1)
        # print('---')
        # print(distance)
        # potential optimization
        # if distance > length // 2:
        #     distance -= length

        # cut it out
        prev = node['last']
        next = node['next']
        prev['next'] = next
        next['last'] = prev

        # find the new destination
        for j in range(distance):
            prev = prev['next']

        # insert it after prev
        next = prev['next']
        prev['next'] = node
        next['last'] = node
        node['next'] = next
        node['last'] = prev

        # cur = zero
        # print('----')
        # while True:
        #     print(cur['value'])
        #     cur = cur['next']
        #     if cur == zero:
        #         break



# get the numbers
sum = 0
for i in [1000,2000,3000]:
    cur = zero
    for j in range(i):
        cur = cur['next']
    print(cur['value'])
    sum += cur['value']

print(sum)

import math


monkeys = [
    {
        # monkey 0
        'items': [62, 92, 50, 63, 62, 93, 73, 50],
        'op': lambda a: a * 7,
        'test': lambda a: a % 2 == 0,
        'iftrue': 7,
        'iffalse': 1,
        'inspections': 0,
    },
    {
        # monkey 1
        'items': [51, 97, 74, 84, 99],
        'op': lambda a: a + 3,
        'test': lambda a: a % 7 == 0,
        'iftrue': 2,
        'iffalse': 4,
        'inspections': 0,
    },
    {
        # monkey 2
        'items': [98, 86, 62, 76, 51, 81, 95],
        'op': lambda a: a + 4,
        'test': lambda a: a % 13 == 0,
        'iftrue': 5,
        'iffalse': 4,
        'inspections': 0,
    },
    {
        # monkey 3
        'items': [53, 95, 50, 85, 83, 72],
        'op': lambda a: a + 5,
        'test': lambda a: a % 19 == 0,
        'iftrue': 6,
        'iffalse': 0,
        'inspections': 0,
    },
    {
        # monkey 4
        'items': [59, 60, 63, 71],
        'op': lambda a: a * 5,
        'test': lambda a: a % 11 == 0,
        'iftrue': 5,
        'iffalse': 3,
        'inspections': 0,
    },
    {
        # monkey 5
        'items': [92, 65],
        'op': lambda a: a * a,
        'test': lambda a: a % 5 == 0,
        'iftrue': 6,
        'iffalse': 3,
        'inspections': 0,
    },
    {
        # monkey 6
        'items': [78],
        'op': lambda a: a + 8,
        'test': lambda a: a % 3 == 0,
        'iftrue': 0,
        'iffalse': 7,
        'inspections': 0,
    },
    {
        # monkey 7
        'items': [84, 93, 54],
        'op': lambda a: a + 1,
        'test': lambda a: a % 17 == 0,
        'iftrue': 2,
        'iffalse': 1,
        'inspections': 0,
    }
]

# monkeys = [
#     {
#         # monkey 0
#         'items': [79,98],
#         'op': lambda a: a * 19,
#         'test': lambda a: a % 23 == 0,
#         'iftrue': 2,
#         'iffalse': 3,
#         'inspections': 0,
#     },
#     {
#         # monkey 1
#         'items': [54, 65, 75, 74],
#         'op': lambda a: a + 6,
#         'test': lambda a: a % 19 == 0,
#         'iftrue': 2,
#         'iffalse': 0,
#         'inspections': 0,
#     },
#     {
#         # monkey 2
#         'items': [79, 60, 97],
#         'op': lambda a: a * a,
#         'test': lambda a: a % 13 == 0,
#         'iftrue': 1,
#         'iffalse': 3,
#         'inspections': 0,
#     },
#     {
#         # monkey 3
#         'items': [74],
#         'op': lambda a: a + 3,
#         'test': lambda a: a % 17 == 0,
#         'iftrue': 0,
#         'iffalse': 1,
#         'inspections': 0,
#     },
# ]

for round in range(10000):
    for mi, m in enumerate(monkeys):
        # print('Monkey %d'%(mi))
        # print(' Has %s'%(m['items']))
        m['inspections'] += len(m['items'])
        for i in m['items']:
            # print(' Monkey inspects an item with worry level %d'%(i))
            i = m['op'](i) # do operation
            # print(' New worry level is %d'%(i))
            i = i % 9699690 # Drop by this instead of a third (it's all the monkey levels togethr)
            # print(' Worry level decreases to %d'%(i))
            target = m['iftrue'] if m['test'](i) else m['iffalse']
            # print(' Throws item to %d'%(target))
            monkeys[target]['items'].append(i)
        m['items'] = []
    # print([m['items'] for m in monkeys])

print(math.prod(list(reversed(sorted([m['inspections'] for m in monkeys])))[:2]))

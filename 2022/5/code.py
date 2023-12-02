stacks = []
stacks.append("")
stacks.append("VCDRZGBW")
stacks.append("GWFCBSTV")
stacks.append("CBSNW")
stacks.append("QGMNJVCP")
stacks.append("TSLFDHB")
stacks.append("JVTWMN")
stacks.append("PFLCSTG")
stacks.append("BDZ")
stacks.append("MNZW")

with open("input.txt") as f:
    for instruction in f:
        words = instruction.split(" ")
        print(stacks)
        print(words)
        no = int(words[1])
        src = int(words[3])
        dest = int(words[5])
        for i in range(no):
            stacks[dest] += stacks[src][-1]
            stacks[src] = stacks[src][:-1]
    print("".join(str([s[len(s)-1] for s in stacks[1:]])))

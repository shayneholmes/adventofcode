with open("input.txt") as f:
    for line in f:
        size = 14
        ringbuff = [" " for i in range(size)]
        i = 0;
        for ch in line:
            i += 1
            ringbuff[i%size] = ch
            # print(ringbuff)
            if i <= size:
                # not enough signal yet
                continue
            d = dict()
            for j in range(size):
                d[ringbuff[j]] = 1
            if len(d) == size:
                print(i)
                break

def take_input():
    n, c = input().split()
    n = int(n)
    c = int(c)
    interval = []
    for j in range(n):
        l, r = input().split()
        l = int(l)
        r = int(r)
        interval.append([l, r])
    new_interval = total_new_intervals(interval, c)
    total_intervals = new_interval + n
    return total_intervals


def total_new_intervals(interval, c):
    h_map = {}
    for i in interval:
        for j in range((i[0]+1), i[1]):
            # print(j)
            if j in h_map:
                h_map[j] = h_map[j]+1
            else:
                h_map[j] = 1
    l = list(h_map.values())
    l.sort()
    l.reverse()
    if c > len(l):
        new_interval = sum(l)
    else:
        new_interval = sum(l[:c])
    return new_interval


test = int(input())
for i in range(test):
    output = take_input()
    print(f"Case #{i+1}: {output}")

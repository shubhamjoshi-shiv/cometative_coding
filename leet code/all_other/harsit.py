def is_similar(e1, e2):
    e1_l = len(e1)
    e2_l = len(e2)
    diff = abs(e1_l-e2_l)

    if diff > 1:
        return False
    elif diff == 0:
        return e1[:-1] == e2[:-1] and e1 != e2
    elif diff == 1:
        if e1_l < e2_l:
            return e1 == e2[:-1]
        else:
            return e2 == e1[:-1]


def similar(words):
    total_similar = 0
    words.sort()
    for i in range(len(words)):
        element1 = words[i]
        num_similar = 0
        j = i+1
        while j < (len(words)):
            element2 = words[j]
            if is_similar(element1, element2):
                num_similar = num_similar+1
            j = j+1

        if num_similar > 0:
            total_similar = total_similar+num_similar+1
    return total_similar


n = int(input())
words = []
for i in range(n):
    words.append(input())
total_similar = similar(words)
print(total_similar)

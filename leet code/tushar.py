def subtract(n, a, b):
    for i in range(n):
        a[i] = a[i]-b[i]
    return a


def is_all_element_positive(a):
    flag = True
    for i in range(len(a)):
        if a[i] < 0:
            flag = False
    return flag


def sum_all_negative(a):
    summ = 0
    for i in range(len(a)):
        if a[i] < 0:
            summ = summ+a[i]
    return summ


def max_num_buillding(n, a, b, s):
    flag = True
    num_new_building = 0

    while(flag):
        a = subtract(n, a, b)
        if is_all_element_positive(a):
            num_new_building = num_new_building+1
        elif sum_all_negative(a)+s >= 0:
            num_new_building = num_new_building+1
            s = s+sum_all_negative(a)
        else:
            flag = False
    return num_new_building


n = int(input())
s = int(input())
a = [int(e) for e in input().split()]
b = [int(e) for e in input().split()]
k = int(input())

max = max_num_buillding(n, a, b, s)

if max >= k:
    print("1", max)
else:
    print("0", max)

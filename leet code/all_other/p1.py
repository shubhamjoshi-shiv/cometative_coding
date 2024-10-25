
def take_array():
    num_list = [[0, 0, 0], [0, 0, 0], [0, 0, 0]]

    (num_list[0][0],	num_list[0][1],	num_list[0][2]) = [int(x)
                                                        for x in input().split()]
    (num_list[1][0],	num_list[1][2]) = [int(x) for x in input().split()]
    (num_list[2][0],	num_list[2][1],	num_list[2][2]) = [int(x)
                                                        for x in input().split()]

    return num_list


def are_in_ap(num_list):
    check_list = [
        [num_list[0][0], num_list[0][1], num_list[0][2], ],
        [num_list[2][0], num_list[2][1], num_list[2][2], ],
        [num_list[0][0], num_list[1][0], num_list[2][0], ],
        [num_list[0][2], num_list[1][2], num_list[2][2], ]]
    check_list

    num_ap = 0
    for i in check_list:
        if i[2]-i[1] == i[1]-i[0]:
            num_ap = num_ap+1
    return num_ap


def may_be_ap(num_list):
    may_be_list = [
        [num_list[0][0], num_list[2][2]],
        [num_list[1][0], num_list[1][2]],
        [num_list[0][1], num_list[2][1]],
        [num_list[2][0], num_list[0][2]],
    ]
    list_2d = []
    for element in may_be_list:
        list_2d.append(((element[1]-element[0])/2)+(element[0]))
    list_2d
    ap = max([list_2d.count(i) for i in list_2d])
    return ap


test = int(input())
for i in range(test):
    num_list = take_array()
    ap_present = are_in_ap(num_list)
    ap_posible = may_be_ap(num_list)
    total_ap = ap_present + ap_posible
    print(f"Case #{i+1}: {total_ap}")

def get_top_k(freq_list:list,k:int)->list:
    top_k =[]
    for i in freq_list:
        if len(top_k)<k:
            top_k+i
        else:
            break
    return top_k

freq_list=[[], [], [], [], [], [], [], [], [1], [2, 3], [], [4, 5], []]
get_top_k(freq_list,3)


def find(num,n):
    l1=[i for i in range(1,n+1)]
    l2=[int((i*(i+1))/2) for i in l1]
    print(l1)
    print(l2)
    if num in l1:
        return [num]
    elif num in l2:
        return l1[:(l2.index(num)+1)]


test_case=int(input())
for t in range(test_case):
    n,x,y=input().split(sep=" ")
    n=int(n)
    x=int(x)
    y=int(y)
    summ=(n*(n+1))/2
    if summ%(x+y)==0:
        print("Case #",(t+1),": ","POSSIBLE",sep="")
        num=x*(summ/(x+y))
        ans=find(num,n)
        print(len(ans))
        print(' '.join(str(x) for x in ans))
    else:
        print("Case #",(t+1),": ","IMPOSSIBLE",sep="")
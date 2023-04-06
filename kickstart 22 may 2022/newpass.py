def ck(ca,st):
    for i in ca:
        if i in st:
            return True
    return False


T = int(input())
c=[0,0,0,0,0]
lc=["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"]
uc=["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]
dt=["0","1","2","3","4","5","6","7","8","9"]
sc=["#","@","*","&"]
for t in range(T):
    N=int(input())
    op=input()
    np=op
    if ck(lc,op):
        pass
    else:
        np=np+lc[0]
        N+=1
    if ck(uc,op):
        pass
    else:
        np=np+uc[0]
        N+=1
    if ck(dt,op):
        pass
    else:
        np=np+dt[0]
        N+=1
    if ck(sc,op):
        pass
    else:
        np=np+sc[0]
        N+=1
    if (N<7):
        np=np+("1"*(7-N))

    print("Case #",(t+1),": ",np,sep="")

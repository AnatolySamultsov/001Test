a=int(input())
b=int(input())
s=0
n=0
x=(b-a)//3
for i in range(a,b+1,):
  if i%3==0:
    s+=i
    n+=1
    i+=3
print(s/n)
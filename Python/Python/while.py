#pay


a=int(input())
b=int(input())
c=int(input())
d=int(input())

for j in range(c, d+1):
  print('  ', j, end='')
print()
for i in range(a, b+1):
  print(i , end=' ')
  for j in range(c, d+1):
    x=i*j
    print(x ,  end=' ')
  print()
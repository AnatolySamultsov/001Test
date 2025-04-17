
hrs = input("Enter Hours:")
hrs= 45
h = float(hrs)
r = float(10.5)

if h <= 40 :
    p = h * r
    print(p)
elif h > 40 :
    p = (h-40) * 1.5 * r +40 * r
    print(p)
    
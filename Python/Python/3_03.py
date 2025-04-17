score = input("Enter Score: ")
 0.85
sc = float(score)

try:
    print(sc)
except:
    print("Error, please enter numeric input >0 and <1")

if sc>=0.9 :
    print("A")
elif sc>= 0.8 :
    print("B")
elif sc>= 0.7 :
    print("C")
elif sc>= 0.6 :
    print("D")
elif sc< 0.6 :
    print("F")

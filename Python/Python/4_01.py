
if not s:
    print('')
else:
    abc = []
    n = 1
    for i in range(1, len(s)):
        if s[i] == s[i - 1]:
            n += 1
        else:
            abc.append(s[i - 1] + str(n))
            n = 1
    abc.append(s[-1] + str(n))  # Добавляем последнюю группу

    print(''.join(abc))
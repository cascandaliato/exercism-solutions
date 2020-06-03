from collections import Counter


def is_isogram(string):
    return string == "" or Counter([c for c in string.lower() if c.isalpha()]).most_common(1)[0][1] == 1

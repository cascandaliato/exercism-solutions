from collections import Counter
import re


def count_words(sentence):
    return Counter(re.findall(r'([a-zA-Z]+(?:\'[a-zA-Z]+)?|\d+)', sentence.lower()))

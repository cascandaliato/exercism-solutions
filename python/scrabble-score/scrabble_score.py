VALUES = {
    1:  ('A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'),
    2:  ('D', 'G'),
    3:  ('B', 'C', 'M', 'P'),
    4:  ('F', 'H', 'V', 'W', 'Y'),
    5:  ('K'),
    8:  ('J', 'X'),
    10: ('Q', 'Z')
}

LETTER_VALUES = {l: v for v in VALUES.keys() for l in VALUES[v]}


def score(word):
    return sum(LETTER_VALUES[l] for l in word.upper())

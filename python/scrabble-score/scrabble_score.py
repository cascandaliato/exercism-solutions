values = {
    1:  ('A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'),
    2:  ('D', 'G'),
    3:  ('B', 'C', 'M', 'P'),
    4:  ('F', 'H', 'V', 'W', 'Y'),
    5:  ('K'),
    8:  ('J', 'X'),
    10: ('Q', 'Z')
}


def score(word):
    letter_values = {l: v for v in values.keys() for l in values[v]}
    return sum(letter_values[l] for l in word.upper())

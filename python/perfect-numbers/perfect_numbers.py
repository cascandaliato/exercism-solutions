from math import sqrt


def classify(number):
    if number <= 0:
        raise ValueError('Not a natural number')
    aliquot_sum = sum(f for f in range(1, number//2+1) if number % f == 0)
    if aliquot_sum == number:
        return 'perfect'
    elif aliquot_sum > number:
        return 'abundant'
    else:
        return 'deficient'

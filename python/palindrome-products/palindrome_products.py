from itertools import product


def largest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("min_factor must be less or equal than max_factor")
    max_product, factors = None, []
    lower_limit = min_factor
    for a in reversed(range(min_factor, max_factor+1)):
        for b in reversed(range(min_factor, max_factor+1)):
            if b < lower_limit:
                break
            product = a*b
            if is_palindrome(product):
                if max_product == None or product > max_product:
                    max_product, factors = product, [[a, b]]
                    lower_limit = b
                elif product == max_product:
                    factors.append([a, b])
                    lower_limit = max(lower_limit, b)
                break
    return [max_product, factors]


def smallest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("min_factor must be less or equal than max_factor")
    min_product, factors = None, []
    upper_limit = max_factor
    for a in range(min_factor, max_factor + 1):
        for b in range(min_factor, max_factor + 1):
            if b > upper_limit:
                break
            if is_palindrome(a * b):
                if min_product == None or a * b < min_product:
                    min_product, factors = a*b, [[a, b]]
                    upper_limit = b
                elif a * b == min_product:
                    factors.append([a, b])
                    upper_limit = min(upper_limit, b)
                break
    return [min_product, factors]


def is_palindrome(number):
    return str(number) == ''.join(reversed(str(number)))

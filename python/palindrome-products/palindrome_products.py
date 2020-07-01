from itertools import product


def largest(min_factor, max_factor):
    ans = []
    factors = range(min_factor, max_factor+1)
    for a, b in product(factors, factors):
        if is_palindrome(a * b) and (not ans or a * b > ans[0]*ans[1]):
            ans = [a, b]
    return ans


def smallest(min_factor, max_factor):
    ans = []
    factors = range(min_factor, max_factor+1)
    for a, b in product(factors, factors):
        if is_palindrome(a * b) and (not ans or a * b < ans[0]*ans[1]):
            ans = [a, b]
    return ans


def is_palindrome(number):
    return str(number) == ''.join(reversed(str(number)))

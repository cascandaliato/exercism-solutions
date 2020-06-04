def is_valid(isbn):
    isbn = isbn.replace("-", "")
    if len(isbn) == 10 and isbn[:-1].isnumeric() and isbn[-1] in '0123456789X':
        check = 10 if isbn[-1] == 'X' else int(isbn[-1])
        return (sum(int(isbn[i])*(10-i) for i in range(9)) + check) % 11 == 0
    else:
        return False

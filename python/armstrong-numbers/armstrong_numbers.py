def is_armstrong_number(number):
    return number == sum(map(lambda n: int(n)**len(str(number)), str(number)))

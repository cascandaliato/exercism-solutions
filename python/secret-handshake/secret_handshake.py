def commands(number):
    commands = [
        (0, lambda c: c.append('wink')),
        (1, lambda c: c.append('double blink')),
        (2, lambda c: c.append('close your eyes')),
        (3, lambda c: c.append('jump')),
        (4, lambda c: c.reverse())
    ]

    code = []
    for pos, action in commands:
        if (number >> pos) & 1:
            action(code)
    return code

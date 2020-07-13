def commands(number):
    actions = ['wink', 'double blink', 'close your eyes', 'jump']

    code = []
    for action in actions:
        if number & 1:
            code.append(action)
        number >>= 1
    return code[::-1] if number & 1 else code

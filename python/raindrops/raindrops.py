from collections import OrderedDict


def convert(number):
    sounds = OrderedDict([[3, 'Pling'], [5, 'Plang'], [7, 'Plong']])

    return ''.join([sound for (factor, sound) in sounds.items() if number % factor == 0]) or str(number)

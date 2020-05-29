from random import randint
from string import ascii_uppercase
import time

known_robots = []


def letter():
    return ascii_uppercase[randint(0, len(ascii_uppercase)-1)]


class Robot:
    def __init__(self):
        self.reset()

    def reset(self):
        self.name = f'{letter()}{letter()}{str(randint(0,999)).zfill(3)}'
        if self.name in known_robots:
            self.reset()
        else:
            known_robots.append(self.name)

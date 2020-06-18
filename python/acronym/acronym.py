import re


def abbreviate(words):
    return ''.join(re.findall('(?:[^a-zA-Z\']|^)([a-zA-Z])', words)).upper()

def recite(start_verse, end_verse):
    if start_verse < 1 or end_verse > 12 or start_verse > end_verse:
        raise ValueError(
            "The start verse should be less or equal than the end verse and both should be in the range 1-12 (1 and 12 included)")

    nth = ['first',
           'second',
           'third',
           'fourth',
           'fifth',
           'sixth',
           'seventh',
           'eighth',
           'ninth',
           'tenth',
           'eleventh',
           'twelfth']

    gifts = ["twelve Drummers Drumming",
             "eleven Pipers Piping",
             "ten Lords-a-Leaping",
             "nine Ladies Dancing",
             "eight Maids-a-Milking",
             "seven Swans-a-Swimming",
             "six Geese-a-Laying",
             "five Gold Rings",
             "four Calling Birds",
             "three French Hens",
             "two Turtle Doves",
             "a Partridge in a Pear Tree"]

    verses = []
    for day in range(start_verse, end_verse + 1):
        verse = f'On the {nth[day-1]} day of Christmas my true love gave to me: '

        last_gift = ('and ' if day > 1 else '') + gifts[-1]
        verse += ', '.join([gifts[-i]
                            for i in range(day, 1, -1)] + [last_gift])
        verse += '.'

        verses.append(verse)
    return verses

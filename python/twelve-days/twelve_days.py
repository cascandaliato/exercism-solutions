def recite(start_verse, end_verse):
    nth = {1: 'first',
           2: 'second',
           3: 'third',
           4: 'fourth',
           5: 'fifth',
           6: 'sixth',
           7: 'seventh',
           8: 'eighth',
           9: 'ninth',
           10: 'tenth',
           11: 'eleventh',
           12: 'twelfth'}

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
        verse = f'On the {nth[day]} day of Christmas my true love gave to me: '

        last_gift = ('and ' if day > 1 else '') + gifts[-1]
        verse += ', '.join([gifts[-i]
                            for i in range(day, 1, -1)] + [last_gift])
        verse += '.'

        verses.append(verse)
    return verses

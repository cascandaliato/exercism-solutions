import re


class PhoneNumber:
    def __init__(self, number):
        match = re.match(
            r'\D*1?\D*([2-9]\d{2})\D*([2-9]\d{2})\D*(\d{4})\D*$', number)

        if not match:
            raise ValueError("Invalid phone number")

        self.area_code, self.exchange_code, self.subscriber_number = match.groups()
        self.number = self.area_code + self.exchange_code + self.subscriber_number

    def pretty(self):
        return f'({self.area_code}) {self.exchange_code}-{self.subscriber_number}'

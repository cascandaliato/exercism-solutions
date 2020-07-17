from collections import defaultdict


class School:
    def __init__(self):
        self.grades = defaultdict(list)

    def add_student(self, name, grade):
        self.grades[grade].append(name)
        self.grades[grade].sort()

    def roster(self):
        return [student for grade in sorted(self.grades.keys()) for student in self.grades[grade]]

    def grade(self, grade_number):
        return self.grades[grade_number]

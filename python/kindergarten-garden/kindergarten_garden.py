default_students = ['Alice', 'Bob', 'Charlie', 'David',
                    'Eve', 'Fred', 'Ginny', 'Harriet',
                    'Ileana', 'Joseph', 'Kincaid', 'Larry']

plants = {
    'C': 'Clover',
    'G': 'Grass',
    'R': 'Radishes',
    'V': 'Violets'
}


class Garden:
    def __init__(self, diagram, students=default_students):
        self.diagram = diagram.split('\n')
        self.students = sorted(students)

        self.studentPlants = {}
        for student in students:
            idx = self.students.index(student)
            plantShortNames = ''.join(
                map(lambda r: r[2 * idx: 2 * idx + 2], self.diagram))
            self.studentPlants[student] = list(
                map(lambda p: plants[p], plantShortNames))

    def plants(self, student):
        return self.studentPlants[student]

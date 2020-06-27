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
        self.studentPlants = {}
        for idx, student in enumerate(sorted(students)):
            plantShortNames = ''.join(r[2 * idx: 2 * idx + 2]
                                      for r in diagram.split('\n'))
            self.studentPlants[student] = [plants[p] for p in plantShortNames]

    def plants(self, student):
        return self.studentPlants[student]

class Matrix:
    def __init__(self, matrix_string):
        self.matrix = [list(map(int, row.split(' ')))
                       for row in matrix_string.split('\n')]

    def row(self, index):
        return self.matrix[index-1]

    def column(self, index):
        return [rows[index-1] for rows in self.matrix]

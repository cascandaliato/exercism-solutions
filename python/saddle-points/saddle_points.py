def saddle_points(matrix):
    for r in range(1, len(matrix)):
        if len(matrix[r]) != len(matrix[0]):
            raise ValueError(
                "Irregular matrix. All rows should have the same length.")

    num_rows = len(matrix)
    num_cols = len(matrix[0]) if num_rows > 0 else 0

    is_saddle = [[max(matrix[r]) == min(matrix[i][c] for i in range(
        num_rows)) for c in range(num_cols)] for r in range(num_rows)]

    return [{'row': r+1, 'column': c+1} for r in range(num_rows) for c in range(num_cols) if is_saddle[r][c]]

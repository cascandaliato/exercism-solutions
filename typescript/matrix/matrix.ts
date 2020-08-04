class Matrix {
  constructor(private matrix: string) {}

  get rows(): number[][] {
    return this.matrix.split('\n').map((row) => row.split(' ').map(Number));
  }

  get columns(): number[][] {
    return this.transpose(this.rows);
  }

  private transpose(matrix: number[][]): number[][] {
    return matrix[0].reduce(
      (columns, _, colIdx) => [...columns, matrix.map((row) => row[colIdx])],
      [] as number[][],
    );
  }
}

export default Matrix;

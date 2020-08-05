class Matrix {
  constructor(private matrix: string) {}

  get rows(): number[][] {
    return this.splitIntoRows(this.matrix);
  }

  get columns(): number[][] {
    return this.transpose(this.rows);
  }

  private splitIntoRows(matrix: string): number[][] {
    return matrix.split('\n').map((row) => row.split(' ').map(Number));
  }

  private transpose(matrix: number[][]): number[][] {
    return matrix[0].reduce(
      (columns, _, colIdx) => [...columns, matrix.map((row) => row[colIdx])],
      [] as number[][],
    );
  }
}

export default Matrix;

class Matrix {
  constructor(private matrix: string) {}

  get rows(): number[][] {
    return this.matrix.split('\n').map((row) => row.split(' ').map(Number));
  }

  get columns(): number[][] {
    return this.rows[0].reduce(
      (columns, _, colIdx) => [...columns, this.rows.map((row) => row[colIdx])],
      [] as number[][],
    );
  }
}

export default Matrix;

class Matrix {
  rows: number[][] = [];
  columns: number[][] = [];

  constructor(matrix: string) {
    matrix.split('\n').forEach((row) => this.rows.push(row.split(' ').map(Number)));

    this.rows[0].forEach((_, colIdx) => {
      const column = this.rows.map((row) => row[colIdx]);
      this.columns.push(column);
    });
  }
}

export default Matrix;

enum TriangleKind {
  Scalene = 'scalene',
  Isosceles = 'isosceles',
  Equilateral = 'equilateral',
}

export default class Triangle {
  private readonly sides: [number, number, number];

  constructor(...sides: [number, number, number]) {
    this.sides = sides;
  }

  private isTriangle(): boolean {
    const sidesSum = this.sides.reduce((a, b) => a + b);
    return this.sides.every((side) => side > 0 && sidesSum >= 2 * side);
  }

  kind() {
    if (!this.isTriangle()) {
      throw new Error('Not a triangle');
    }

    const [a, b, c] = this.sides;
    if (a == b) {
      if (b == c) return TriangleKind.Equilateral;
      else return TriangleKind.Isosceles;
    } else {
      if (b == c || a == c) return TriangleKind.Isosceles;
      else return TriangleKind.Scalene;
    }
  }
}

type Where = {
  minFactor?: number;
  maxFactor: number;
  sum?: number;
};

class Triplet {
  constructor(
    private readonly a: number,
    private readonly b: number,
    private readonly c: number
  ) {}

  isPythagorean(): boolean {
    return (
      this.a >= 0 &&
      this.b >= 0 &&
      this.c >= 0 &&
      this.a < this.b &&
      this.b < this.c &&
      this.a ** 2 + this.b ** 2 === this.c ** 2
    );
  }

  product(): number {
    return this.a * this.b * this.c;
  }

  sum(): number {
    return this.a + this.b + this.c;
  }

  static where({ minFactor = 0, maxFactor, sum }: Where): Triplet[] {
    const triplets: Triplet[] = [];

    for (let a = minFactor; a < maxFactor; a++)
      for (let b = a + 1; b < maxFactor; b++)
        for (let c = b + 1; c <= maxFactor; c++) {
          const triplet = new Triplet(a, b, c);
          if (
            triplet.isPythagorean() &&
            Triplet.isSumConstraintSatisfied(triplet, sum)
          )
            triplets.push(triplet);
        }

    return triplets;
  }

  private static isSumConstraintSatisfied(
    triplet: Triplet,
    sum?: number
  ): boolean {
    return !sum || triplet.sum() === sum;
  }
}

export default Triplet;

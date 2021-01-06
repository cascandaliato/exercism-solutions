class SumOfMultiples {
  constructor(private multipliers: number[]) {}

  to(limit: number): number {
    let sum = 0;
    for (let i = 0; i < limit; i++)
      if (this.multipliers.some((n) => i % n === 0)) sum += i;
    return sum;
  }
}

function factory(multipliers: number[]) {
  return new SumOfMultiples(multipliers);
}

export default factory;

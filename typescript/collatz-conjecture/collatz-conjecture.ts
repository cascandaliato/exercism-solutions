class CollatzConjecture {
  static steps(n: number): number {
    const COLLATZ_SEQUENCE_END = 1;

    if (n < 1) {
      throw new Error('Only positive numbers are allowed');
    }

    let steps = 0;
    while (n !== COLLATZ_SEQUENCE_END) {
      steps++;
      if (n % 2 == 0) {
        n /= 2;
      } else {
        n = 3 * n + 1;
      }
    }
    return steps;
  }
}

export default CollatzConjecture;
